package main

import (
	"createmodule/greet"
	"fmt"
)

func main() {
	message := greet.Hello("Obito")
	fmt.Println(message)
}
