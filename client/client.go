package client

import (
	"fmt"
	"strings"
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
		args := strings.Split(ln, " ")
		cmd := strings.ToLower(args[0])
		if cmd != "store" {
			fmt.Println("Command not found")
		} else if len(args) < 2 {
			fmt.Println("Not enough arguments")
		} else {
			action := strings.ToLower(args[1])
			switch action {
			case "add":
				fmt.Println(CallUpload(args))
			case "ls":
				fmt.Println(CallList(args))
			case "rm":
				fmt.Println(CallRemove(args))
			case "update":
				fmt.Println(CallUpdate(args))
			}

		}

	}
}

func CallUpload(args []string) string {
	return ""
}

func CallList(args []string) string {
	return ""
}

func CallRemove(args []string) string {
	return ""
}

func CallUpdate(args []string) string {
	return ""
}
