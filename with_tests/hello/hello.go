package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const english = "English"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "

func greetingPrefix(language string) (prefix string) {
    switch language {
    case spanish:
        prefix = spanishHelloPrefix
    case french:
        prefix = frenchHelloPrefix
    case english:
        prefix = englishHelloPrefix
    default:
        prefix = portugueseHelloPrefix
    }
    return
}

func Hello(name, language string) string {
    if name == "" {
        name = "Mundo"
    }

    return greetingPrefix(language) + name
}

func main() {
    fmt.Println(Hello("Batata", ""))
}
