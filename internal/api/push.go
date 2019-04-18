package api

import (
    types "github.com/coditva/bunker/internal/types"
)

func (api Api) Push(args *types.Args, reply *types.Reply) error {
    *reply = "Pushing image"
    return nil
}
