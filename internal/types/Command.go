package types

import (
    "os"
)

type Streams struct {
    In      os.File
    Out     os.File
    Err     os.File
}

type Command struct {
    Name    string
    Method  string
    Args    Args
    ArgsLen int
    Streams  Streams
}

func (command *Command) AddArg(arg string) {
    command.Args[command.ArgsLen] = arg
    command.ArgsLen += 1
}
