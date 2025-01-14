package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type ClockNum [5][3]bool

const (
    drawSymbol = string(9608)
)

func main() {

    for {

        clearWindow()

        h := time.Now().Hour()
        m := time.Now().Minute()
        s := time.Now().Second()

        strTime := fmt.Sprintf("%02d%02d%02d", h, m, s)
        arrTime := strings.Split(strTime, "")

        var tt [6]ClockNum
        for i, v := range arrTime {
            num,_ := strconv.Atoi(v)
            tt[i] = getNum(num)
        }

        drawTime(tt, s)

        time.Sleep(time.Second)
    }
}

func clearWindow() {

    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func drawTime(time [6]ClockNum, s int) {

    seps := " "
    sepl := "     "
    sepd := "  " + drawSymbol + "  "

    for row := 0; row < len(time[0]); row++ {

        sep := sepl
        if row % 2 != 0 && s % 2 == 0 {
            sep = sepd
        }

        draw(time[0][row], seps)
        draw(time[1][row], sep)
        draw(time[2][row], seps)
        draw(time[3][row], sep)
        draw(time[4][row], seps)
        draw(time[5][row], seps)

        fmt.Println()
    }
}

func draw(arr [3]bool, sep string) {

    for _, v := range arr {
        if v {
            fmt.Printf(drawSymbol)
        } else {
            fmt.Printf(" ")
        }
    }
    fmt.Printf(sep)
}

func getNum(num int) ClockNum {

    switch num {
        default: return get0()
        case 1: return get1()
        case 2: return get2()
        case 3: return get3()
        case 4: return get4()
        case 5: return get5()
        case 6: return get6()
        case 7: return get7()
        case 8: return get8()
        case 9: return get9()
    }
}

func get0() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{true,false,true,},
        [3]bool{true,false,true,},
        [3]bool{true,false,true,},
        [3]bool{true,true,true,},
    }
}

func get1() ClockNum {

    return ClockNum {
        [3]bool{true,true,false,},
        [3]bool{false,true,false,},
        [3]bool{false,true,false,},
        [3]bool{false,true,false,},
        [3]bool{true,true,true,},
    }
}

func get2() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{false,false,true,},
        [3]bool{true,true,true,},
        [3]bool{true,false,false,},
        [3]bool{true,true,true,},
    }
}

func get3() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{false,false,true,},
        [3]bool{true,true,true,},
        [3]bool{false,false,true,},
        [3]bool{true,true,true,},
    }
}

func get4() ClockNum {

    return ClockNum {
        [3]bool{true,false,true,},
        [3]bool{true,false,true,},
        [3]bool{true,true,true,},
        [3]bool{false,false,true,},
        [3]bool{false,false,true,},
    }
}

func get5() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{true,false,false,},
        [3]bool{true,true,true,},
        [3]bool{false,false,true,},
        [3]bool{true,true,true,},
    }
}

func get6() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{true,false,false,},
        [3]bool{true,true,true,},
        [3]bool{true,false,true,},
        [3]bool{true,true,true,},
    }
}

func get7() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{false,false,true,},
        [3]bool{false,false,true,},
        [3]bool{false,false,true,},
        [3]bool{false,false,true,},
    }
}

func get8() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{true,false,true,},
        [3]bool{true,true,true,},
        [3]bool{true,false,true,},
        [3]bool{true,true,true,},
    }
}

func get9() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{true,false,true,},
        [3]bool{true,true,true,},
        [3]bool{false,false,true,},
        [3]bool{true,true,true,},
    }
}

