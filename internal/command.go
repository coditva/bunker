package bunker

// Command is an interface type for all API commands in bunker.
type Command interface {

    // Returns the name of the command.
    Name()                  string

    // Executes the command and returns error when it fails.
    Execute()               error

    // Returns help information about the command as string.
    Help()                  string
}
