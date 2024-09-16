package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("🧙🏻 Guessing Game!")
	fmt.Println("🫥 A random number will be choosed")
	fmt.Println("😄 You've 10 chances guess the right number")
	fmt.Println("😵 It's a integer between 0 and 100")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	guesses := [10]int64{}

	for i := range guesses {
		fmt.Println("\n🧐 What is your guess?")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInt, err := strconv.ParseInt(guess, 10, 64)
		if err != nil {
			fmt.Println("😡 Your guess must be a integer number")
			return
		}

		switch {
		case guessInt > x:
			fmt.Printf("😫 You missed! The right number is lower ⬇️ than %d. You've %d Try again!", guessInt, 10-(i+1))
		case guessInt < x:
			fmt.Printf("😫 You missed! The right number is higher ⬆️ than %d. You've %d Try again!", guessInt, 10-(i+1))
		case guessInt == x:
			fmt.Println("\n\n😃 You're goddamn righ!")
			fmt.Printf("👏 You took %d guesses to find the number", i+1)
			return
		}
	}

	fmt.Println("\n\n😭 You loose!")
}
