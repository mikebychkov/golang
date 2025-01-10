package main

import ("fmt"; "unicode/utf8")

func main() {

    // STRING LITERAL
    str1 := "<html>\n\t<body>\"Hello\"</body>\n</html>"

    // RAW STRING LITERAL (NO NEED TO USE ESCAPE CHARS)
    str2 := `
    <html>
        <body>"Hello"</body>
    <html>
    `

    fmt.Println(str1)
    fmt.Println(str2)

    // STRING LENGTH

    // len() - for strings byte size

    fmt.Println("Mike", len("Mike"))
    fmt.Println("Миша", len("Миша"))

    fmt.Println("Mike", utf8.RuneCountInString("Mike"))
    fmt.Println("Миша", utf8.RuneCountInString("Миша"))

}
