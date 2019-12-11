package common

import (
	"log"
	"os"
)

const WORK_DIR = "/var/lib/keylime"

func ChangeDir(path string, mode int) {
	os.Mkdir(path, os.FileMode(mode))
}

func GetCWD() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}
