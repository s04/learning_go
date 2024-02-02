package main

const spanish = "Spanish"
const french = "French"
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + name

}

func main() {
	// fmt.Println(Hello("Chris"))
}
