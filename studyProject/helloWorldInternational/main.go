package main

import (
	"fmt"
	"os"
)

func getGreeting(lang string) string {
	switch lang {
	case "es":
		return "¡Hola Mundo!"
	case "fr":
		return "Bonjour le monde!"
	case "de":
		return "Hallo Welt!"
	case "it":
		return "Ciao Mondo!"
	case "pt":
		return "Olá Mundo!"
	case "zh":
		return "你好，世界！"
	case "ja":
		return "こんにちは世界！"
	case "ru":
		return "Привет, мир!"
	case "ar":
		return "مرحبا بالعالم!"
	case "hi":
		return "नमस्ते दुनिया!"
	case "ko":
		return "안녕하세요 세계!"
	case "tr":
		return "Merhaba Dünya!"
	default:
		return "Hello World!"
	}
}

func main() {
	lang := "tr" // Default language is Turkish
	// Check if a language argument is provided
	// If provided, it will override the default language
	// Example usage: go run main.go es
	// This will print "¡Hola Mundo!" for Spanish
	// If no argument is provided, it will print "Merhaba Dünya!" for Turkish
	if len(os.Args) > 1 {
		lang = os.Args[1]
	}
	fmt.Println(getGreeting(lang))
}
