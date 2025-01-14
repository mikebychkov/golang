package main

import (
    "fmt"
)

func main() {

    arr := [3]int{3,4,7}

    type bookcase [3]int
    a := bookcase{3,4,7}

    fmt.Printf("%#v == %#v is %v\n", arr, a, arr == a)
    
    type cabinet [3]int
    b := cabinet{3,4,7}

    fmt.Printf("%#v == %#v is %v\n", a, b, a == bookcase(b))
}
