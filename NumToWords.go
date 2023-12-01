package main

import (
	"fmt"
	"strings"
)

var (
	singleDigits = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	teenDigits   = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tensDigits   = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
)

func NumberToWords(number int) string {
	if number < 0 {
		return "minus " + NumberToWords(-number)
	}

	if number < 10 {
		return singleDigits[number]
	} else if number < 20 {
		return teenDigits[number-10]
	} else if number < 100 {
		tens := number / 10
		remainder := number % 10

		result := tensDigits[tens]
		if remainder > 0 {
			result += "-" + singleDigits[remainder]
		}
		return result
	} else if number < 1000 {
		hundreds := number / 100
		remainder := number % 100

		result := singleDigits[hundreds] + " hundred"
		if remainder > 0 {
			result += " " + NumberToWords(remainder)
		}
		return result
	} else if number < 1000000 {
		thousands := number / 1000
		remainder := number % 1000

		result := NumberToWords(thousands) + " thousand"
		if remainder > 0 {
			result += " " + NumberToWords(remainder)
		}
		return result
	} else {
		millions := number / 1000000
		remainder := number % 1000000

		result := NumberToWords(millions) + " million"
		if remainder > 0 {
			result += " " + NumberToWords(remainder)
		}
		return result
	}
}

func main() {
	number := 12345
	words := NumberToWords(number)
	fmt.Printf("%d in words: %s\n", number, strings.Title(words))
}
