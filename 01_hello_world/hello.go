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
	return englishPrefix + name
}
func main() {
	fmt.Println(Hello("world", ""))
}
