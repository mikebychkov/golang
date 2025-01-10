package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {

    guess := 10

    rand.NewSource(time.Now().UnixNano())

    for i := 0; i < 10; i++ {
        n := rand.Intn(guess + 1)
        fmt.Printf("%d ", n)
    }
}
