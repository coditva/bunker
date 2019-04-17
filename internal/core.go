package lib

import "github.com/coditva/bunker/internal/types"

func ExecuteCommand(command types.Command) error {

    if command.Name == types.CommandPull {
        PullImage(types.Image{Name: "some/image"})
    }
    return nil
}
