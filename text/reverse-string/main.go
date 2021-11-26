package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Reverse(input string) (output string) {
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]

	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}

	output = string(rune)
	return
}

func main() {
	fmt.Print("Enter a string to reverse: ")
	in := bufio.NewReader(os.Stdin)
	if line, err := in.ReadString('\n'); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(Reverse(line))
	}
}
