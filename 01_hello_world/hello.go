package main

import "fmt"

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

	prefix := englishPrefix
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	case swahili:
		prefix = swahiliPrefix

	}
	return prefix + name
}
func main() {
	fmt.Println(Hello("world", ""))
}
