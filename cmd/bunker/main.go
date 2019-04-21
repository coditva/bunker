package main

import (
    "os"

    lib "github.com/coditva/bunker/internal"
    rpc "github.com/coditva/bunker/internal/rpc"
)

func main() {
    lib.InitLogger("bunker", "/tmp/bunker.log")

    command, err := lib.ParseArgs(os.Args)
    if err != nil {
        lib.PrintHelp(err)
        os.Exit(1)
    }

    lib.Logger.Info("Connecting to bunkerd")
    client := rpc.NewClient(lib.RPCSocketPath)
    if err := client.Connect(); err != nil {
        lib.Logger.Error(err)
    }
    defer client.Close()

    lib.Logger.Info("Executing command: ", command.Name)
    var reply string
    if err := client.Call("Api.Pull", &command.Args, &reply); err != nil {
        lib.Logger.Error(err)
    }

    lib.Logger.Info("Exiting")
}
