package lib

import "fmt"
import "errors"
import "github.com/coditva/bunker/internal/types"

func PullImage(image types.Image) error {

    // check for image locally
    err := pullLocalImage(image)
    if err != nil {
        fmt.Println(err)
    } else {
        return nil
    }

    // pull image from registry
    // this is to be implemented using containerd
    fmt.Printf("Pulling image %v from registry\n", image.Name)

    return nil
}

func pullLocalImage(image types.Image) error {
    return errors.New(fmt.Sprintf("Image \"%v\" not found locally", image.Name))
}
