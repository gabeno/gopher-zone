package romans

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		NumberToConvert int
		Want            string
	}{
		{NumberToConvert: 1, Want: "I"},
		{NumberToConvert: 2, Want: "II"},
		{NumberToConvert: 3, Want: "III"},
		{NumberToConvert: 4, Want: "IV"},
		{NumberToConvert: 5, Want: "V"},
		{NumberToConvert: 6, Want: "VI"},
		{NumberToConvert: 9, Want: "IX"},
		{NumberToConvert: 10, Want: "X"},
		{NumberToConvert: 12, Want: "XII"},
		{NumberToConvert: 14, Want: "XIV"},
		{NumberToConvert: 18, Want: "XVIII"},
		{NumberToConvert: 20, Want: "XX"},
		{NumberToConvert: 39, Want: "XXXIX"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("convert %d", test.NumberToConvert), func(t *testing.T) {
			got := ConvertToRomans(test.NumberToConvert)
			want := test.Want

			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}
