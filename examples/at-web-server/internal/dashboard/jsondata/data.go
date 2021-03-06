package jsondata

import "github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/enum"

type RoomStatus string

const (
	RoomStatusOperating       RoomStatus = "operating"
	RoomStatusOthersOperating RoomStatus = "others-operating"
	RoomStatusNotOperating    RoomStatus = "not-operating"
)

type ErrorCode int

const (
	ErrorUnknown ErrorCode = iota + 1
	ErrorDeviceNotRegistered
	ErrorJoinRoomFailed
	ErrorLeaveRoomFailed
	ErrorContentNotFound
)

func (c ErrorCode) String() string {
	switch c {
	case ErrorDeviceNotRegistered:
		return "デバイスが登録されていません。"
	case ErrorJoinRoomFailed:
		return "入室に失敗しました。"
	case ErrorLeaveRoomFailed:
		return "退室に失敗しました。"
	case ErrorContentNotFound:
		return "コンテンツが見つかりません。"
	default:
		return "エラーが発生しました。"
	}
}

type Error struct {
	Code    ErrorCode `json:"code,omitempty"`
	Message string    `json:"message,omitempty"`
}

type DashboardData struct {
	Rooms    []*Room    `json:"rooms,omitempty"`
	Room     *Room      `json:"room,omitempty"`
	Contents []*Content `json:"contents,omitempty"`
	Error    *Error     `json:"error,omitempty"`
}

type Room struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`

	OperatorDeviceID int64                      `json:"operator_device_id"`
	ConnectionState  enum.DeviceConnectionState `json:"connection_state"`

	Status RoomStatus `json:"status,omitempty"`
}

type Content struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (r *Room) SetStatus(deviceID int64) {
	status := RoomStatusNotOperating
	if r.OperatorDeviceID > 0 {
		if r.OperatorDeviceID == deviceID {
			status = RoomStatusOperating
		} else {
			status = RoomStatusOthersOperating
		}
	}
	r.Status = status
}

type EventType int

const (
	RoomCreated EventType = 0 // ルームが作成された。
	RoomDeleted EventType = 1 // ルームが削除された。

	DeviceJoined  EventType = 10 // デバイスが入室した。
	DeviceLeaved  EventType = 11 // デバイスが退室した。
	DeviceDeleted EventType = 12 // デバイスが削除された。

	DeviceOffline EventType = 20 // デバイスがオフラインになった。
	DeviceOnline  EventType = 21 // デバイスがオンラインになった。
)

type Event struct {
	Type EventType `json:"type"`
	Room *Room     `json:"room,omitempty"`
}
