package main

import "fmt"

const englishPrefix = "Hello, "
const french = "French"
const spanish = "Spanish"
const frenchPrefix = "Bonjour, "
const spanishPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("George", "French"))
}
