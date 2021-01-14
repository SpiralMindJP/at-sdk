// Package config は、at-server アプリケーションの設定ファイルを読み込みます。
package config

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
)

type Config struct {
	BindHost string `toml:"bind_host"`
	BindPort int    `toml:"bind_port"`

	ServerHost string `toml:"server_host"`
	ServerPort int    `toml:"server_port"`

	EnableTLS   bool   `toml:"enable_tls"`
	TLSCertFile string `toml:"tls_cert_file"`
	TLSKeyFile  string `toml:"tls_key_file"`

	Core     *Core     `toml:"core"`
	Firebase *Firebase `toml:"firebase"`
}

func (c *Config) Validate() error {
	if c.Firebase == nil {
		return errors.New("firebase settings are not found")
	}
	if c.Core == nil {
		return errors.New("core server settings are not found")
	}

	if c.Core != nil {
		if err := c.Core.Validate(); err != nil {
			return err
		}
	}
	if c.Firebase != nil {
		if err := c.Firebase.Validate(); err != nil {
			return err
		}
	}

	if c.EnableTLS {
		if c.TLSCertFile == "" {
			return errors.New("TLS cert file is not specified")
		}
		if c.TLSKeyFile == "" {
			return errors.New("TLS key file is not specified")
		}
	}

	return nil
}

func (c *Config) Merge(other *Config) {
	if c == nil || other == nil {
		return
	}

	if other.BindHost != "" {
		c.BindHost = other.BindHost
	}
	if other.BindPort > 0 {
		c.BindPort = other.BindPort
	}
	if other.ServerHost != "" {
		c.ServerHost = other.ServerHost
	}
	if other.ServerPort > 0 {
		c.ServerPort = other.ServerPort
	}
	if other.EnableTLS {
		c.EnableTLS = other.EnableTLS
	}
	if other.TLSCertFile != "" {
		c.TLSCertFile = other.TLSCertFile
	}
	if other.TLSKeyFile != "" {
		c.TLSKeyFile = other.TLSKeyFile
	}
}

func (c *Config) MergeToDefault() *Config {
	config := defaultConfig

	config.Merge(c)

	if c.Core != nil {
		core := defaultCore
		core.Merge(c.Core)
		config.Core = &core
	}
	if c.Firebase != nil {
		firebase := defaultFirebase
		firebase.Merge(c.Firebase)
		config.Firebase = &firebase
	}

	return &config
}

var defaultConfig = Config{
	BindHost:   "localhost",
	BindPort:   8888,
	ServerHost: "localhost",
	ServerPort: 8888,
}

func LoadFile(name string) (*Config, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	return Load(file)
}

func Load(r io.Reader) (*Config, error) {
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// 環境変数を展開
	str := os.ExpandEnv(string(buf))

	c := new(Config)
	if _, err := toml.Decode(str, c); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	config := c.MergeToDefault()

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) BindAddr() string {
	return net.JoinHostPort(c.BindHost, strconv.Itoa(c.BindPort))
}
