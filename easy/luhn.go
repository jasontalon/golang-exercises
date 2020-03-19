package easy

import (
	"fmt"
	"strconv"
)

//LuhnValidationResponse for the luhn validation
type LuhnValidationResponse struct {
	modifiedCreditCard string
	isValid            bool
}

func isEvenNumber(number int) bool {
	return number%2 == 0
}

func isOddNumber(number int) bool {
	return !isEvenNumber(number)
}

func isValidCreditCard(creditCard string) bool {
	sum := 0
	for _, digit := range creditCard {

		numDigit, _ := strconv.Atoi(string(digit))
		sum += numDigit
	}
	fmt.Printf("total sum: %d\n", sum)
	return sum%10 == 0
}

//Luhn validates credit card number using Luhn algorithm
func Luhn(creditCard string) LuhnValidationResponse {
	modifiedCreditCard := ""

	for index, digit := range creditCard {
		strDigit := string(digit)
		numDigit, _ := strconv.Atoi(strDigit)

		if isOddNumber(index + 1) {
			doubleUp := numDigit * 2
			if doubleUp > 9 {
				doubleUp -= 9
			}
			strDoubleUp := strconv.FormatInt(int64(doubleUp), 10)
			modifiedCreditCard += strDoubleUp
		} else {
			modifiedCreditCard += strDigit
		}
	}
	fmt.Printf("Modified Credit Card: %s \n", modifiedCreditCard)

	isValid := isValidCreditCard(modifiedCreditCard)

	fmt.Printf("Valid: %t\n", isValid)

	return LuhnValidationResponse{isValid: isValid, modifiedCreditCard: modifiedCreditCard}
}
