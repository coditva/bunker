package lib

import (
    "os"
    "fmt"
    "errors"

    types "github.com/coditva/bunker/internal/types"
)

func InitCLI() error {
    InitLogger("bunker", "/tmp/bunker.log")

    if err := InitContainerd(); err != nil {
        return err
    }
    return nil
}

func ExecuteCLI() error {
    var reply string

    command, err := ParseArgs(os.Args)
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


func ParseArgs(args []string) (*types.Command, error) {
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
