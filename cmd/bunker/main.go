package main

import (
    "os"

    "github.com/coditva/bunker/internal"
)

func main() {
    lib.InitLogger("bunker", "/tmp/bunker.log")

    lib.Logger.Info("Starting application")

    lib.Logger.Info("Parsing command line arguments")
    cmd, err := lib.ParseArgs(os.Args)
    if err != nil {
        lib.PrintHelp(err)
        os.Exit(1)
    } else {
        lib.Logger.Info("Command: ", cmd.Name)
    }

    lib.Logger.Info("Executing command")
    err = lib.ExecuteCommand(cmd)
    if err != nil {
        lib.Logger.Error(err)
    }

    lib.Logger.Info("Exiting")
}
