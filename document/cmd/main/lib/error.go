package lib

import (
	"log"
	"path/filepath"
	"runtime"
)

func CheckError(err error) error {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("[%s:%d] %v", filepath.Base(file), line, err)
	}
	return err
}
