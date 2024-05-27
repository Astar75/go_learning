package greet

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Привет, %v! Добро пожаловать!", name)
	return message
}
