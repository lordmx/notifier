package main

import(
	"os"
	"log"
	"io"
)

func InitLogger(path string) error {
	f, err := os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	defer f.Close()

	if err != nil {
	    return err
	}


    log.SetOutput(f)

    return nil
}