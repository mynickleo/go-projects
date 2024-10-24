package main

import (
	"fmt"
	"strings"
)

var bannedWords = []string{"badword", "непристойно", "🤬"}

func filterMessage(message string) string {
	returningMessage := []rune(message)

	for _, bannedWord := range bannedWords {
		bannedRunes := []rune(bannedWord)
		inputRunes := returningMessage

		for i := 0; i <= len(returningMessage)-len(bannedRunes); i++ {
			if strings.HasPrefix(string(inputRunes[i:]), bannedWord) {
				for j := 0; j < len(bannedRunes); j++ {
					returningMessage[i+j] = '*'
				}
			}
		}
	}

	return string(returningMessage)
}

func main() {
	message := "Hello World 🤬"
	fmt.Println(filterMessage(message))
}
