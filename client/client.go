package client

import (
	"fmt"
)

func Start() {
	fmt.Println("File store Operations")
	for {
		fmt.Print("$")
		var ln string
		fmt.Scanln(&ln)
		if ln == "exit" {
			break
		}

	}
}
