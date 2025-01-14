package main

import (
	"fmt"
	"os"
	"strings"
)

const data = "qwas Mike cake love qwas STONE super paw PAtrol love deatH robots MIKE"

func main() {

    input := os.Args[1:]

    dataSlice := strings.Fields(data)

    main: for _, iv := range input {

        for i, dv := range dataSlice {

            if strings.ToLower(iv) == strings.ToLower(dv) {

                fmt.Printf("#%-2d: %s\n", i + 1, dv)
                break main
            }
        }
    }
}
