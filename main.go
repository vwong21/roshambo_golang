package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func getUserChoice() string {

	chooseAgain := true
	var userChoice string
	for chooseAgain {

		fmt.Print("Select Rock, Paper, or Scissors: ")
		fmt.Scan(&userChoice)
		userChoice = strings.ToLower(userChoice)

		if userChoice == "rock" || userChoice == "paper" || userChoice == "scissors" {
			return userChoice
		} else {
			fmt.Print("Invalid selection. Please choose again.\n")
		}
	}
	return ""
}

func getComChoice() string {
	choices := []string{"rock", "paper", "scissors"}
	comChoice := choices[rand.IntN(3)]
	return comChoice
}

func getResult(userChoice string, comChoice string) {
	if userChoice == comChoice {
		fmt.Print("\nDraw!\n")
	} else if (userChoice == "rock" && comChoice == "scissors") || (userChoice == "paper" && comChoice == "rock") || (userChoice == "scissors" && comChoice == "paper") {
		fmt.Print("\nYou are the winner!\n")
	} else {
		fmt.Print("\nYou lose!\n")
	}
}

func main() {
	game := true
	var userChoice, comChoice string
	for game {
		var playAgain string
		userChoice = getUserChoice()
		comChoice = getComChoice()
		fmt.Print("Computer has selected ", comChoice)
		getResult(userChoice, comChoice)
		fmt.Print("Play again? (Y/N): ")
		fmt.Scan(&playAgain)
		playAgain = strings.ToLower(playAgain)
		if playAgain == "n" {
			game = false
		}

	}
}
