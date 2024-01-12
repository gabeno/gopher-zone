package romans

import "strings"

func ConvertToRomans(n int) string {
	var results strings.Builder

	for n > 0 {
		switch {
		case n > 4:
			results.WriteString("V")
			n -= 5
		case n > 3:
			results.WriteString("IV")
			n -= 4
		default:
			results.WriteString("I")
			n--
		}
	}

	return results.String()
}
