package romans

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRomans(n int) string {
	var results strings.Builder

	for _, numeral := range allRomanNumerals {
		for n >= numeral.Value {
			results.WriteString(numeral.Symbol)
			n -= numeral.Value
		}
	}

	return results.String()
}

func ConvertToArabic(s string) int {
	total := 0
	for range s {
		total++
	}
	return total
}
