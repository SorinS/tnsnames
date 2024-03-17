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
		fmt.Printf("Failed to file: %s\n", err)
		os.Exit(1)
	}
	// Create the Lexer
	lexer := tnsnames.NewtnsnamesLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := tnsnames.NewtnsnamesParser(stream)
	p.BuildParseTrees = true
	// Do the parsing
	addrs := []Address{}
	crtAddr := Address{}
	listener := &CustomTnsNamesListener{addrs, crtAddr, false}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Tnsnames())
	fmt.Printf("Got addresses: %d\n", len(listener.Addresses))
	for _, addr := range listener.Addresses {
		fmt.Printf("%s://%s:%s\n", addr.Protocol, addr.Host, addr.Port)
	}
}
