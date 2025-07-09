package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	message := greetings.Hello("Jian")
	fmt.Println(message)
}
