package utils

import (
	"errors"
	"log"
	"os"
)

/*
Create a working directory folder if not exist

param path folder to be created
return nil
*/
func CreateFolder(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
