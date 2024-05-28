package main

import (
	"createmodule/greet"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(log.LstdFlags)

	message, err := greet.Hello("Obito")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
