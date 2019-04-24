package bunker

import (
    "context"

    "github.com/containerd/containerd"
    "github.com/containerd/containerd/namespaces"
)

// Containerd is a wrapper over containerd client library. It is takes care of
// storing a containerd client instance which is used by the CLI.
type Containerd struct {

    // Client is the initialied containerd client.
    Client  *containerd.Client

    // Context is the containerd context with a separate namespace for bunker.
    Context context.Context
}

// NewContainerd return a new initialized instance of containerd client
func NewContainerd() (*Containerd, error) {
    client := new(Containerd)

    Logger.Info("Connecting to containerd")
    if c, err := containerd.New(ContainerdSocketPath); err != nil {
        return nil, err
    } else {
        Logger.Info("Connected to containerd as client")
        client.Client = c
    }

    Logger.Info("Creating containerd client context")
    client.Context = namespaces.WithNamespace(context.Background(),
            ContainerdNamespace)

    return client, nil
}
