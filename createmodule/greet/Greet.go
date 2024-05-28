package greet

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// возвращает map которая связывает каждого из названных людей
// с приветственными сообщением.
func Hellos(names []string) (map[string]string, error) {
	// карта для связи имен с сообщениями
	messages := make(map[string]string)
	// Проходим по полученному списку имен, вызывая функцию Hello,
	// что бы получить сообщение для каждого имени.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// в карте свяжем полученое сообщение с именем
		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
