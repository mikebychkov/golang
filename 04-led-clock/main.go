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

