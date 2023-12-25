package main

import "fmt"

const englishPrefix = "Hello, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishPrefix
	switch language {
	case "spanish":
		prefix = "Hola, "
	case "french":
		prefix = "Bonjour, "
	case "swahili":
		prefix = "Hujambo, "

	}
	return prefix + name
}
func main() {
	fmt.Println(Hello("world", ""))
}
