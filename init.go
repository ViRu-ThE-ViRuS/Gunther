package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ViRu-ThE-ViRuS/Gunther/src/lexer"
	"github.com/ViRu-ThE-ViRuS/Gunther/src/parser"
	"github.com/joho/godotenv"
)

var (
	args = os.Args[1:]
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	if len(os.Getenv("DEBUG")) == 0 {
		log.SetOutput(ioutil.Discard)
	}
}

func main() {
	log.Printf("arguments: %v", args)

	temp, err := lexer.Tokenize("(+ 1 2)")
	if err != nil {
		log.Fatalln("lexer crashed")
		os.Exit(0)
	}

	temp, err = parser.Parse(temp)
}
