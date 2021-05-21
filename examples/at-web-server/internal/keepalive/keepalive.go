// Package keepalive は、gRPC コネクションの Keepalive パラメーターを提供します。
package keepalive

import (
	"time"

	"google.golang.org/grpc/keepalive"
)

// ClientParameters は、Keepalive のクライアント設定を返します。
func ClientParameters() keepalive.ClientParameters {
	return keepalive.ClientParameters{
		// PING フレームをサーバーに送信する間隔
		Time: 10 * time.Second,
		// PING フレームに対する ACK が返るまでのタイムアウト
		Timeout: 5 * time.Second,
		// アクティブなストリームが無い時は PING を送信しない
		PermitWithoutStream: false,
	}
}

func EnforcementPolicy() keepalive.EnforcementPolicy {
	return keepalive.EnforcementPolicy{
		// この期間中に一定数の PING が送られると切断する（DoS 対策）
		MinTime: 5 * time.Second,
		// アクティブなストリームが無い時に PING を受信すると切断する
		PermitWithoutStream: false,
	}
}
