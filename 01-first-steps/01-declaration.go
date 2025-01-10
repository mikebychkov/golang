package main

import ("fmt"; "path")

func main() {

    fullPath := "/home/golang/hello.world"

    var dir, file string

    dir, file = path.Split(fullPath)

    fmt.Println("dir:", dir)
    fmt.Println("file:", file)

    
    // BLANK IDENTIFIER

    var file1 string

    _, file1 = path.Split(fullPath)

    fmt.Println("file:", file1)


    // SHORT DECLARATION FORM

    _, file2 := path.Split(fullPath)
    fmt.Println("file:", file2)

    
    // GROUP DECLARATION

    var (
        color = "green"
        width = 100
        height = 50
    )
    fmt.Println("group vars:", color, width, height)
}
