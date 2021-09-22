package main

import (
	"log"
	"os"
)

func LogError(err error) error {
	if err != nil {
		log.Fatal(err.Error())
	}
	return err
}

func ThenTerminate(err error) error {
	if err != nil {
		os.Exit(1)
	}
	return err
}
