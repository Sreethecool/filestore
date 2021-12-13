package utils

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Sreethecool/filestore/server/models"
)

func IsAllowedCommand(cmd string) bool {
	for _, v := range models.CmdList {
		if v == strings.ToLower(cmd) {
			return true
		}
	}
	return false
}

func Contains(list []string, key string) bool {
	for _, v := range list {
		if v == key {
			return true
		}
	}
	return false
}

func CreateDirIfNotExists(dirName string) error {
	var err error
	if _, err = os.Stat(dirName); os.IsNotExist(err) {
		err = os.Mkdir(dirName, os.ModePerm)
		if err != nil {
			fmt.Println("Cant create upload folder")
		}
	}
	return err
}

func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}
