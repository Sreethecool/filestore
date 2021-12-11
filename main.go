package main

import (
	"fmt"

	"github.com/Sreethecool/filestore/client"
	"github.com/Sreethecool/filestore/command"
)

func main() {

	fmt.Println(command.Execute("ls", []string{}))
	client.Start()

}
