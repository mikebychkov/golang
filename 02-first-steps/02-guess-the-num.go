package main

import (
    "fmt"
    "math/rand"
    "time"
    "strconv"
    "os"
)

const maxTurns = 5
const maxGuess = 10
const winRate = 100
const bonusRate = 50
const doubleGuessDivider = 3

func main() {
    
    winRate := spinTheGame()
    fmt.Println("Your win rate is", winRate)
}

func spinTheGame() int {

    rand.NewSource(time.Now().UnixNano())

    guess := getArg(1)
    guessNum, err := strconv.Atoi(guess)
    if err != nil {
        fmt.Println("Ivalid input")
        return 0
    }
    if guessNum < 0 || guessNum > maxGuess {
        fmt.Println("Ivalid input")
        return 0
    }

    guess2 := getArg(2)
    guess2Num, err := strconv.Atoi(guess2)
    guess2Set := true
    if err != nil {
        guess2Set = false
    }

    for i := 0; i < maxTurns; i++ {
        n := rand.Intn(maxGuess + 1)
        fmt.Printf("%d ", n)
        if isWin, rate := checkIfWin(guessNum, n, i, guess2Set); isWin {
            return rate
        }
        if guess2Set {
            if isWin, rate := checkIfWin(guess2Num, n, i, guess2Set); isWin {
                return rate
            }
        }
    }
    fmt.Println("You LOSE!!!")
    return 0
}

func checkIfWin(guess int, n int, i int, doubleGuess bool) (bool, int) {

    if guess == n {
        rate := winRate
        fmt.Println("You WIN!!!")
        if i == 0 {
            rate += bonusRate
            fmt.Println("First turn winner BONUS!!!")
        }
        if doubleGuess {
            rate /= doubleGuessDivider
        }
        return true, rate
    }
    return false, 0
}

func getArg(num int) string {
    if len(os.Args) >= num + 1 {
        return os.Args[num]
    }
    return ""
}
