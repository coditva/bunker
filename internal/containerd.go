package lib

import (
    "fmt"
    "context"

    "github.com/containerd/containerd"
    "github.com/containerd/containerd/namespaces"
)

type Containerd struct {
    Client  *containerd.Client
    Context context.Context
    Ns      context.Context
}

var ContainerdClient *Containerd


func InitContainerd() error {
    ContainerdClient = new(Containerd)
    var err error

    Logger.Info("Creating new client")
    ContainerdClient.Client, err = containerd.New(ContainerdSocketPath)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Could not connect to containerd")
        return err
    }
    //defer ContainerdClient.Client.Close()

    Logger.Info("Creating new containerd client context")
    ContainerdClient.Context = context.Background()
    ContainerdClient.Ns = namespaces.WithNamespace(ContainerdClient.Context, "bunker")
    return nil
}
