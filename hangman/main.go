package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var secret_word string
var gameOver bool
var correct_guesses []string
var failed_guesses []string
var chances int = 5

func main() {

	secret_word = SelectWord()

	for !gameOver {

		PrintGameStatus(secret_word)

		fmt.Println("guess a character : ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}
		char := strings.ToLower(string(line[0]))
		if isAlphabet := CheckIfAlphabet(char); !isAlphabet {
			fmt.Println("please provide a valid alphabet")
			continue
		}

		if x := Contains(correct_guesses, char); x {
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
		}

		dashifiedString := DashifyString(secret_word)

		if hasDash := strings.Contains(dashifiedString, "-"); !hasDash {
			gameOver = true
			fmt.Printf("you won, the word was -%v-", secret_word)
		}
	}
}
