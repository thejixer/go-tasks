package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

var words = [50]string{
	"python", "apple", "orange",
	"existansialism", "disgusting", "christ", "yummy", "ronaldo",
	"bitcoin", "football", "cryptocurrency", "corridor", "execution",
	"whiteboard", "delicious", "exciting", "immortality", "dangrous",
	"cosmon", "homeless", "psycology", "necklace", "zebra", "rhinoceros",
	"japan", "spider", "monnument", "pyramid", "avenue", "antidisestablishmentarianism",
	"lantern", "liberty", "shoulder", "finger", "thumb", "locket", "monsoone", "avadacadabra",
	"ethereum", "pronunciation", "naruto", "neverland", "golang", "france",
	"hogrider", "bottle", "heavymetal", "water", "windrunner", "bloodmoon",
}

func getRandomItem() string {
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func checkIfAlphabet(r string) bool {
	return unicode.IsLetter([]rune(r)[0])
}

var secret_word string
var gameOver bool
var correct_guesses []string
var failed_guesses []string
var chances int = 5

type WordStruct struct {
	Word          string `json:"word"`
	Definition    string `json:"definition"`
	Pronunciation string `json:"pronunciation"`
}

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
		c := string(char)
		if x := contains(correct_guesses, c); x {
			str += c
		} else {
			str += "-"
		}
	}
	return str
}

func printGameStatus(secret_word string) {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("#######################################")
	fmt.Printf("## the word has %v letters \n", utf8.RuneCountInString(secret_word))
	fmt.Printf("## the word is: %v \n", dashifyString(secret_word))
	fmt.Printf("## u can make %v more mistakes \n", chances-len(failed_guesses))
	fmt.Print("## failed letters : ", failed_guesses, "\n")
	fmt.Println("#######################################")
}

func selectWord() string {
	resp, err := http.Get("https://random-words-api.vercel.app/word")
	// in case there was an issue with the connection, choose a random word from our own dictionary
	if err != nil {
		return getRandomItem()
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		sb := string(body)

		var theRes []WordStruct
		idk := []byte(sb)
		json.Unmarshal(idk, &theRes)

		return strings.ToLower(theRes[0].Word)
	}
}

func finishGame(hasWon bool) {
	if hasWon {
		fmt.Printf("poor thing u lost, the correct word was -%v- \n", secret_word)
	} else {
		fmt.Printf("you won, the word was -%v-", secret_word)
	}
}

func main() {

	secret_word = selectWord()

	for !gameOver {

		printGameStatus(secret_word)

		fmt.Println("guess a character : ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}
		char := strings.ToLower(string(line[0]))
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
			finishGame(false)
			gameOver = true
		}

		dashifiedString := dashifyString(secret_word)

		if hasDash := strings.Contains(dashifiedString, "-"); !hasDash {
			finishGame(true)
			gameOver = true
		}

	}

}
