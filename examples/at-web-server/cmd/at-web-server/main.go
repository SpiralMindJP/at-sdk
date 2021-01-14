package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/config"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/log"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/log/std"
	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/server"
)

func _main() error {
	var configFile string
	var verbose bool
	flag.StringVar(&configFile, "c", "config.toml", "config file")
	flag.BoolVar(&verbose, "verbose", false, "makes ATP server verbose during operation")
	flag.Parse()

	logger := std.NewStdLogger()

	if verbose {
		logger.SetLevel(log.Debug)
	} else {
		logger.SetLevel(log.Info)
	}

	config, err := config.LoadFile(configFile)
	if err != nil {
		return err
	}

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	srv, err := server.NewServer(ctx, config, server.Logger(logger))
	if err != nil {
		logger.Errorf("failed to initialize server: %v", err)
	}
	defer srv.Close()

	if err := srv.Serve(ctx); err != nil {
		logger.Errorf("failed to start server: %v", err)
		os.Exit(1)
	}

	return nil
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
