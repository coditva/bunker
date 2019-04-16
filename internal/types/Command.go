package types

type CommandName int
const (
    CommandEmpty    CommandName = 0
    CommandUnknown  CommandName = 1
    CommandPull     CommandName = 2
    CommandPush     CommandName = 3
    CommandRun      CommandName = 4
    CommandPs       CommandName = 5
    CommandBuild    CommandName = 6
)

type Command struct {
    Name CommandName
}
