package main

import "fmt"
import (
    "os"
    "os/exec"
    "syscall"
)
import "github.com/coditva/bunker/internal"

var helpText =
`Usage:
start
stop
restart
enable
disable
`

func start() error {
    binary, lookErr := exec.LookPath("containerd")
    if lookErr != nil {
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

    pid, err := syscall.ForkExec(binary, args, &procAttr)
    fmt.Println("Daemon Process: ", pid)
    return err
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
    } else {
        printHelp()
        os.Exit(1)
    }
}
