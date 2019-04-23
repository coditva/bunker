package api

import (
    "fmt"
    "github.com/containerd/containerd"
    "github.com/containerd/containerd/oci"

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
    *reply = fmt.Sprintf("Running command %v on image %v in container ID %v",
            runCommand, imageName, container.ID())

    lib.Logger.Info(*reply)

    return nil
}
