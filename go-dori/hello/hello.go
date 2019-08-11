package main

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishPrefix = "Hola, "
const koreanPrefix = "안녕, "
const french = "French"
const spanish = "Spanish"
const korean = "Korean"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishPrefix
	case korean:
		prefix = koreanPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
