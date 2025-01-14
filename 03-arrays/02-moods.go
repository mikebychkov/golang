package main

import (
	"fmt"
	"math/rand"
    "os"
)

func main() {

    name := getArg(1)
    mood := getMood()
    fmt.Printf("%s feels %s\n", name, mood)
}

func getMood() string {

    moods := [...]string {
        "terrible",
        "bad",
        "sad",
        "good",
        "happy",
        "awesome",
    }
    
    i := rand.Intn(len(moods))

    return moods[i]
}

func getArg(num int) string {
    if len(os.Args) >= num + 1 {
        return os.Args[num]
    }
    return ""
}
