package romans

import "strings"

func ConvertToRomans(n int) string {
	var results strings.Builder

	for i := 0; i < n; i++ {
		results.WriteString("I")
	}

	return results.String()
}
