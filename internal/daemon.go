package bunker

import (
    "os"
    "fmt"
    "os/exec"
    "syscall"
)

type Daemon struct {
}

func NewDaemon() *Daemon {
    return &Daemon{}
}

func (daemon *Daemon) Execute() error {
    InitLogger("bunker", "/tmp/bunkerd.log")

    if len(os.Args) < 2 {
        daemon.Help()
        os.Exit(1)
    } else if os.Args[1] == "start" {
        daemon.Start()
    } else if os.Args[1] == "stop" {
        daemon.Stop()
    } else if os.Args[1] == "enable" {
        daemon.Enable()
    } else if os.Args[1] == "disable" {
        daemon.Disable()
    } else if os.Args[1] == "restart" {
        daemon.Stop()
        daemon.Start()
    } else if os.Args[1] == "status" {
        daemon.Status()
    } else {
        daemon.Help()
        os.Exit(1)
    }
    return nil
}

func (daemon *Daemon) Start() error {
    Logger.Info("Starting bunkerd")

    binary, err := exec.LookPath("containerd")
    if err != nil {
        Logger.Error("Could not find containerd")
        os.Exit(1)
    }
    args := []string{"containerd", "--address", ContainerdSocketPath,
            "--log-level", "fatal"}
    env := os.Environ()
    procAttr := syscall.ProcAttr{
        Dir: "/tmp",
        Env: env,
    }

    if pid, err := syscall.ForkExec(binary, args, &procAttr); err != nil {
        return err
    } else {
        Logger.Info("Started containerd daemon PID: ", pid)
    }

    return nil
}

func (daemon *Daemon) Stop() error {
    return nil
}

func (daemon *Daemon) Enable() error {
    return nil
}

func (daemon *Daemon) Disable() error {
    return nil
}

func (daemon *Daemon) Status() error {
    return nil
}

func (daemon *Daemon) Help() {
    const helpText =
            "Usage: bunkerd start | stop | restart | enable | disable | status\n"
    fmt.Printf(helpText)
}
