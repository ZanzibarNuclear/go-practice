package utils

import "fmt"

func HelloUtil(name string) string {
	message := fmt.Sprintf("Hey, %v. You are such a tool.", name)
	return message
}
