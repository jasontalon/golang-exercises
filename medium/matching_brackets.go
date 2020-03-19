package medium

import (
	"errors"
	"fmt"
	"strings"
)

type bracket struct {
	start string
	end   string
}

var brackets = []bracket{{start: "{", end: "}"}, {start: "(", end: ")"}, bracket{start: "[", end: "]"}}

func findBracketByStart(start string) (*bracket, error) {
	for _, bracketItem := range brackets {
		if bracketItem.start == start {
			return &bracketItem, nil
		}
	}
	return nil, errors.New("not exists")
}

func isBracket(letter rune, startingBrackets bool) bool {
	strLetter := string(letter)

	for _, bracketItem := range brackets {
		bracket := bracketItem.end
		if startingBrackets {
			bracket = bracketItem.start
		}
		if strLetter == string(bracket) {
			return true
		}
	}

	return false
}

func hasBrackets(text string) bool {
	for _, letter := range text {
		if (isBracket(letter, true) || isBracket(letter, false)) == true {
			return true
		}
	}

	return false
}

//BracketValidator validates the bracket inside the text
func BracketValidator(text string) bool {
	trimmed := text
	for _, letter := range text {
		if !isBracket(letter, true) {
			continue
		}

		bracket, err := findBracketByStart(string(letter))
		if err != nil {
			fmt.Println(err.Error())
			return false
		}

		lenBefore := len(trimmed)
		trimmed = strings.Replace(trimmed, bracket.end, "", 1)

		if lenBefore > len(trimmed) {
			trimmed = strings.Replace(trimmed, bracket.start, "", 1)
		}
	}

	return hasBrackets(trimmed)
}
