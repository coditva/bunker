package bunker

import (
    "os"

    "github.com/google/logger"
)

var Logger *logger.Logger

func InitLogger(name string, path string) {
    out, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
    if err != nil {
        logger.Fatalf("Failed to open logfile")
    }

    Logger = logger.Init("defaultLogger", false, true, out)
}
