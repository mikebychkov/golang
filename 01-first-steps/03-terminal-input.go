package main

import ("fmt"; "os")

func main() {

    fmt.Println(os.Args[0])
    fmt.Println(os.Args[1])
    fmt.Println(os.Args[2])

    fmt.Println("%#v\n", os.Args, len(os.Args))
}
