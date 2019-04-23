package lib

import (
    "os"
    "fmt"
    "errors"

    types "github.com/coditva/bunker/internal/types"
)

type CLI struct {
}

func NewCLI() *CLI {
    return &CLI{}
}


func (cli *CLI) Execute() error {
    var reply string
    InitLogger("bunker", "/tmp/bunker.log")

    if err := InitContainerd(); err != nil {
        Logger.Error(err)
        return err
    }

    command, err := cli.ParseArgs(os.Args)
    if err != nil {
        PrintHelp(err)
        os.Exit(1)
    }

    if err := command.Method(&command.Args, &reply); err != nil {
        Logger.Error(err)
        os.Exit(1)
    }
    fmt.Println(reply)

    return nil
}


func (cli *CLI) ParseArgs(args []string) (*types.Command, error) {
    var err error
    var command *types.Command

    Logger.Info("Parsing command line arguments")

    if len(args) < 2 {
        err = errors.New("No command")
        return nil, err
    }

    command, err = NewCommand(args[1])
    if err != nil {
        return nil, err
    }

    for i := 2; i < len(args); i++ {
        if args[i] != "" {
            command.AddArg(args[i])
        }
    }

    return command, nil
}

func NewCommand(name string) (*types.Command, error) {
    command := new(types.Command)
    command.ArgsLen = 0
    command.Name = name

    if name == "pull" {
        command.Method = Pull
    } else if name == "images" {
        command.Method = Images
    } else if name == "run" {
        command.Method = Run
    } else if name == "containers" {
        command.Method = Containers
    } else {
        return nil, errors.New(fmt.Sprintf("Unknown command %v", name))
    }

    return command, nil
}
