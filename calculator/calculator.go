package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	yyErrorVerbose = true
	interpreter := interpreter{}

	interpreter.vars = make(map[string]float64)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Bye.")
			return
		}

		interpreter.position = 0
		interpreter.input = input
		interpreter.evaluationFailed = false

		yyParse(&interpreter)
	}
}

const EOF = 0

type interpreter struct {
	input            string
	position         int
	evaluationFailed bool
	vars             map[string]float64
}

func (i *interpreter) Error(e string) {
	fmt.Println(e)
	i.evaluationFailed = true
}

type tokenDef struct {
	regex *regexp.Regexp
	token int
}

var tokens = []tokenDef{
	tokenDef{
		regex: regexp.MustCompile(`^[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?`),
		token: NUMBER,
	},
	tokenDef{
		regex: regexp.MustCompile(`^[_a-zA-Z][_a-zA-Z0-9]*`),
		token: IDENTIFIER,
	},
}

func (l *interpreter) Lex(lval *yySymType) int {
	// Skip spaces
	for ; l.position < len(l.input) && isSpace(l.input[l.position]); l.position++ {
	}
	if l.position == len(l.input) {
		return EOF
	}
	for _, tokDef := range tokens {
		str := tokDef.regex.FindString(l.input[l.position:])
		if str != "" {
			lval.String = str
			l.position += len(str)
			return tokDef.token
		}
	}
	ret := int(l.input[l.position])
	l.position++
	return ret
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n'
}
