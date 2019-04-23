package lib

import (
    "fmt"

    types "github.com/coditva/bunker/internal/types"
)

func Images(args *types.Args, reply *string) error {
    containerd, err := NewContainerd()
    if err != nil {
        Logger.Error(err)
        return err
    }

    Logger.Info("Getting images from containerd")
    images, err := containerd.Client.ListImages(containerd.Context, "")
    if err != nil {
        return err
    }

    *reply = "Size\tImage\n------\t--------------"
    for _, image := range images{
        name := image.Name()
        size, err := image.Size(containerd.Context)
        if err != nil {
            Logger.Warning("Unknown size for image ", name, ": ", err)
            size = 0
        }
        *reply = fmt.Sprintf("%v\n%v\t%v", *reply, Util.ByteToString(size), name)
    }

    return nil
}
