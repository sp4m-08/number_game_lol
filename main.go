package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	for {
		log.Printf("Welcome to this game!")
		log.Println("Please select your difficulty level or type 'quit' to exit: ")
		log.Println(" ")
		log.Println("1. Easy (5 chances)")
		log.Println("2. Medium (3 chances)")
		log.Println("3. Hard (1 chance - Jackpot)")
		log.Println("4. Quit")
		log.Println(" ")
		log.Println("Enter your choice: ")

		var choiceInput string
		fmt.Scan(&choiceInput)

		if choiceInput == "quit" || choiceInput == "4" {
			log.Println("Thanks for playing! Goodbye!")
			break
		}

		var chances int
		_, err := fmt.Sscan(choiceInput, &chances)
		if err != nil || chances < 1 || chances > 3 {
			log.Println("Error! Please select a valid difficulty!")
			log.Println("")
			continue
		}

		computer_number := rand.Intn(100)
		var userNumber int

		switch chances {
		case 1:
			log.Println("Easy Mode! You can go for 5 chances!")
			log.Println("")
			n := 5
			for i := 0; i < n; i++ {
				log.Printf("Enter your Guess: ")
				fmt.Scan(&userNumber)
				if userNumber != computer_number {
					log.Printf("Ooops! Lost that one you have %d chances left!\n", n-i-1)
				} else {
					log.Println("That was correct! Congratulations! You're truly the gambler of all time!")
					break
				}
			}
		case 2:
			log.Println("Medium Mode! You can go for 3 chances! Be careful!")
			log.Println("")
			n := 3
			for i := 0; i < n; i++ {
				log.Printf("Enter your Guess: ")
				fmt.Scan(&userNumber)
				if userNumber != computer_number {
					log.Printf("Ooops! Lost that one you have %d chances left!\n", n-i-1)
				} else {
					log.Println("That was correct! Congratulations! You're truly the gambler of all time!")
					break
				}
			}
		case 3:
			log.Println("WOW Thrilling! You get only 1 chance!")
			log.Println("")
			log.Printf("Enter your Guess: ")
			fmt.Scan(&userNumber)
			if userNumber != computer_number {
				log.Println("AAH UNFORTUNTE! This gamble did not pay off well :D")
			} else {
				log.Println("That was correct! Congratulations! You're truly the gambler of all time!")
			}
			log.Println("")
		}

		log.Println("The correct number was:", computer_number)
		log.Println("")
	}
}
