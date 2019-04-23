package lib

import (
    "fmt"

    types "github.com/coditva/bunker/internal/types"
)

func Containers(args *types.Args, reply *string) error {
    containerd, err := NewContainerd()
    if err != nil {
        Logger.Error(err)
        return err
    }

    Logger.Info("Getting list of containers")
    containers, err := containerd.Client.Containers(containerd.Context, "")
    if err != nil {
        return err
    }

    *reply = "Name\t\tImage\n--------------\t--------------"
    for _, container := range containers {
        imageName := "-"

        name := container.ID()
        image, err := container.Image(containerd.Context)
        if err != nil {
            Logger.Warning("Unknown image for container ", name, ": ", err)
        }
        imageName = image.Name()

        *reply = fmt.Sprintf("%v\n%v\t%v", *reply, name, imageName)
    }

    return nil
}
