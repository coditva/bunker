package api

import (
    "fmt"

    lib "github.com/coditva/bunker/internal"
    types "github.com/coditva/bunker/internal/types"
)

func (api Api) Containers(args *types.Args, reply *string) error {
    lib.Logger.Info("Getting containers from containerd")
    containers, err := lib.ContainerdClient.Client.Containers(lib.ContainerdClient.Ns, "")
    if err != nil {
        return err
    }

    for _, container := range containers {
        imageName := "-"

        name := container.ID()
        image, err := container.Image(lib.ContainerdClient.Ns)
        if err != nil {
            lib.Logger.Warning("Unknown image for container ", name, ": ", err)
        }
        imageName = image.Name()

        *reply = fmt.Sprintf("%v%v\t%v\n", *reply, name, imageName)
    }

    return nil
}
