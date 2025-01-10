package main

import ("fmt"; "os"; "strconv")

func main() {

    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Converted value:", n)
    }

    // SIMPLE STATEMENT
    if n, err := strconv.Atoi(os.Args[1]); err == nil {
        fmt.Println("There is no error", n)
    }
}
