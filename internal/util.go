package bunker

import(
    "fmt"
    "time"
    "github.com/goombaio/namegenerator"
)

type util struct {
}

var Util util

func (u util) NewRandomName() string {
    seed := time.Now().UTC().UnixNano()
    return namegenerator.NewNameGenerator(seed).Generate()
}

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

func (u util) ImageNameToRegistryURL(name string) string {
    // TODO
    return name
}
