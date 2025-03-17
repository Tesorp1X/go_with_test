package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	germanHelloPrefix  = "Hallo, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix
	switch language {
	case "German":
		prefix = germanHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Liz", "English"))
}
