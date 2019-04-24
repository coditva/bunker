package lib

import (
    "errors"
)

type Pull struct {
    args    *Args
}

func NewPullCommand(rawArgs *[]string) (*Pull, error) {
    args := make(Args)
    if len(*rawArgs) > 2 {
        args["command"] = (*rawArgs)[1]
        args["image"] = (*rawArgs)[2]
    } else {
        return nil, errors.New("pull: No enough arguments")
    }
    return &Pull{ args: &args }, nil
}

func (cmd *Pull) Name() string {
    return "pull"
}

func (cmd *Pull) Help() string {
    return "pull [image]"
}

func (cmd *Pull) Execute() error {
    if cmd.args.Value("image") == "" {
        err := errors.New("No image name specified")
        Logger.Warning(err)
        return err
    }
    imageName := Util.ImageNameToRegistryURL(cmd.args.Value("image"))

    containerd, err := NewContainerd()
    if err != nil {
        Logger.Error(err)
        return err
    }

    image, err := containerd.Client.Pull(containerd.Context, imageName)
    if err != nil {
        Logger.Warning(err)
        return err
    }
    Logger.Info("Pulled image ", image.Name())
    return nil
}
