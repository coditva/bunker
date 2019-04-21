package api

import (
    "fmt"

    lib "github.com/coditva/bunker/internal"
    types "github.com/coditva/bunker/internal/types"
)

func (api Api) Pull(args *types.Args, reply *string) error {
    imageName := (*args)[0]
    lib.Logger.Info("Pulling image ", imageName)
    image, err := lib.ContainerdClient.Client.Pull(lib.ContainerdClient.Ns, imageName)
    if err != nil {
        lib.Logger.Error(err)
        return err
    }

    *reply = fmt.Sprintf("Pulled image %v (%v)", imageName, image.Name)
    lib.Logger.Info(*reply)
    return err
}
