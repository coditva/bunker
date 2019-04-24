package lib

import (
    "fmt"
    "errors"
)

type Images struct {
    args    *Args
}

func NewImagesCommand(rawArgs *[]string) (*Images, error) {
    args := make(Args)
    if len(*rawArgs) > 1 {
        args["command"] = (*rawArgs)[1]
    } else {
        return nil, errors.New("run: No enough arguments")
    }
    return &Images{ args: &args }, nil
}

func (cmd *Images) Name() string {
    return "images"
}

func (cmd *Images) Help() string {
    return "images"
}


func (cmd *Images) Execute() error {
    containerd, err := NewContainerd()
    if err != nil {
        Logger.Error(err)
        return err
    }

    Logger.Info("Getting list of images from containerd")
    images, err := containerd.Client.ListImages(containerd.Context, "")
    if err != nil {
        return err
    }

    fmt.Println("Size\tImage\n------\t--------------")
    for _, image := range images{
        name := image.Name()
        size, err := image.Size(containerd.Context)
        if err != nil {
            Logger.Warning("Unknown size for image ", name, ": ", err)
            size = 0
        }
        fmt.Printf("%v\t%v\n", Util.ByteToString(size), name)
    }

    return nil
}
