package romans

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description     string
		NumberToConvert int
		Want            string
	}{
		{Description: "1 converts to I", NumberToConvert: 1, Want: "I"},
		{Description: "2 converts to II", NumberToConvert: 2, Want: "II"},
		{Description: "3 converts to III", NumberToConvert: 3, Want: "III"},
		{Description: "4 converts to IV", NumberToConvert: 4, Want: "IV"},
		{Description: "5 converts to V", NumberToConvert: 5, Want: "V"},
		{Description: "6 converts to VI", NumberToConvert: 6, Want: "VI"},
		{Description: "9 converts to IX", NumberToConvert: 9, Want: "IX"},
		{Description: "10 converts to X", NumberToConvert: 10, Want: "X"},
		{Description: "12 converts to XII", NumberToConvert: 12, Want: "XII"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRomans(test.NumberToConvert)
			want := test.Want

			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}
