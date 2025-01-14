package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

    args := os.Args[1:]

    err := validate(args)
    if err != nil {
        fmt.Println("Error occured:", err)
        return
    }

    fmt.Printf("%#v\n", args)

    for {
        swapped := false
        for i := 0; i < len(args) - 1; i++ {
            if compareNumericStrs(args[i], args[i + 1]) > 0 {
                t := args[i + 1]
                args[i + 1] = args[i]
                args[i] = t
                swapped = true
            }
        }
        if !swapped {
            break
        }
    }

    fmt.Printf("%#v\n", args)
}

func compareNumericStrs(a string, b string) int {
    n1,_ := strconv.Atoi(a) 
    n2,_ := strconv.Atoi(b) 
    if n1 > n2 {
        return 1
    } else if n1 < n2 {
        return -1
    } else {
        return 0
    }
}

func validate(a []string) (error) {
    
    if len(a) > 5 {
        return fmt.Errorf("Max 5 elements")
    }

    for _, v := range a {
        _, err := strconv.Atoi(v)
        if err != nil {
            return fmt.Errorf("Parsing error %s to integer", v)
        }
    }

    return nil
}
