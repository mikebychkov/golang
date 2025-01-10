package main

import ("fmt"; "os"; "strings")

func main() {

    city := getArg(1)

    switch city {
        case "Paris", "Marsel": fmt.Println("France")
            // break - added automaticaly
        case "Tokio", "Kioto", "Osaka": fmt.Println("Japan"); break
        default: fmt.Println("Other")
    }

    numberType(122)
    numberType(81)
}

func numberType(i int) {

    fmt.Println(strings.Repeat("=", 50))
    fmt.Println("Number", i)

    switch { // for bool cases by default - true (so can be omitted)
        case i > 100: fmt.Println("big positive number"); fallthrough
        case i > 0: 
            fmt.Println("positive number")
            fallthrough
        default: fmt.Println("number")
    }
}

func getArg(num int) string {
    if len(os.Args) >= num + 1 {
        return os.Args[num]
    }
    return ""
}
