package utils

import (
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
