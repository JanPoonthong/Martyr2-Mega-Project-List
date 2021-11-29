package main

import (
	"fmt"
	"strconv"
)

func Piglatin(userInput string) (output string) {
	firstChar := userInput[0]
	fmt.Println(string(firstChar))
	return
}

func Intchecker(userInput string) string {
	if _, err := strconv.Atoi(userInput); err == nil {
		return fmt.Sprintf("%q looks like a number.\n", userInput)
	}
        return nil
}

func main() {
	var userInput string
	fmt.Print("Enter a string(not int): ")
	fmt.Scan(&userInput)

	Intchecker(userInput)
	Piglatin(userInput)
}
