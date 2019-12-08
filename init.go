package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ViRu-ThE-ViRuS/Gunther/src/lexer"
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

	temp := lexer.Tokenize("(+ 1 2 (+ 1 2) )")
	fmt.Println(temp, len(temp))
}
