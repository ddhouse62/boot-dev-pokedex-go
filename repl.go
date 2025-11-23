package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	textNoWhiteSpace := strings.TrimSpace(lowerText)
	listWords := strings.Split(textNoWhiteSpace, " ")
	cleanedInput := []string{}

	for _, str := range listWords {
		if str != "" {
			cleanedInput = append(cleanedInput, str)
		}
	}
	return cleanedInput
}
