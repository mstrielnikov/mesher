package main

import (
	"log"
	"os"
)

func Ok(err error) error {
	if err != nil {
		log.Fatal(err.Error())
	}
	return err
}

func Must(err error) {
	if err != nil {
		os.Exit(-1)
	}
}
