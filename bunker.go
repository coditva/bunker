package main

import "fmt"
import (
    "github.com/google/logger"
    "os"
)

func main() {
    var logger_out = os.Stdout
    defer logger.Init("defaultLogger", false, true, logger_out).Close()

    fmt.Printf("CLI for the bunker\n")

    logger.Info("Info")
    logger.Warning("Warning")
    logger.Error("Error")
    logger.Fatal("Fatal")
}
