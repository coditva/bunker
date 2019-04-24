package bunker

import (
    "os"
    "fmt"
    "errors"
)

// Rm implements Command interface. It takes care of removing a container
// locally.
type Rm struct {
    args    *Args
}

// NewRmCommand returns an instance of the command. It takes os arguments and
// parses them to create the command which can be executed.
func NewRmCommand(rawArgs *[]string) (*Rm, error) {
    args := make(Args)
    if len(*rawArgs) > 2 {
        args["command"] = (*rawArgs)[1]
        args["container"] = (*rawArgs)[2]
    } else {
        return nil, errors.New("rm: No enough arguments")
    }
    return &Rm{ args: &args }, nil
}

// Name returns the name of the command.
func (cmd *Rm) Name() string {
    return "rm"
}

// Help returns the help information about the command.
func (cmd *Rm) Help() string {
    return "rm [container id]"
}

// Execute runs the command and returns error upon failure.
func (cmd *Rm) Execute() error {
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

    for i := 2; i < len(os.Args); i++ {
        Logger.Info("Iterating over list of containers")
        deleted := false
        for _, container := range containers {
            if container.ID() == os.Args[i] {
                Logger.Info("Found container ", container.ID())
                deleted = true
                if err := container.Delete(containerd.Context); err != nil {
                    fmt.Println("Failed to delete container")
                    return err
                }
                break
            }
        }

        if deleted {
            Logger.Info("Deleted container")
            fmt.Println(os.Args[i])
        } else {
            Logger.Info("Container not found")
            return errors.New("Container does not exist")
        }
    }

    return nil
}
