package lib

type Command interface {
    Name()                  string
    Execute()               error
    Help()                  string
}
