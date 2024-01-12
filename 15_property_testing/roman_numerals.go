package romans

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
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
