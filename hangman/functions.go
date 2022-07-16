package hangman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func GetRandomItem() string {
	rand.Seed(time.Now().UnixNano())
	return Words[rand.Intn(len(Words))]
}

func CheckIfAlphabet(r string) bool {
	return unicode.IsLetter([]rune(r)[0])
}

func Contains(x []string, v string) bool {
	for _, s := range x {
		if v == s {
			return true
		}
	}
	return false
}

func DashifyString(secret_word string) string {
	str := ""
	for _, char := range secret_word {
		c := string(char)
		if x := Contains(correct_guesses, c); x {
			str += c
		} else {
			str += "-"
		}
	}
	return str
}

func PrintGameStatus(secret_word string) {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("#######################################")
	fmt.Printf("## the word has %v letters \n", utf8.RuneCountInString(secret_word))
	fmt.Printf("## the word is: %v \n", DashifyString(secret_word))
	fmt.Printf("## u can make %v more mistakes \n", chances-len(failed_guesses))
	fmt.Print("## failed letters : ", failed_guesses, "\n")
	fmt.Println("#######################################")
}


func SelectWord() string {
	resp, err := http.Get("https://random-words-api.vercel.app/word")
	if err != nil {
		return GetRandomItem()
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