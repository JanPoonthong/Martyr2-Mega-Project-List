package main

import (
    "fmt"
)

func Reverse(str string) (output string) {
    for _, char := range str {
        output = string(char) + output
    }
    return
}

func main() {
    var userInput string
    fmt.Print("Enter a word: ")
    fmt.Scanln(&userInput)
    fmt.Println(Reverse(userInput))
}
