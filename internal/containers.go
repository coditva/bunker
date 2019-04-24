package bunker

import (
    "fmt"
    "errors"
)

// Containers implements Command interface. It is responsible to list containers
// that exist locally.
type Containers struct {
    args    *Args
}

// NewContainersCommand returns an instance of the command. It takes os
// arguments and parses them to create the command which can be executed.
func NewContainersCommand(rawArgs *[]string) (*Containers, error) {
    args := make(Args)
    if len(*rawArgs) > 1 {
        args["command"] = (*rawArgs)[1]
    } else {
        return nil, errors.New("run: No enough arguments")
    }
    return &Containers{ args: &args }, nil
}

// Name returns the name of the command.
func (cmd *Containers) Name() string {
    return "containers"
}

// Help returns the help information about the command.
func (cmd *Containers) Help() string {
    return "containers"
}

// Execute runs the command and returns error upon failure.
func (cmd *Containers) Execute() error {
    containerd, err := NewContainerd()
    if err != nil {
        Logger.Error(err)
        return err
    }

    Logger.Info("Getting list of containers from containerd")
    containers, err := containerd.Client.Containers(containerd.Context, "")
    if err != nil {
        return err
    }

    fmt.Println("Name\t\tImage\n--------------\t--------------")
    for _, container := range containers {
        name := container.ID()
        image, err := container.Image(containerd.Context)
        if err != nil {
            Logger.Error("Unknown image for container ", name, ": ", err)
            return err
        }
        fmt.Printf("%v\t%v\n", name, image.Name())
    }

    return nil
}
