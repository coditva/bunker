package types

type Command struct {
    Name    string
    Method  string
    Args    Args
    ArgsLen int
}

func (command *Command) AddArg(arg string) {
    command.Args[command.ArgsLen] = arg
    command.ArgsLen += 1
}
