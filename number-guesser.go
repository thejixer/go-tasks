package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var chances int = 5
var shouldContinue bool = true
var isWon bool

// var target int = rand.Intn(100)
func absDiffInt(x, y int64) int64 {
	if x < y {
		return y - x
	}
	return x - y
}

func generateRandomNumberInRange(min int, max int) int64 {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(max-min+1) + min
	return int64(x)
}

func printStatus() {
	// str := "22"
	fmt.Printf("%v chances left \n", chances)

	// fmt.Printf("%c", target)
}

func finishGame(target int64) {
	if isWon {
		fmt.Printf("You Won, GZ!")
	} else {
		fmt.Printf("poor thing, u lost, the current number was %v", target)
	}
}

func main() {
	var target int64
	target = generateRandomNumberInRange(1, 99)

	for shouldContinue {
		if chances == 0 {
			shouldContinue = false
			break
		}
		printStatus()
		var userGuess int64
		fmt.Println("input userGuess:")
		// get user input
		_, err := fmt.Scanf("%d\n", &userGuess)
		if err != nil {
			fmt.Print("lol")
			log.Fatal(err)
		}

		// fmt.Printf("\n read userGuess: %d\n", userGuess)
		// fmt.Printf("%v %T \n", userGuess, userGuess)

		if userGuess < 1 || userGuess > 99 {
			fmt.Println("please provide a valid  number ")
			continue
		}
		dif := absDiffInt(userGuess, target)
		var temp string

		if dif <= 5 {
			temp = "HOT"
		} else if dif > 5 && dif <= 15 {
			temp = "WARM"
		} else {
			temp = "COLD"
		}

		if userGuess > target {
			fmt.Printf("the provided number is greater than the Number, %v \n", temp)

			chances--
			continue
		} else if userGuess < target {
			fmt.Printf("the provided number is less than the Number %v \n", temp)

			chances--
			continue
		} else {
			isWon = true
			break
		}

	}

	finishGame(target)

}
