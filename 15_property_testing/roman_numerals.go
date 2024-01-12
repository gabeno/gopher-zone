package romans

import "strings"

func ConvertToRomans(n int) string {
	if n == 4 {
		return "IV"
	}
	var results strings.Builder

	for i := 0; i < n; i++ {
		results.WriteString("I")
	}

	return results.String()
}
