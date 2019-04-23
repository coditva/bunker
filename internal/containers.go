package lib

import (
    "fmt"

    types "github.com/coditva/bunker/internal/types"
)

func Containers(args *types.Args, reply *string) error {
    Logger.Info("Getting containers from containerd")
    containers, err := ContainerdClient.Client.Containers(ContainerdClient.Ns, "")
    if err != nil {
        return err
    }

    *reply = "Name\t\tImage\n--------------\t--------------"
    for _, container := range containers {
        imageName := "-"

        name := container.ID()
        image, err := container.Image(ContainerdClient.Ns)
        if err != nil {
            Logger.Warning("Unknown image for container ", name, ": ", err)
        }
        imageName = image.Name()

        *reply = fmt.Sprintf("%v\n%v\t%v", *reply, name, imageName)
    }

    return nil
}
