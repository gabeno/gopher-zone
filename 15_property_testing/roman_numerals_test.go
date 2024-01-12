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
