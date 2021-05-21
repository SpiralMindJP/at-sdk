// Package enum は、AT サーバーで使用する各種列挙体を定義します。
package enum

//go:generate stringer -output enum_string.go -type=DeviceType,DeviceConnectionState,ContentType,VideoFrameType

// DeviceType は、デバイスの種類を表します。
type DeviceType int

const (
	Operator DeviceType = 0 // Operator デバイス。
	Customer DeviceType = 1 // Customer デバイス。
)

// DeviceConnectionState は、デバイスの接続状態を表します。
type DeviceConnectionState int

const (
	Offline DeviceConnectionState = 0 // デバイスがオフライン。
	Online  DeviceConnectionState = 1 // デバイスがオンライン。
)

// ContentType は、コンテンツの種類を表します。
type ContentType int32

const (
	Image     ContentType = 0   // 画像コンテンツ。
	Video     ContentType = 1   // ビデオコンテンツ。
	Avatar    ContentType = 2   // アバターVRMコンテンツ。
	Animation ContentType = 3   // アニメーションコンテンツ。
	Other     ContentType = 999 // その他のコンテンツ。
)

// VideoFrameType は、ビデオフレームの種類を表します。
type VideoFrameType int

const (
	Unknown VideoFrameType = 0 // 不明なビデオフレーム。
	JPEG    VideoFrameType = 1 // JPEG ビデオフレーム。
)
