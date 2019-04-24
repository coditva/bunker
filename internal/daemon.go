package bunker

import (
    "os"
    "fmt"
    "strconv"
    "os/exec"
    "syscall"
    "io/ioutil"
)

// DaemonPID is the PID of the running bunkerd daemon.
var DaemonPID = 0

// DaemonIsRunning is a boolean value true if daemon is running.
var DaemonIsRunning = false


// Daemon is the starting point for the bunkerd daemon. It is responsible for
// starting a new containerd instance.
type Daemon struct {
}

// NewDaemon returns an instance of Daemon.
func NewDaemon() *Daemon {
    return &Daemon{}
}

// Start execution of bunkerd daemon.
func (daemon *Daemon) Execute() error {
    InitLogger("bunker", "/tmp/bunkerd.log")

    if daemonPID, err := ioutil.ReadFile(ContainerdDaemonPIDFile); err != nil {
        DaemonIsRunning = false
        Logger.Info("Daemon is not running")
    } else {
        DaemonIsRunning = true
        if n, err := strconv.Atoi(string(daemonPID)); err == nil {
            DaemonPID = n
        } else {
            Logger.Error(err)
        }
        Logger.Info("Daemon is running")
    }

    if len(os.Args) < 2 {
        daemon.Help()
        Logger.Info("No arguments given")
        os.Exit(1)

    } else if os.Args[1] == "start" {
        Logger.Info("Starting daemon")
        daemon.Start()

    } else if os.Args[1] == "stop" {
        Logger.Info("Stopping daemon")
        daemon.Stop()

    } else if os.Args[1] == "enable" {
        Logger.Info("Enabling daemon")
        daemon.Enable()

    } else if os.Args[1] == "disable" {
        Logger.Info("Disabling daemon")
        daemon.Disable()

    } else if os.Args[1] == "restart" {
        Logger.Info("Restarting daemon")
        daemon.Stop()
        daemon.Start()

    } else if os.Args[1] == "status" {
        Logger.Info("Printing status information")
        daemon.Status()

    } else {
        daemon.Help()
        os.Exit(1)
    }
    return nil
}

// Start starts the daemon if not already running.
func (daemon *Daemon) Start() error {
    if DaemonIsRunning {
        fmt.Println("Daemon already running")
        return nil
    }

    Logger.Info("Starting bunkerd")

    // TODO: Check if daemon is already running and abort if it is.

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
        ioutil.WriteFile(ContainerdDaemonPIDFile, []byte(strconv.Itoa(pid)),
                0660)
        DaemonIsRunning = true
    }

    fmt.Println("Started daemon")
    return nil
}

// Stop stops the running daemon.
func (daemon *Daemon) Stop() error {
    if DaemonIsRunning == false {
        fmt.Println("Daemon not running")
        return nil
    }

    Logger.Info("Killing daemon")
    if daemon, err := os.FindProcess(DaemonPID); err != nil {
        Logger.Error(err)
    } else {
        daemon.Kill()
        DaemonIsRunning = false
    }

    os.Remove(ContainerdDaemonPIDFile)
    fmt.Println("Stopped daemon")
    return nil
}

// Enable enables the starting of daemon on startup.
func (daemon *Daemon) Enable() error {
    return nil
}

// Disable disables the starting of daemon on startup.
func (daemon *Daemon) Disable() error {
    return nil
}

// Status prints the status of the daemon.
func (daemon *Daemon) Status() error {
    if DaemonIsRunning {
        fmt.Println("Daemon is running")
    } else {
        fmt.Println("Daemon is not running")
    }
    return nil
}

// Help prints the help information on how to use bunkerd daemon.
func (daemon *Daemon) Help() {
    const helpText =
            "Usage: bunkerd start | stop | restart | enable | disable | status\n"
    fmt.Printf(helpText)
}
