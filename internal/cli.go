package bunker

import (
    "os"
    "fmt"
    "errors"
)

// Text that is displayed when an unknown command or help command is issued
const helpText =
`Usage:
pull        Pull new image
run         Run a command in a new containers for the image
images      List images
containers  List containers
`

// CLI is the starting point for the bunker cli client. It takes care of
// argument parsing and calling of correct API methods.
type CLI struct {
}

// Create a new CLI instance.
func NewCLI() *CLI {
    return &CLI{}
}


// Start execution of the CLI client.
func (cli *CLI) Execute() error {
    var command Command
    var err error

    InitLogger("bunker", "/tmp/bunker.log")

    if len(os.Args) < 2 {
        Logger.Info("Empty command")
        cli.printHelp("Need command to execute")
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

    case "image":
        if command, err = NewImageCommand(&os.Args); err != nil {
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
        cli.printHelp(err)
        return err
    }

    if err := command.Execute(); err != nil {
        Logger.Error(err)
        os.Exit(1)
    }

    return nil
}

// Prints help information.
func (cli *CLI) printHelp(prefix interface{}) {
    fmt.Println(prefix, "\n")
    fmt.Print(helpText)
}
