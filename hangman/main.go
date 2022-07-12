package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

var words = [40]string{
	"python","apple","orange",
	"existansialism","disgusting","christ","yummy","ronaldo",
	"bitcoin","football","cryptocurrency","corridor","execution",
	"whiteboard","delicious","exciting","immortality","dangrous",
	"cosmon","homeless","psycology","necklace","zebra","rhinoceros",
	"japan","spider","monnument","pyramid","avenue","antidisestablishmentarianism",
	"lantern","liberty","shoulder","finger","thumb","locket","monsoone","avadacadabra",
	"ethereum","pronunciation",
}

func getRandomItem()string {
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func checkIfAlphabet(r string)bool {
	return unicode.IsLetter([]rune(r)[0])
}

var secret_word string;
var gameOver bool
var correct_guesses []string
var failed_guesses []string
var chances int = 5

func contains(x []string, v string) bool {
	for _, s := range x {
		if v == s {
			return true
		}
	}
	return false
}

func dashifyString(secret_word string) string {

	str := ""

	for _, char := range secret_word {
		var c string
		c = string(char)
    if x := contains(correct_guesses, c); x {
			str += c
		} else {
			str += "-"
		}
	}

	return str
}

func printGameStatus(secret_word string) {
	fmt.Println("#######################################")
  fmt.Printf("the word has %v letters \n", utf8.RuneCountInString(secret_word))
  fmt.Printf("the word is: %v \n", dashifyString(secret_word))
  fmt.Printf("u can make %v more mistakes \n", chances - len(failed_guesses))
  fmt.Println("#######################################")
}

func main() {
	secret_word = getRandomItem()
	fmt.Println("the secret word is " + secret_word)


	// utf8.RuneCountInString(s)
	// reader := bufio.NewReader(os.Stdin)
	// userGuess, err := reader.ReadString('\n')

	// if err != nil {
	// 	log.Fatal(err)
	// }


	// fmt.Printf("output: %v", userGuess)
	// x := len(userGuess)
	// fmt.Print(x)


	for !gameOver {

		// if ()

		printGameStatus(secret_word)

		fmt.Println("input text:")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
	
		
		if err != nil {
			log.Fatal(err)
		}
		char := string(line[0])
		if isAlphabet := checkIfAlphabet(char); !isAlphabet {
			fmt.Println("please provide a valid alphabet")
			continue
		}

		if x := contains(correct_guesses, char); x {
			fmt.Println("you already guessed this letter")
			continue
		}

		if trueguess := strings.Contains(secret_word, char); trueguess {
			correct_guesses = append(correct_guesses, char)
		} else {
			failed_guesses = append(failed_guesses, char)
		}

		if len(failed_guesses) == chances {
			fmt.Printf("poor thing u lost, the correct word was -%v- \n", secret_word)
			gameOver = true
			// break
		}

		var dashifiedString string 
		dashifiedString = dashifyString(secret_word)

		if hasDash := strings.Contains(dashifiedString, "-"); !hasDash {
			gameOver = true
			fmt.Println("you won")
		}
		
	}

	// fmt.Print())
	// fmt.Print()
}