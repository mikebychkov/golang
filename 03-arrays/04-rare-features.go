package main

import (
    "fmt"
)

func main() {

    rates := [3]float64 {
        0: 0.5,
        1: 2.5,
        2: 1.5,
    }
    rates2 := [3]float64 {
        2: 1.5,
        1: 2.5,
        0: 0.5,
    }
    rates3 := [3]float64 {
        2: 1.5,
    }
    rates4 := [...]float64 {
        5: 1.5,
    }
    rates5 := [...]float64 {
        5: 1.5,
        2.5,
        1: 0.5,
    }
    const (
        AB = iota
        CD
        VW
    )
    rates6 := [...]float64 {
        AB: 1.5,
        CD: 2.5,
        VW: 0.5,
    }

    fmt.Printf("%#v %#v %#v %#v %#v %#v\n", rates, rates2, rates3, rates4, rates5, rates6)
}
