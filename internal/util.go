package bunker

import(
    "fmt"
    "time"
    "github.com/goombaio/namegenerator"
)

// util is a wrapper struct for util functions.
type util struct {
}

// Util is a global instance of util struct.
var Util util

// NewRandomName generates and returns a random string name.
func (u util) NewRandomName() string {
    seed := time.Now().UTC().UnixNano()
    return namegenerator.NewNameGenerator(seed).Generate()
}

// ByteToString converts a given number in bytes and return a string with the
// number in human readable format.
func (u util) ByteToString(size int64) string {
    unit := []string{"", "B", "KB", "MB", "GB", "TB"}

    scale := 0
    res := size
    temp := size

    for {
        if temp == 0 {
            break
        }
        res = temp
        temp /= 1000
        scale += 1
    }

    return fmt.Sprintf("%v%v", res, unit[scale])
}

// ImageNameToRegistryURL converts the name of an image into the URL for the
// docker registry which can be used to pull a container image.
func (u util) ImageNameToRegistryURL(name string) string {
    // TODO
    return name
}
