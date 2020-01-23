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

		interpreter.input = input
		interpreter.evaluationFailed = false

		yyParse(&interpreter)

		if !interpreter.evaluationFailed {
			interpreter.eval(interpreter.parseResult)
		}
	}
}

const EOF = 0

type interpreter struct {
	input            string
	evaluationFailed bool
	vars             map[string]float64
	parseResult      expr
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
	// Skip spaces.
	for ; len(l.input) > 0 && isSpace(l.input[0]); l.input = l.input[1:] {
	}

	// Check if the input has ended.
	if len(l.input) == 0 {
		return EOF
	}

	// Check if one of the regular expressions matches.
	for _, tokDef := range tokens {
		str := tokDef.regex.FindString(l.input)
		if str != "" {
			// Pass string content to the parser.
			lval.String = str
			l.input = l.input[len(str):]
			return tokDef.token
		}
	}

	// Otherwise return the next letter.
	ret := int(l.input[0])
	l.input = l.input[1:]
	return ret
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n'
}
