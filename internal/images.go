package lib

import (
    "fmt"

    util "github.com/coditva/bunker/internal/util"
    types "github.com/coditva/bunker/internal/types"
)

func Images(args *types.Args, reply *string) error {
    Logger.Info("Getting images from containerd")
    images, err := ContainerdClient.Client.ListImages(ContainerdClient.Ns, "")
    if err != nil {
        return err
    }

    *reply = "Size\tImage\n------\t--------------"
    for _, image := range images{
        name := image.Name()
        size, err := image.Size(ContainerdClient.Ns)
        if err != nil {
            Logger.Warning("Unknown size for image ", name, ": ", err)
            size = 0
        }
        *reply = fmt.Sprintf("%v\n%v\t%v", *reply, util.SizeToString(size), name)
    }

    return nil
}
