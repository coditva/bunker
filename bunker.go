package main

import (
    "github.com/google/logger"
    "./internal"
    "os"
)

func main() {
    logPath := "/tmp/bunker.log"
    loggerOut, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
    if err != nil {
        logger.Fatalf("Failed to open logfile")
    }
    defer logger.Init("defaultLogger", false, true, loggerOut).Close()

    logger.Info("Starting application")

    logger.Info("Parsing command line arguments")
    cmd, err := lib.ParseArgs(os.Args)
    if err != nil {
        lib.PrintHelp(err)
        os.Exit(1)
    } else {
        logger.Info("Command: ", cmd.Name)
    }

    logger.Info("Executing command")
    err = lib.ExecuteCommand(cmd)
    if err != nil {
        logger.Error(err)
    }

    logger.Info("Exiting")
}
