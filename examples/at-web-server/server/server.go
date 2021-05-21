// Package server は、Avatar Teleporter の Web サーバーを提供します。
package server

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	webauth "github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/auth"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/config"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/keepalive"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/log"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/path"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/template"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	config *config.Config
	opts   *serverOptions
	app    *firebase.App
	client *auth.Client
	conn   *grpc.ClientConn
}

func NewServer(ctx context.Context, config *config.Config, opts ...ServerOption) (*Server, error) {
	op := defaultOptions
	for _, opt := range opts {
		opt.apply(&op)
	}

	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(config.Firebase.Credentials))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase Authentication client: %w", err)
	}

	grpcOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters()),
	}
	if config.Core.Secure {
		creds := credentials.NewTLS(&tls.Config{})
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(creds))
	} else {
		grpcOpts = append(grpcOpts, grpc.WithInsecure())
	}

	addr := config.Core.Addr()
	conn, err := grpc.DialContext(ctx, addr, grpcOpts...)
	if err != nil {
		return nil, fmt.Errorf("faield to dial to core server: %s: %w", addr, err)
	}

	return &Server{
		config: config,
		opts:   &op,
		app:    app,
		client: client,
		conn:   conn,
	}, nil
}

var (
	_ webauth.Server = (*Server)(nil)
)

func (s *Server) Close() error {
	return s.conn.Close()
}

func (s *Server) Serve(ctx context.Context) error {
	config := s.config

	template.LoadTemplate(TemplateFS())

	router, err := s.NewRouter()
	if err != nil {
		return err
	}

	addr := config.BindAddr()
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		<-ctx.Done()
		_ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return server.Shutdown(_ctx)
	})

	g.Go(func() error {
		var err error
		if config.EnableTLS {
			err = server.ListenAndServeTLS(config.TLSCertFile, config.TLSKeyFile)
		} else {
			err = server.ListenAndServe()
		}
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("failed to serve HTTP: %w", err)
	}

	return nil
}

func (s *Server) RedirectToLogin(w http.ResponseWriter, r *http.Request, addRedirect bool) {
	location := s.BuildURL(path.Login)

	if addRedirect {
		currentURL := s.BuildURL(r.URL.String())
		query := location.Query()
		query.Add("redirect", url.QueryEscape(currentURL.String()))
		location.RawQuery = query.Encode()
	}

	w.Header().Add("Location", location.String())
	w.WriteHeader(http.StatusSeeOther)
}

func (s *Server) BuildURL(path string) *url.URL {
	baseURL := s.baseURL(false)

	_url, err := baseURL.Parse(path)
	if err != nil {
		panic(err)
	}

	return _url
}

func (s *Server) BuildWebSocketURL(path string) *url.URL {
	baseURL := s.baseURL(true)

	_url, err := baseURL.Parse(path)
	if err != nil {
		panic(err)
	}

	return _url
}

func (s *Server) baseURL(webSocket bool) *url.URL {
	config := s.config

	var scheme string
	if webSocket {
		if config.EnableTLS {
			scheme = "wss"
		} else {
			scheme = "ws"
		}
	} else {
		if config.EnableTLS {
			scheme = "https"
		} else {
			scheme = "http"
		}
	}

	var port int
	port, err := defaultPort(scheme)
	if err != nil {
		panic(err)
	}

	host := config.ServerHost
	if config.ServerPort != port {
		host = net.JoinHostPort(host, strconv.Itoa(config.ServerPort))
	}

	baseURL := &url.URL{
		Scheme: scheme,
		Host:   host,
	}

	return baseURL
}

var defaultPortMap = map[string]int{
	"http":  80,
	"https": 443,
	"ws":    80,
	"wss":   443,
}

func defaultPort(service string) (port int, err error) {
	port, ok := defaultPortMap[service]
	if !ok {
		err = fmt.Errorf("invalid service name [%s]", service)
	}

	return
}

var defaultOptions = serverOptions{
	logger: log.NewNopLogger(),
}

type serverOptions struct {
	logger log.Logger
}

type ServerOption interface {
	apply(opt *serverOptions)
}

type serverOptionFunc func(opt *serverOptions)

func (f serverOptionFunc) apply(opt *serverOptions) {
	f(opt)
}

func Logger(logger log.Logger) ServerOption {
	if logger == nil {
		panic("logger is nil")
	}

	return serverOptionFunc(func(opt *serverOptions) {
		opt.logger = logger
	})
}
