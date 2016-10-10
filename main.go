package main

import "strings"

var (
	romans = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	digits = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
)

// Ditit2Roman - convert number to roman
func Ditit2Roman(num int) string {
	var roman = ""

	for i := range romans {
		for num >= digits[i] {
			roman += romans[i]
			num -= digits[i]
		}
	}

	return roman
}

// Roman2Digit - convert roman to number
func Roman2Digit(roman string) int {
	var number = 0

	for i := range romans {
		for strings.Index(roman, romans[i]) == 0 {
			number += digits[i]
			roman = strings.Replace(roman, romans[i], "", 1)
		}
	}

	return number
}
