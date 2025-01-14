package main

import (
    "fmt"
)

func main() {

    var magazines [5]string
    magazines[0] = "Leads health"
    var published [len(magazines)]bool
    published[0] = true
    fmt.Printf("%#v %#v\n", magazines, published)

    // SHORT DECLARATION
    books := [4]string {
        "Book #1",
        "Book #2",
        "Book #3",
        "Book #4",
    }
    for _, v := range books {
        fmt.Println(v)
    }

    // ELLIPSIS OPERATOR - SIZE = INIT VALUES NUMBER
    books2 := [...]string {
        "Book #1",
        "Book #2",
        "Book #3",
    }
    for _, v := range books2 {
        fmt.Println(v)
    }
}
