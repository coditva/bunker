package bunker

import (
    "errors"
)

// Pull implements Command interface. It takes care of pulling and image from
// the registry.
type Pull struct {
    args    *Args
}

// NewPullCommand returns an instance of the command. It takes os arguments and
// parses them to create the command which can be executed.
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

// Name returns the name of the command.
func (cmd *Pull) Name() string {
    return "pull"
}

// Help returns the help information about the command.
func (cmd *Pull) Help() string {
    return "pull [image]"
}

// Execute runs the command and returns error upon failure.
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
