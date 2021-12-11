package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Sreethecool/filestore/client"
	"github.com/Sreethecool/filestore/server"
	"github.com/Sreethecool/filestore/server/command"
	"github.com/Sreethecool/filestore/server/models"
)

func main() {
	res, err := command.Execute("ls -l")
	fmt.Println("res:", res)
	fmt.Println("Err:", err)
	go server.RunServer()
	var resp models.Response
	json.Unmarshal([]byte(client.CallList([]string{})), &resp)

	fmt.Println(resp.Message)
	files := strings.Split(resp.Data.(string), "\n")
	fmt.Println(files)
	var a string
	fmt.Scanln(&a)
}
