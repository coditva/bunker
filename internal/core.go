package lib

import (
    types "github.com/coditva/bunker/internal/types"
    rpc "github.com/coditva/bunker/internal/rpc"
)

func ExecuteCommand(command types.Command) error {

    if command.Name == types.CommandPush {
        client := rpc.NewClient(RPCSocketPath)
        if err := client.Connect(); err != nil {
            return err
        }
        defer client.Close()

        var args types.Args
        var reply types.Reply
        if err := client.Call("Api.Push", &args, &reply); err != nil {
            return err
        }
    }
    return nil
}
