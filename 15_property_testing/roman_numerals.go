package romans

import "strings"

func ConvertToRomans(n int) string {
	var results strings.Builder

	for i := n; i > 0; i-- {
		if i == 4 {
			results.WriteString("IV")
			break
		}
		results.WriteString("I")
	}

	return results.String()
}
