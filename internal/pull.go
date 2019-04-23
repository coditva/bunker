package lib

import (
    "fmt"

    types "github.com/coditva/bunker/internal/types"
)

func Pull(args *types.Args, reply *string) error {
    imageName := (*args)[0]
    if imageName == "" {
        *reply = "No image specified to pull from registry"
        Logger.Warning(*reply)
        return nil
    }

    retry := 2
    for retry > 0 {
        Logger.Info("Pulling image ", imageName)
        image, err := ContainerdClient.Client.Pull(ContainerdClient.Ns, imageName)
        if err != nil {
            Logger.Warning("Failed to pull image: ", err)
            imageName = fmt.Sprintf("docker.io/library/%v", imageName)
            retry -= 1
        } else {
            *reply = fmt.Sprintf("Pulled image %v (%v)", imageName, image.Name)
            Logger.Info(*reply)
            return nil
        }
    }

    *reply = fmt.Sprintf("Could not find image %v to pull", (*args)[0])

    return nil
}
