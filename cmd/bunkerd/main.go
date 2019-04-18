package main

import "fmt"
import (
    "os"
    "os/exec"
    "syscall"
)
import (
    lib "github.com/coditva/bunker/internal"
    rpc "github.com/coditva/bunker/internal/rpc"
    api "github.com/coditva/bunker/internal/api"
)

var helpText =
"Usage: bunkerd start | stop | restart | enable | disable | status\n"

func start() error {
    binary, err := exec.LookPath("containerd")
    if err != nil {
        fmt.Println("Could not find containerd")
        os.Exit(1)
    }
    args := []string{"containerd", "--address", lib.ContainerdSocketPath,
            "--log-level", "fatal"}
    env := os.Environ()
    procAttr := syscall.ProcAttr{
        Dir: "/tmp",
        Env: env,
    }

    if pid, err := syscall.ForkExec(binary, args, &procAttr); err != nil {
        return err
    } else {
        fmt.Println("Daemon Process: ", pid)
    }

    server := rpc.NewServer("/tmp/rpc.sock")
    if err := server.Serve(api.New()); err != nil {
        return err
    }
    defer server.Close()

    return nil
}

func stop() error {
    return nil
}

func enable() error {
    return nil
}

func disable() error {
    return nil
}

func status() error {
    return nil
}

func printHelp() {
    fmt.Printf(helpText)
}

func main() {
    if len(os.Args) < 2 {
        printHelp()
        os.Exit(1)
    } else if os.Args[1] == "start" {
        start()
    } else if os.Args[1] == "stop" {
        stop()
    } else if os.Args[1] == "enable" {
        enable()
    } else if os.Args[1] == "disable" {
        disable()
    } else if os.Args[1] == "restart" {
        stop()
        start()
    } else if os.Args[1] == "status" {
        status()
    } else {
        printHelp()
        os.Exit(1)
    }
}
