package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	args = os.Args[1:]
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .enf file found")
	}

	if len(os.Getenv("DEBUG")) == 0 {
		log.SetOutput(ioutil.Discard)
	}
}

func main() {
	log.Printf("arguments: %v", args)
}
