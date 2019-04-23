package api

import (
    "fmt"
    "github.com/containerd/containerd"
    "github.com/containerd/containerd/oci"
    "github.com/containerd/containerd/cio"

    lib "github.com/coditva/bunker/internal"
    util "github.com/coditva/bunker/internal/util"
    types "github.com/coditva/bunker/internal/types"
)

func (api Api) Run(args *types.Args, reply *string) error {
    imageName := (*args)[0]
    runCommand := (*args)[1]

    if imageName == "" {
        *reply = "No image to run from"
        lib.Logger.Warning(*reply)
        return nil
    }

    if runCommand == "" {
        *reply = "No command given to run"
        lib.Logger.Warning(*reply)
        return nil
    }

    image, err := lib.ContainerdClient.Client.Pull(lib.ContainerdClient.Ns, imageName, containerd.WithPullUnpack)
    if err != nil {
        lib.Logger.Error(err)
        return err
    }

    id := util.NewRandomName()
    container, err := lib.ContainerdClient.Client.NewContainer(lib.ContainerdClient.Ns, id,
            containerd.WithNewSnapshot(id, image),
            containerd.WithNewSpec(oci.WithImageConfig(image)))
    if err != nil {
        lib.Logger.Error(err)
        return nil
    }

    task, err := container.NewTask(lib.ContainerdClient.Ns, cio.LogFile("/tmp/task.log"))
    if err != nil {
        *reply = "Could not create new task"
        lib.Logger.Error(err)
    }
    defer task.Delete(lib.ContainerdClient.Ns)
    task.Start(lib.ContainerdClient.Ns)

    *reply = fmt.Sprintf("Running command %v on image %v in container ID %v",
            runCommand, imageName, container.ID())

    lib.Logger.Info(*reply)

    return nil
}
