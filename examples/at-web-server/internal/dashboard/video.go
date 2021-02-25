package dashboard

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"time"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/log"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/middleware"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/webutil"
	pb "github.com/SpiralMindJP/at-sdk/go/pb/core"
)

type motionJPEGStream struct {
	w  http.ResponseWriter
	m  *multipart.Writer
	h  textproto.MIMEHeader
	st string
}

func newMotionJPEGStream(w http.ResponseWriter) *motionJPEGStream {
	m := multipart.NewWriter(w)

	w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary="+m.Boundary())
	w.Header().Set("Connection", "close")

	return &motionJPEGStream{
		w:  w,
		m:  m,
		h:  textproto.MIMEHeader{},
		st: fmt.Sprint(time.Now().Unix()),
	}
}

func (s *motionJPEGStream) Write(b []byte) error {
	s.h.Set("Content-Type", "image/jpeg")
	s.h.Set("Content-Length", fmt.Sprint(len(b)))
	s.h.Set("X-StartTime", s.st)
	s.h.Set("X-TimeStamp", fmt.Sprint(time.Now().Unix()))

	mw, err := s.m.CreatePart(s.h)
	if err != nil {
		return fmt.Errorf("failed to create a new part for multipart: %w", err)
	}

	_, err = mw.Write(b)
	if err != nil {
		return fmt.Errorf("failed to write JPEG frame: %w", err)
	}

	if flusher, ok := mw.(http.Flusher); ok {
		flusher.Flush()
	}

	return nil
}

func (s *motionJPEGStream) Close() error {
	return s.m.Close()
}

type videoStream struct {
	client pb.DashboardService_VideoStreamClient
	frames chan *pb.VideoFrame
	exit   func()
	err    error
}

func readVideoStream(ctx context.Context, client pb.DashboardService_VideoStreamClient) *videoStream {
	ctx, cancel := context.WithCancel(ctx)

	s := &videoStream{
		client: client,
		frames: make(chan *pb.VideoFrame, 128),
		exit:   cancel,
	}

	go s.read(ctx)

	return s
}

func (s *videoStream) Close() error {
	s.exit()
	return s.err
}

func (s *videoStream) Frames() <-chan *pb.VideoFrame {
	return s.frames
}

func (s *videoStream) read(ctx context.Context) {
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
		}

		frame, err := s.client.Recv()
		if errors.Is(err, io.EOF) {
			break loop
		}
		if err != nil {
			s.err = fmt.Errorf("failed to receive video stream: %w", err)
			break loop
		}

		select {
		case s.frames <- frame:
		default:
		}
	}
}

func makeBlankJPEG() []byte {
	const w = 320
	const h = 240
	blank := image.NewRGBA(image.Rect(0, 0, w, h))

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			blank.SetRGBA(x, y, color.RGBA{0, 0, 0, 0xFF})
		}
	}

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, blank, &jpeg.Options{Quality: 80}); err != nil {
		panic("failed to create blank JPEG: " + err.Error())
	}

	return buf.Bytes()
}

// VideoStreamHandler は、以下のリクエストを処理するハンドラーです。
// POST /dashboard/rooms/[:room_id]/video
func VideoStreamHandler() http.HandlerFunc {
	blank := makeBlankJPEG()

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.UserFromContext(ctx)
		room := middleware.RoomFromRequest(r)
		grpcConn := middleware.GRPCConnFromContext(ctx)

		teamID := user.TeamID()
		roomID := room.GetRoomId()

		dashboardService := pb.NewDashboardServiceClient(grpcConn)

		logger := log.FromContext(ctx).With("method", "dashboard.VideoStreamHandler", "team_id", teamID)

		client, err := dashboardService.VideoStream(ctx, &pb.VideoStreamRequest{
			TeamId: teamID,
			RoomId: roomID,
		})
		if err != nil {
			err = fmt.Errorf("failed to get video stream: %w", err)
			logger.Errorw("failed to get video stream", "err", err)
			webutil.WriteError(w, r, "failed to get video stream", err, http.StatusInternalServerError)
			return
		}
		defer func() {
			err := client.CloseSend()
			if errors.Is(err, context.Canceled) {
				return
			}
			if err != nil {
				logger.Errorw("failed to close send stream", "err", err)
			}
		}()

		vs := readVideoStream(ctx, client)
		defer func() {
			if err := vs.Close(); err != nil {
				logger.Errorw("failed to receive video stream", "err", err)
			}
		}()

		s := newMotionJPEGStream(w)
		defer s.Close()
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			case frame := <-vs.Frames():
				switch frame.GetType() {
				case pb.VideoFrameType_VIDEO_FRAME_JPEG:
					if err := s.Write(frame.Payload); err != nil {
						err = fmt.Errorf("failed to write JPEG frame: %w", err)
						logger.Errorw("failed to write JPEG frame", "err", err)
						break loop
					}
				}
			case <-time.After(1 * time.Second):
				if err := s.Write(blank); err != nil {
					err = fmt.Errorf("failed to write JPEG frame: %w", err)
					logger.Errorw("failed to write JPEG frame", "err", err)
					break loop
				}
			}

		}
	}
}
