package lib

import "errors"
import "./types"

func ParseArgs(args []string) (types.Command, error) {
    var command types.CommandName
    var err error

    if len(args) < 2 {
        command = types.CommandEmpty
        err = errors.New("No command")
    } else if args[1] == "build" {
        command = types.CommandBuild
    } else if args[1] == "pull" {
        command = types.CommandPull
    } else if args[1] == "push" {
        command = types.CommandPush
    } else {
        command = types.CommandUnknown
        err = errors.New("Command unknown")
    }

    return types.Command{Name: command}, err
}
