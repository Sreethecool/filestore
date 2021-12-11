package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Sreethecool/filestore/server/models"
	"github.com/Sreethecool/filestore/server/utils"
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
			case "wc":
				fmt.Println(CallWordCount(args))
			case "freq-words":
				fmt.Println(CallFrequentWords(args))
			default:
				fmt.Println("Action Command Not found")
			}
		}
	}
}

func CallUpload(args []string) string {

	if len(args) < 3 {
		return "Error: Not enough arguments"
	}
	args = args[2:]
	ls := CallList(args)
	if strings.Contains(ls, "Error") {
		return "Error: Failed to Get the list of files"
	}
	files := strings.Split(ls, "\n")
	for _, v := range args {
		if utils.Contains(files, v) {
			return fmt.Sprintf("Error: One of the File %s already Present in server. To check list of Files in server use ls.", v)
		}
	}

	res := uploadFiles(args)
	if strings.Contains(res, "Error") {
		return res
	}
	return "Upload Sucess"
}

func CallList(args []string) string {

	url := "http://localhost:8080/list"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "Error: Failed to Prepare Request"
	}
	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return "Error: Request Failed"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "Error: Failed to Read Response"
	}
	return getResponse(body)
}

func CallRemove(args []string) string {

	if len(args) > 3 {
		return "Error: syntax error: rm <filename> cannot delete multiple files."
	} else if args[2] == "" {
		return "Error: syntax error: file missing."
	}
	ls := CallList(args)
	if strings.Contains(ls, "Error") {
		return "Error: Failed to Get the list of files"
	}
	args = args[2:]
	files := strings.Split(ls, "\n")
	for _, v := range args {
		if !utils.Contains(files, v) {
			return fmt.Sprintf("Error: Some Files %s are not found in server. To check the list of files use ls", v)
		}
	}

	url := "http://localhost:8080/delete"
	method := "POST"

	var request models.DeleteRequest
	request.Filename = args[0]
	payload, err := json.Marshal(request)
	if err != nil {
		return "Error: Failed to marshal request"
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(string(payload)))
	if err != nil {
		return "Error: Failed to Prepare Request"
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return "Error: Request Failed"
	}
	defer res.Body.Close()

	return "File removed from Server"
}

func CallUpdate(args []string) string {
	if len(args) < 3 {
		return "Error: Not enough arguments"
	}
	args = args[2:]
	res := uploadFiles(args)
	if strings.Contains(res, "Error") {
		return res
	}
	return "Update Sucess"
}

func CallWordCount(args []string) string {
	if len(args) > 2 {
		return "Error: wc should not have arguments"
	}
	return callExecute("wc", []string{})
}

func CallFrequentWords(args []string) string {
	if len(args) > 6 {
		return "Error: Have more arguments"
	}
	return callExecute("freq-words", args[2:])
}

func callExecute(cmd string, args []string) string {
	url := "http://localhost:8080/run"
	method := "POST"

	var request models.ExecuteRequest
	request.Cmd = cmd
	request.Args = args[:]

	payload, err := json.Marshal(request)
	if err != nil {
		return "Error: Failed to marshal request"
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(string(payload)))
	if err != nil {
		return "Error: Failed to Prepare Request"
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return "Error: Request Failed"
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "Error: Failed to Read Response"
	}
	return getResponse(body)
}
func uploadFiles(args []string) string {

	url := "http://localhost:8080/upload"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for _, v := range args {

		file, err := os.Open(v)
		if err != nil {
			return "Error: Failed to Open/Read file:" + v
		}
		defer file.Close()
		part, err := writer.CreateFormFile("files", filepath.Base(v))
		if err != nil {
			return "Error: Failed to create form data from file:" + v
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return "Error: Failed to Prepare Read Contents of File:" + v
		}
		err = writer.Close()
		if err != nil {
			return "Error: Falied to Read File" + v
		}
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "Error:Failed to Prepare Request"
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil && res.StatusCode != http.StatusOK {
		return "Error:Failed to Upload"
	}
	defer res.Body.Close()

	return "Sucess"
}
func getResponse(body []byte) string {
	var resp models.Response
	json.Unmarshal(body, &resp)
	return resp.Data.(string)
}
