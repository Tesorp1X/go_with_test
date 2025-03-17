package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	germanHelloPrefix  = "Hallo, "
	spanishHelloPrefix = "Hola, "
	russianHelloPrefix = "Привет, "
)

const (
	english = "English"
	german  = "German"
	spanish = "Spanish"
	russian = "Russian"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := getPrefix(language)
	return prefix + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case german:
		prefix = germanHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case russian:
		prefix = russianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Liz", "English"))
}
