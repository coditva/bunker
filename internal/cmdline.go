package lib

import (
    "fmt"
    "errors"

    types "github.com/coditva/bunker/internal/types"
)

func ParseArgs(args []string) (*types.Command, error) {
    var err error
    var command *types.Command

    Logger.Info("Parsing command line arguments")

    if len(args) < 2 {
        err = errors.New("No command")
        return nil, err
    }

    command, err = NewCommand(args[1])
    if err != nil {
        return nil, err
    }

    for i := 2; i < len(args); i++ {
        if args[i] != "" {
            command.AddArg(args[i])
        }
    }

    return command, nil
}

func NewCommand(name string) (*types.Command, error) {
    command := new(types.Command)
    command.ArgsLen = 0
    command.Name = name

    if name == "pull" {
        command.Method = "Api.Pull"
    } else if name == "images" {
        command.Method = "Api.Images"
    } else if name == "run" {
        command.Method = "Api.Run"
    } else {
        return nil, errors.New(fmt.Sprintf("Unknown command %v", name))
    }

    return command, nil
}
