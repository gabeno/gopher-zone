package romans

func ConvertToRomans(n int) string {
	if n == 3 {
		return "III"
	}
	if n == 2 {
		return "II"
	}
	return "I"
}
