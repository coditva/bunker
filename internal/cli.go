package lib

import (
    "os"
    "fmt"
    "errors"
)

type CLI struct {
}

func NewCLI() *CLI {
    return &CLI{}
}


func (cli *CLI) Execute() error {
    var command Command
    var err error

    InitLogger("bunker", "/tmp/bunker.log")

    // empty command
    if len(os.Args) < 2 {
        Logger.Info("No command name given")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "pull":
        if command, err = NewPullCommand(&os.Args); err != nil {
            Logger.Info(err)
            fmt.Println(err)
            return nil
        }

    case "run":
        if command, err = NewRunCommand(&os.Args); err != nil {
            Logger.Info(err)
            fmt.Println(err)
            return nil
        }

    case "images":
        if command, err = NewImagesCommand(&os.Args); err != nil {
            Logger.Info(err)
            fmt.Println(err)
            return nil
        }

    case "containers":
        if command, err = NewContainersCommand(&os.Args); err != nil {
            Logger.Info(err)
            fmt.Println(err)
            return nil
        }

    default:
        err := errors.New("Unknown command")
        Logger.Info(err)
        PrintHelp(err)
        return err
    }

    if err := command.Execute(); err != nil {
        Logger.Error(err)
        os.Exit(1)
    }

    return nil
}
