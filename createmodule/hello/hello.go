package main

import (
	"createmodule/greet"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(log.LstdFlags)

	names := []string{"Naruto", "Hinata", "Obito"}

	message, err := greet.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
