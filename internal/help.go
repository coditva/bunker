package bunker

import "fmt"

var helpText string =
`Usage:
pull        Pull new image
push        Update image
`

func PrintHelp(prefix error) {
    fmt.Println(prefix, "\n")
    fmt.Print(helpText)
}
