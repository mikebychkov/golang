package main

import (
	"fmt"
	"strings"
)

func main() {

    for i := 1; i < 100; i++ {
        fmt.Println(i)
    }

    k := 0
    for ;k < 20; {
        fmt.Println(k)
        k++
    }

    for k > 0 {
        fmt.Println(k)
        k--
    }

    // infinite loop
    n := 5
    for {
        if n > 0 {
            break // continue - as expected
        }
        n--
    }
    
    str := "Hello world"
    arr := strings.Split(str, "")
    for i := 0; i < len(str); i++ {
        fmt.Println(arr[i])
    }

    arr2 := strings.Fields("Hel      lo   wo   rld")
    for i := 0; i < len(arr2); i++ {
        fmt.Println(arr2[i])
    }

    for i, v := range arr2 {
        fmt.Println(i, v)
    }

    for i, v := range arr2 {
        fmt.Println(i, v)
    }

    for i, v := range arr2[2:] {
        fmt.Println(i, v)
    }

    for i, v := range "whaat" {
        fmt.Println(i, string(v))
    }

    fmt.Println('w', int('w'), rune(int('w')), string('w'))
}
