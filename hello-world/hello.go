package helloworld

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, lan string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(lan) + name
}

func greetingPrefix(lan string) (prefix string) {
	switch lan {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}
