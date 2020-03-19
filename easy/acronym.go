package easy

import (
	"strings"
)

//Acronym Convert a phrase to its acronym.
func Acronym(text string) string {

	text = strings.TrimSuffix(text, "\n")

	wordArr := strings.Fields(text)

	var acronym string

	for _, word := range wordArr {
		acronym += strings.ToUpper(string(word[0]))
	}

	return acronym
}
