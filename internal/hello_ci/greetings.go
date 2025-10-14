package main

import "fmt"

// Greet возвращает приветствие в зависимости от языка.
func Greet(language, name string) string {
	switch language {
	case "en":
		return fmt.Sprintf("Hello, %s", name)
	case "es":
		return fmt.Sprintf("Hola, %s", name)
	case "ru":
		return fmt.Sprintf("Привет, %s", name)
	default:
		return fmt.Sprintf("Hello, %s", name)
	}
}
