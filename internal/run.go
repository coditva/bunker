package bunker

import (
    "os"
    "fmt"
    "errors"
    containerdlib "github.com/containerd/containerd"
    "github.com/containerd/containerd/oci"
    "github.com/containerd/containerd/cio"
)

type Run struct {
    args    *Args
}

func NewRunCommand(rawArgs *[]string) (*Run, error) {
    args := make(Args)
    if len(*rawArgs) > 3 {
        args["command"] = (*rawArgs)[1]
        args["image"] = (*rawArgs)[2]
        args["binary"] = (*rawArgs)[3]
    } else {
        return nil, errors.New("run: No enough arguments")
    }
    return &Run{ args: &args }, nil
}

func (cmd *Run) Name() string {
    return "run"
}

func (cmd *Run) Help() string {
    return "run [image] [command] [command args]"
}

func (cmd *Run) Execute() error {
    containerd, err := NewContainerd()
    if err != nil {
        Logger.Error(err)
        return err
    }

    imageName := Util.ImageNameToRegistryURL(cmd.args.Value("image"))
    runCommand := cmd.args.Value("command")

    if imageName == "" {
        err = errors.New("No image to run from")
        Logger.Warning(err)
        return nil
    }

    if runCommand == "" {
        err = errors.New("No command given to run")
        Logger.Warning(err)
        return nil
    }

    image, err := containerd.Client.Pull(containerd.Context, imageName, containerdlib.WithPullUnpack)
    if err != nil {
        return err
    }

    id := Util.NewRandomName()
    container, err := containerd.Client.NewContainer(containerd.Context, id,
            containerdlib.WithNewSnapshot(id, image),
            containerdlib.WithNewSpec(oci.WithImageConfig(image)))
    if err != nil {
        Logger.Error(err)
        return nil
    }

    //os.Mkdir(ContainerStreamBasePath, 0660)

    //containerOut, err := os.OpenFile(fmt.Sprintf("%v%v.out", ContainerStreamBasePath, id),
            //os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
    //if err != nil {
        //Logger.Error(err)
        //return err
    //}

    //containerIn, err := os.OpenFile(fmt.Sprintf("%v%v.in", ContainerStreamBasePath, id),
            //os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0660)
    //if err != nil {
        //Logger.Error(err)
        //return err
    //}

    //containerErr, err := os.OpenFile(fmt.Sprintf("%v%v.err", ContainerStreamBasePath, id),
            //os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
    //if err != nil {
        //Logger.Error(err)
        //return err
    //}

    task, err := container.NewTask(containerd.Context,
            cio.NewCreator(cio.WithStreams(os.Stdin, os.Stdout, os.Stderr)))
    if err != nil {
        return err
    }
    defer task.Delete(containerd.Context)

    Logger.Info(fmt.Sprintf("Running command %v on image %v in container ID %v",
            runCommand, imageName, container.ID()))
    task.Start(containerd.Context)

    return nil
}
