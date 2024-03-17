package main

import (
	"bramp.net/antlr4/tnsnames"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TnsNamesListener interface {
	antlr.ParseTreeListener
	EnterPort(c *tnsnames.PortContext)
}

type exampleListener struct {
	*tnsnames.BasetnsnamesParserListener
}

func (l *exampleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}
