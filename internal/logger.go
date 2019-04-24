package bunker

import (
    "os"

    "github.com/google/logger"
)

// Logger is a global logger instance. This is initialized once by the CLI or
// Daemon.
var Logger *logger.Logger


// InitLogger creates a new instance of logger with the given name and path of
// the file to log to. The instance is stored as a global logger instance.
func InitLogger(name string, path string) {
    out, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
    if err != nil {
        logger.Fatalf("Failed to open logfile")
    }

    Logger = logger.Init("defaultLogger", false, true, out)
}
