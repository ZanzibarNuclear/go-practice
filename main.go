package main

import (
	"fmt"

	"github.com/ZanzibarNuclear/go-practice/internal/utils"
	"github.com/ZanzibarNuclear/greetings"
)

func main() {
	message := greetings.Hello("皆さん")
	altMessage := utils.HelloUtil("皆さん")
	fmt.Println(message)
	fmt.Println(altMessage)
}
