package lib

import "fmt"
import "./types"

func PullImage(image types.Image) error {

    // this is to be implemented using containerd
    fmt.Printf("Pulling image %v...\n", image.Name)

    return nil
}
