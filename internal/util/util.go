package util

import(
    "time"
    "github.com/goombaio/namegenerator"
)

func NewRandomName() string {
    seed := time.Now().UTC().UnixNano()
    return namegenerator.NewNameGenerator(seed).Generate()
}
