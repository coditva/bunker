package lib

import "./types"

func ExecuteCommand(command types.Command) error {

    if command.Name == types.CommandPull {
        PullImage(types.Image{Name: "some/image"})
    }
    return nil
}
