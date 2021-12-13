package main

import (
	"github.com/Sreethecool/filestore/client"
)

func main() {
	//URL of the server need to passed.
	cli := client.GetClient("http://localhost:8080")
	cli.Start()
}
