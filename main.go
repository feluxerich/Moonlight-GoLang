package main

import (
	"Moonlight/lexer"
	"Moonlight/token"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	var source = "Hello\tThis is just for testing...,12.34"
	println(source)
	l := lexer.Lexer{Source: source}
	tokens := lexer.GenTokens(&l)
	iterateTokens(tokens)
	iterateTokens(lexer.ImproveTyping(tokens))
}

func usage() {
	println(os.Stderr, "Usage: %s <file>", os.Args[0])
	os.Exit(2)
}

func iterateTokens(tokens []token.Token) {
	for _, currentToken := range tokens {
		switch currentToken.Value {
		case "\r":
			fallthrough
		case "\n":
			println("<Newline>\t", currentToken.TokenType, currentToken.Value)
		case "\t":
			println("<Tab>\t\t", currentToken.TokenType, currentToken.Value)
		case " ":
			println("<Space>\t\t", currentToken.TokenType, currentToken.Value)
		default:
			println("<"+currentToken.TokenType+">\t", currentToken.Value)
		}
	}
}
