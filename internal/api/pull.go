package api

import (
    lib "github.com/coditva/bunker/internal"
    types "github.com/coditva/bunker/internal/types"
)

func (api Api) Pull(args *types.Args, reply *types.Reply) error {
    lib.Logger.Info("Pulling image")
    image, err := lib.ContainerdClient.Client.Pull(lib.ContainerdClient.Ns,
            "docker.io/library/redis:latest")
    if err != nil {
        lib.Logger.Error(err)
        return err
    }
    lib.Logger.Info("Pulled image: ", image.Name)

    *reply = "Pulled image"
    return err
}
