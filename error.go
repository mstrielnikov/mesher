package main

import (
	"fmt"
	"log"
	"os"
)

func Ok(err error) error {
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func Log(err error) error {
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
