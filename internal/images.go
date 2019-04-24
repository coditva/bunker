package bunker

import (
    "fmt"
    "errors"
)

// Images implements Command interface. It list the images pulled from the
// registry that exist locally.
type Images struct {
    args    *Args
}

// NewImagesCommand returns an instance of the command. It takes os arguments
// and parses them to create the command which can be executed.
func NewImagesCommand(rawArgs *[]string) (*Images, error) {
    args := make(Args)
    if len(*rawArgs) > 1 {
        args["command"] = (*rawArgs)[1]
    } else {
        return nil, errors.New("run: No enough arguments")
    }
    return &Images{ args: &args }, nil
}

// Name returns the name of the command.
func (cmd *Images) Name() string {
    return "images"
}

// Help returns the help information about the command.
func (cmd *Images) Help() string {
    return "images"
}

// Execute runs the command and returns error upon failure.
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
