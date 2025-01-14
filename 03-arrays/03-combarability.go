package main

import (
    "fmt"
)

func main() {

    blue := [3]int{3,6,9}
    red := [3]int{3,6,9}
    yellow := [3]int{6,9,3}

    fmt.Printf("%#v and %#v are equal? - %v\n", blue, red, blue == red)
    fmt.Printf("%#v and %#v are equal? - %v\n", blue, yellow, blue == yellow)

    arr1 := [5]int{1,2,3}
    var arr2 [5]int
    // var arr3 [3]int
    arr4 := [5]int{4,4,4,4,4}

    fmt.Printf("%#v , %#v , %#v\n", arr1, arr2, arr4)
    arr2 = arr1
    arr4 = arr1
    // arr3 = arr1 // cannot be assigned
    fmt.Printf("%#v , %#v , %#v\n", arr1, arr2, arr4)
}
