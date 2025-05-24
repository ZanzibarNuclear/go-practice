package main

import (
	"fmt"

	"github.com/ZanzibarNuclear/greetings"
)

func main() {
	message := greetings.Hello("boo")
	fmt.Println(message)
}