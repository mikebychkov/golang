package main

import ("fmt"; "os"; "strconv"; "time")

func main() {

    fmt.Println("Current time:", time.Now())
    fmt.Println("Day: ", time.Now().Day())
    fmt.Println("Day of year: ", time.Now().YearDay())
    y, m, d := time.Now().Date()
    fmt.Printf("Year: %v, Month: %v, Day: %v", y, m, d)
    fmt.Println("\n")

    rsl, err := isInputEven()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    if rsl {
        fmt.Println("Number is even")
    } else {
        fmt.Println("Number is odd")
    }
}

func isInputEven() (bool, error) {

    if len(os.Args) != 2 {
        return false, fmt.Errorf("Incorect input")
    }

    param := os.Args[1]

    num, err := strconv.Atoi(param)
    if err != nil {
        return false, fmt.Errorf("Parse error: %v", err)
    }

    fmt.Println("Input number:", num)

    return num % 2 == 0, nil
}
