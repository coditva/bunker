package main

import (
    "os"
    "fmt"

    lib "github.com/coditva/bunker/internal"
    rpc "github.com/coditva/bunker/internal/rpc"
)

func main() {
    var reply string

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
        os.Exit(1)
    }
    defer client.Close()

    lib.Logger.Info("Executing command: ", command.Name)
    if err := client.Call(command.Method, &command.Args, &reply); err != nil {
        lib.Logger.Error(err)
        os.Exit(1)
    }
    fmt.Println(reply)

    lib.Logger.Info("Exiting")
}
