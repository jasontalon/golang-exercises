package medium

import (
	"strings"
)

//RailFence object
type RailFence struct {
}

//Encode encode text to rail code cipher
func (RailFence) Encode(text string) string {
	text, topRow, bottomRow, middleRow := strings.ReplaceAll(text, " ", ""), []int{}, []int{}, []int{}

	for i:= 0; len(text) > i; i++ {
		if(i %4 == 0) {
			topRow = append(topRow, i)			
		} else if (i % 2 == 0) {
			bottomRow = append(bottomRow, i)			
		} else {
			middleRow = append(middleRow, i)
		}
	}

	encodedText := ""
	
	for _, item := range append(append(append([]int{}, topRow...), middleRow...), bottomRow...) {
		encodedText += string(text[item])
	}
	
	return encodedText
}

//Decode rail fence to plain text
func (RailFence) Decode(encodedText string) string {
	topRow, bottomRow, middleRow := []int{}, []int{}, []int{}
	decodedText := ""
	
	for i:= 0; len(encodedText) > i; i++ {
		decodedText += "*"
		if(i %4 == 0) {
			topRow = append(topRow, i)			
		} else if (i % 2 == 0) {
			bottomRow = append(bottomRow, i)			
		} else {
			middleRow = append(middleRow, i)
		}
	}

	for index, item := range append(append(append([]int{}, topRow...), middleRow...), bottomRow...) {
		decodedText = replace(decodedText, rune(encodedText[index]), item)
	}
	
	return strings.ReplaceAll(decodedText, "*", "")
}

func replace(str string, replacement rune, index int) string {	
	return str[:index] + string(replacement) + str[index+1:]
}