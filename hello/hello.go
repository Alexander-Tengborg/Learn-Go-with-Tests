package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const russian = "Russian"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const russianHelloPrefix = "Privyet, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case russian:
		prefix = russianHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Linda", ""))
}