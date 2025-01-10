package main

import ("fmt")

func main() {

    var number int
    var str string

    number = 33

    str = string(number)

    fmt.Println(number, str)

    // --

    speed := 100 // int
    force := 2.5 // float64

    rsl := speed * int(force)
    fmt.Println(rsl)

    rsl1 := float64(speed) * force
    fmt.Println(rsl1)
}
