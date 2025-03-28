package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	polish  = "Polish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	polishHelloPrefix  = "Cześć, "
)

func main() {
	fmt.Println(Hello("Marek", "Polish"))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case polish:
		prefix = polishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
