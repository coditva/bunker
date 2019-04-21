package api

import (
    "fmt"

    lib "github.com/coditva/bunker/internal"
    types "github.com/coditva/bunker/internal/types"
)

func (api Api) Images(args *types.Args, reply *string) error {
    lib.Logger.Info("Getting images from containerd")
    images, err := lib.ContainerdClient.Client.ListImages(lib.ContainerdClient.Ns, "")
    if err != nil {
        return err
    }

    for _, image := range images{
        *reply = fmt.Sprintf("%v%v\n", *reply, image.Name())
    }

    return nil
}
