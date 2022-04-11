package main

import (
	"fmt"
	"log"
	"os"
)

func Ok(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

func Log(err error) error {
	if err != nil {
		log.Fatal(err.Error())
	}
	return err
}
