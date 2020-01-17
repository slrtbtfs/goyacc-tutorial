package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	yyErrorVerbose = true
	interpreter := interpreter{}

	interpreter.vars = make(map[string]int)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Bye.")
			return
		}

		interpreter.position = 0
		interpreter.input = input
		yyParse(&interpreter)
	}
}

const EOF = 0

type interpreter struct {
	input    string
	position int
	vars     map[string]int
}

func (*interpreter) Error(e string) {
	fmt.Println(e)
}

func (l *interpreter) Lex(lval *yySymType) int {
	tokenStart := l.position
	var tokenEnd int

	for l.position < len(l.input) {
		c := l.input[l.position]
		l.position++

		switch {
		case isDigit(c):
			tokenEnd = l.position
			if l.position == len(l.input) || !isDigit(l.input[l.position]) {
				lval.String = l.input[tokenStart:tokenEnd]
				return NUMBER
			}
		case isLetter(c):
			for ; l.position < len(l.input) && isLetter(l.input[l.position]) || isDigit(l.input[l.position]); l.position++ {
				tokenEnd = l.position
			}
			lval.String = l.input[tokenStart:tokenEnd]
			return VARIABLE
		case isSpace(c): //ignore spaces
			tokenStart = l.position
		default:
			return int(c)
		}
	}
	return EOF
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}
