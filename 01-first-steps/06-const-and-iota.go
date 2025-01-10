package main

import ("fmt")

func main() {
    example1()
    example2()
    example3()
    example4()
    fmt.Printf("%s - %s - %d - %T - %v - %[1]v - %v\n", "qwas", 12, 13, 14, 15)
    fmt.Printf("%T - %.2f\n", 12.123, 12.123)
}

func example1() {
    const (
        monday = 0
        tuesday = 1
        wednesday = 2
        thursday = 3
        friday = 4
        saturday = 5
        sunday = 6
    )
    fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}

func example2() {
    const (
        monday = iota
        tuesday
        wednesday
        thursday
        friday
        saturday
        sunday
    )
    fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}

func example3() {
    const (
        monday = iota + 1
        tuesday
        wednesday
        thursday
        friday
        saturday
        sunday
    )
    fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}

func example4() {
    // EST is -5
    // MST is -7
    // PST is -8
    const (
        EST = -(5 + iota)
        _
        MST
        PST
    )
    fmt.Println(EST, MST, PST)
}
