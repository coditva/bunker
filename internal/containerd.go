package bunker

import (
    "context"

    "github.com/containerd/containerd"
    "github.com/containerd/containerd/namespaces"
)

type Containerd struct {
    Client  *containerd.Client
    Context context.Context
}

func NewContainerd() (*Containerd, error) {
    client := new(Containerd)

    Logger.Info("Connecting to containerd")
    if c, err := containerd.New(ContainerdSocketPath); err != nil {
        Logger.Error(err)
        return nil, err
    } else {
        Logger.Info("Connected to containerd as client")
        client.Client = c
    }
    //defer ContainerdClient.Client.Close()

    Logger.Info("Creating containerd client context")
    client.Context = namespaces.WithNamespace(context.Background(), "bunker")
    return client, nil
}
