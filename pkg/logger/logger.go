package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func init() {
	if Logger == nil {
		Logger = slog.New(slog.NewTextHandler(os.Stdout,nil))
	}
}


