package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    fmt.Println("String:", random(10))
}

func random (x int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return rand.Intn(x)
}
