package hangman

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const colorRed = "\033[1;31m"
const colorGreen = "\033[1;32m"
const colorReset = "\033[0m"

func LetterChoose() string { //func that return a string contane what user write in terminal
	fmt.Print("choose a letter :")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	if text == "é\n" || text == "è\n" {
		text = "e\n"
	}
	if text == "ç\n" {
		text = "c\n"
	}
	return text
}
func IsPresent(wordToFind string, letterChoose string) bool { // func returne true if letter choose by user is present in word to find
	if len(letterChoose) > 1 {
		for ii, valueLetterChoose := range letterChoose {
			for jj, valueWordToFind := range wordToFind {
				if ii == jj && valueWordToFind != valueLetterChoose {
					return false
				}
			}
		}
		return true
	} else {

		for _, valueWord := range wordToFind {
			for _, valueLettreChoose := range letterChoose {
				if string(valueWord) == string(valueLettreChoose) {
					return true
				}
			}
		}
	}
	return false
}

func AlreadySaid(letterChoose string, wordSaid string) string {
	err := false
	var said []string
	if len(letterChoose) == 1 {
		for _, valueWordSaid := range wordSaid {
			if letterChoose == string(valueWordSaid) {
				return wordSaid
			}

		}
	} else {
		for _, valueLetterChoose := range letterChoose {
			if valueLetterChoose == ' ' {
				err = true
			}
		}
		if err == false {
			wordSaid = wordSaid + letterChoose
			said = strings.Split(wordSaid, "\n")
			wordString := strings.Join(said, " ")
			fmt.Println("Already tried :", wordString, "\n")
		} else {

		}
	}
	return wordSaid

}

func VerifeChar(wordToFind string, wordUncomplet string) string {
	attempts := 11
	var wordSaid string
	wordInProgresse := wordUncomplet
	for attempts > 1 {
		letterChoose := LetterChoose()
		wordSaid = AlreadySaid(letterChoose, wordSaid)
		letterChoose = strings.Replace(letterChoose, "\n", "", -1)
		wordInProgresse = Reveal(wordToFind, wordInProgresse, letterChoose)
		fmt.Println(wordInProgresse)
		fmt.Println()
		if IsPresent(wordToFind, letterChoose) == true {
			Position(attempts)
			fmt.Println(string(colorGreen), "__________________________________________", string(colorReset))
			fmt.Println("remaining try :", attempts-1)
		} else {
			attempts--
			Position(attempts)
			fmt.Println(string(colorRed), "__________________________________________", string(colorReset))
			fmt.Println("remaining try :", attempts-1)
		}
		fmt.Print("\n")
		if wordInProgresse == wordToFind {
			return WinOrLoose(attempts, wordToFind)
		}
	}
	return WinOrLoose(attempts, wordToFind)
}

func Reveal(wordToFind string, wordInProgresse string, letterChoose string) string {
	word := []rune(wordInProgresse)
	index := 0
	if len(letterChoose) > 1 {
		for ii, valueLetterChoose := range letterChoose {
			for jj, valueWordToFind := range wordToFind {
				if ii == jj && valueWordToFind != valueLetterChoose {
					return wordInProgresse
				}
			}
		}
		return wordToFind
	} else {

		for _, letter := range wordToFind {
			for _, valueLettreChoose := range letterChoose {
				if string(letter) == string(valueLettreChoose) {
					word[index] = rune(letter)
				}
			}
			index++
		}
	}
	return string(word)
}
func Position(attempts int) {
	hangFile, _ := ioutil.ReadFile("../hangman.txt")
	file := strings.Split(string(hangFile), "\n\n")
	position := 11 - attempts
	if position < 0 {
		fmt.Println(file[0])
	} else {
		fmt.Println(file[position])
	}

}

func WinOrLoose(attempts int, wordToFind string) string {
	var endPrint string
	if attempts == 1 {
		endPrint = "Dommage ! Vous avez perdu, le mot était :" + wordToFind
	} else {
		endPrint = "Bravo ! Vous avez gagné, le mot était :" + wordToFind
	}
	return endPrint
}
