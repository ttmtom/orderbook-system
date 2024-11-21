package logger

import (
	"github.com/dpotapov/slogpfx"
	"log/slog"
	"os"
)

var logger *slog.Logger

func Init() {
	s := slog.NewTextHandler(os.Stderr, nil)

	prefixed := slogpfx.NewHandler(s, &slogpfx.HandlerOptions{
		PrefixKeys: []string{"service"},
	})

	logger = slog.New(prefixed)

	slog.SetDefault(logger)
}
