package config

import (
	"net"
	"strconv"
)

// Core は Core サーバーの接続設定を表します。
type Core struct {
	// ホスト。
	Host string `toml:"host"`
	// ポート。
	Port int `toml:"port"`
	// true を設定すると、トランスポートセキュリティを有効にします。
	Secure bool `toml:"secure"`
}

var defaultCore = Core{
	Host: "localhost",
	Port: 8080,
}

func (c *Core) Validate() error {
	return nil
}

func (c *Core) Merge(other *Core) {
	if c == nil || other == nil {
		return
	}

	if other.Host != "" {
		c.Host = other.Host
	}
	if other.Port > 0 {
		c.Port = other.Port
	}
	if other.Secure {
		c.Secure = other.Secure
	}
}

func (c *Core) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}
