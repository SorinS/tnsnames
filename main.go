package main

import (
	"bramp.net/antlr4/tnsnames"
	"flag"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
)

func newCharStream(filename string) (antlr.CharStream, error) {
	var input antlr.CharStream
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func main() {
	fmt.Println("TNSNames parser")
	var filepath string
	flag.StringVar(&filepath, "f", "", "file to parse")

	flag.Parse()
	input, err := newCharStream(filepath)
	if err != nil {
		fmt.Printf("Failed to open example file: %s\n", err)
		os.Exit(1)
	}
	// Create the Lexer
	lexer := tnsnames.NewtnsnamesLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := tnsnames.NewtnsnamesParser(stream)
	p.BuildParseTrees = true
	// add the listener

	// Finally test
	antlr.ParseTreeWalkerDefault.Walk(&exampleListener{}, p.Tnsnames())
}
