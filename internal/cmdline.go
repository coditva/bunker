package lib

import (
    "fmt"
    "errors"

    types "github.com/coditva/bunker/internal/types"
)

func ParseArgs(args []string) (types.Command, error) {
    var command types.CommandName
    var commandArgs string
    var err error

    Logger.Info("Parsing command line arguments")

    if len(args) < 2 {
        command = types.CommandEmpty
        err = errors.New("No command")
    } else if args[1] == "build" {
        command = types.CommandBuild
    } else if args[1] == "pull" {
        command = types.CommandPull
        if len(args) < 3 {
            err = errors.New("Need image name to pull")
        } else {
            commandArgs = args[2]
        }
    } else if args[1] == "push" {
        command = types.CommandPush
    } else {
        command = types.CommandUnknown
        err = errors.New(fmt.Sprintf("Unknown command: %v", args[1]))
    }

    return types.Command{Name: command, Args: []string{commandArgs}}, err
}
