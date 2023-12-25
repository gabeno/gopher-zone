package main

import (
	"fmt"
)

const (
	spanish = "spanish"
	french  = "french"
	swahili = "swahili"

	englishPrefix = "Hello, "
	swahiliPrefix = "Hujambo, "
	frenchPrefix  = "Bonjour, "
	spanishPrefix = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		return spanishPrefix
	case french:
		return frenchPrefix
	case swahili:
		return swahiliPrefix
	default:
		return englishPrefix
	}
}

func main() {
	fmt.Println(Hello("world", ""))
}
