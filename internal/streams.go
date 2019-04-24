package lib

import (
    "os"
)

type Streams struct {
    In      os.File
    Out     os.File
    Err     os.File
}
