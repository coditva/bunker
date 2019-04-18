package lib

import "fmt"
import "github.com/containerd/containerd"
import "os"

func ClientNew() {
    client, err := containerd.New(ContainerdSocketPath)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Could not connect to bunkerd")
        os.Exit(1)
    } else {
        defer client.Close()
    }
}
