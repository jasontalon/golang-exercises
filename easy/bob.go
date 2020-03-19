package easy

import (
	"fmt"
	"strings"
)

//Bob is a lackadaisical teenager. In conversation, his responses are very limited...
func Bob(text string) string {
	fmt.Print("Enter text: ")
	text = strings.TrimSuffix(text, "\n")

	isQuestion := strings.HasSuffix(text, "?")
	isAllUpperCase := strings.ToUpper(text) == text
	isNothing := len(text) == 0

	if isQuestion && isAllUpperCase {
		return "Calm down, I know what I'm doing!"
	} else if isNothing {
		return "Fine. Be that way!"
	} else if isAllUpperCase {
		return "Whoa, chill out!"
	} else if isQuestion {
		return "Sure."
	} else {
		return "Whatever."
	}
}
