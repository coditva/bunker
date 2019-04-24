package bunker

import (
    "errors"
)

// Image implements the Command interface. It is responsible to running various
// actions on an image such as remove/update etc.
type Image struct {
    args    *Args
}

// NewImageCommand returns an instance of the command. It takes os arguments
// and parses them to create the command which can be executed.
func NewImageCommand(rawArgs *[]string) (*Image, error) {
    args := make(Args)
    args["command"] = Util.ArrayValueAtIndex(rawArgs, 1)
    args["subcommand"] = Util.ArrayValueAtIndex(rawArgs, 2)

    switch args.Value("subcommand") {
    case "rm":
        args["image"] = Util.ArrayValueAtIndex(rawArgs, 3)

    default:
        return nil, errors.New("image: Unknown subcommand")
    }

    return &Image{ args: &args }, nil
}

// Name returns the name of the command.
func (cmd *Image) Name() string {
    return "image"
}

// Help returns the help information about the command.
func (cmd *Image) Help() string {
    return "image [subcommand] [image]"
}

// Execute runs the command and returns error upon failure.
func (cmd *Image) Execute() error {
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

    //TODO
    return nil
}
