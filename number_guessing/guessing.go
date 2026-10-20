package main

import (
	"fmt"
	"math/rand"
)

var gameON bool = true
var userGuess int
var userTrials int

func main() {
	userTrials = 3

	//Generate random number
	randomNumber := rand.Intn(11)

	for gameON {

		if userTrials == 0 {
			fmt.Println("You have 0 chance(s) left")
			fmt.Printf("The correct number was %d \n", randomNumber)
			fmt.Print("GAME OVER")
			break

		} else {
			//Ask for user to input number
			fmt.Print("Enter your guess (1 to 10): ")
			_, err := fmt.Scan(&userGuess)

			if err != nil {
				userTrials -= 1
				fmt.Printf("Wrong input. %d chance(s) left", userTrials)

			} else {
				//Check input against random number
				if userGuess > randomNumber {
					fmt.Println("Guess too high! Try again")
					userTrials -= 1
					fmt.Printf("%d chance(s) left \n\n", userTrials)

				} else if userGuess < randomNumber {
					fmt.Println("Guess too low! Try again")
					userTrials -= 1
					fmt.Printf("%d chance(s) left \n\n", userTrials)

				} else {
					fmt.Println("Hurray! you got it.")
					fmt.Printf("The correct guess is %d, ", randomNumber)
					break
				}
			}
		}
	}
}
