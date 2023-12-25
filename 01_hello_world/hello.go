package main

import "fmt"

const englishPrefix = "Hello, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "spanish" {
		return "Hola, " + name
	}
	if language == "french" {
		return "Bonjour, " + name
	}
	if language == "swahili" {
		return "Hujambo, " + name
	}
	return englishPrefix + name
}
func main() {
	fmt.Println(Hello("world", ""))
}
