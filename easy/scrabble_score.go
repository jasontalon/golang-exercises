package easy

import (
	"fmt"
	"strings"
)

//ScrabbleScore - Given a word, compute the Scrabble score for that word
func ScrabbleScore(text string) int {
	text = strings.TrimSuffix(text, "\n")

	score := 0

	scoreMap := mapScore()
	for _, letter := range text {
		score += scoreMap[strings.ToUpper(string(letter))]
	}

	fmt.Println("Your score: ", score)

	return score
}

func mapScore() map[string]int {
	m := map[string]int{}

	for _, letter := range []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"} {
		m[letter] = 1
	}

	for _, letter := range []string{"D", "G"} {
		m[letter] = 2
	}

	for _, letter := range []string{"B", "C", "M", "P"} {
		m[letter] = 3
	}

	for _, letter := range []string{"F", "H", "V", "W", "Y"} {
		m[letter] = 3
	}
	m["K"] = 5
	m["J"] = 8
	m["X"] = 8
	m["Q"] = 10
	m["Z"] = 10
	return m
}
