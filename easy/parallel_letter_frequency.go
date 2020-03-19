package easy

import ("fmt"
		"time"
		"strings")

//LetterCount information for counting letters for ParallelLetterFrequency
type LetterCount struct {
	letter string
	count int
}

//ParallelLetterFrequency counts the frequency of letters in texts using parallel computation
func ParallelLetterFrequency(text string) []LetterCount {
	start := time.Now()
	text = strings.ToUpper(text)

	m, letterCollection := []LetterCount{}, getLetterCollection(text)
	c := make(chan LetterCount)

	for _, letter := range letterCollection {
		 go countLetter(letter, text, c)
		 m = append(m, <-c)		 		 
	}
	
	fmt.Println(letterCollection, m, time.Since(start))
	return m
}

func countLetter(letter string, text string, c chan LetterCount)  {
	count := 0
	for _,  l := range text {
		if string(l) == letter {
			count++
		}
	}
	c <- LetterCount{letter : letter, count : count}
}

func getLetterCollection(text string) []string {
	letters := []string{}

	for _, letter := range strings.ReplaceAll(text, " ", "") {
		if !isExists(letters, letter) {
			letters = append(letters, string(letter))
		}
	}
	return letters
}
func isExists(letterCollection []string,  letter rune ) bool {
	for _,  l := range letterCollection  {
		if string(l) == string(letter) {
			return true
		}
	}

	return false
}
