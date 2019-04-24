package bunker

import (
    "os"
    "fmt"
    "errors"
    containerdlib "github.com/containerd/containerd"
    "github.com/containerd/containerd/oci"
    "github.com/containerd/containerd/cio"
)

// Run implements Command interface. It takes care of running a task in a new
// or existing container.
type Run struct {
    args    *Args
}

// NewRunCommand returns an instance of the command. It takes os arguments and
// parses them to create the command which can be executed.
func NewRunCommand(rawArgs *[]string) (*Run, error) {
    args := make(Args)
    if len(*rawArgs) > 2 {
        args["command"] = (*rawArgs)[1]
        args["image"] = (*rawArgs)[2]
    } else {
        return nil, errors.New("run: No enough arguments")
    }
    return &Run{ args: &args }, nil
}

// Name returns the name of the command.
func (cmd *Run) Name() string {
    return "run"
}

// Help returns the help information about the command.
func (cmd *Run) Help() string {
    return "run [image] [command] [command args]"
}

// Execute runs the command and returns error upon failure.
func (cmd *Run) Execute() error {
    containerd, err := NewContainerd()
    if err != nil {
        Logger.Error(err)
        return err
    }

    imageName := Util.ImageNameToRegistryURL(cmd.args.Value("image"))

    if imageName == "" {
        err = errors.New("No image to run from")
        Logger.Warning(err)
        return nil
    }

    fmt.Println("Getting image", imageName)
    image, err := containerd.Client.Pull(containerd.Context, imageName, containerdlib.WithPullUnpack)
    if err != nil {
        return err
    }

    id := Util.NewRandomName()

    fmt.Println("Creating new container", id)
    container, err := containerd.Client.NewContainer(containerd.Context, id,
            containerdlib.WithImage(image),
            containerdlib.WithNewSnapshot(id, image),
            containerdlib.WithNewSpec(oci.WithImageConfig(image)))
    if err != nil {
        Logger.Error(err)
        return nil
    }

    task, err := container.NewTask(containerd.Context,
            cio.NewCreator(cio.WithStreams(os.Stdin, os.Stdout, os.Stderr)))
    if err != nil {
        return err
    }
    defer task.Delete(containerd.Context)

    task.Exec(containerd.Context, "/bin/bash", nil,
    cio.NewCreator(cio.WithStreams(os.Stdin, os.Stdout, os.Stderr)))

    task.Start(containerd.Context)

    return nil
}
