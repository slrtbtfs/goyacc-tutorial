package main

import (
	"fmt"
	"strconv"
)

type expr interface{}

type binaryExpr struct {
	Op  byte
	lhs expr
	rhs expr
}

type unaryExpr struct {
	expr expr
}

type astRoot struct {
	expr expr
}

type parenExpr struct {
	expr expr
}
type variable struct {
	name string
}

type number struct {
	val string
}

type assignment struct {
	variable string
	expr     expr
}

func (i *interpreter) eval(e expr) float64 {
	switch t := e.(type) {
	case *binaryExpr:
		lhs := i.eval(t.lhs)
		rhs := i.eval(t.rhs)

		switch t.Op {
		case '+':
			return lhs + rhs
		case '-':
			return lhs - rhs
		case '*':
			return lhs * rhs
		case '/':
			return lhs / rhs
		default:
			panic("invalid operation")
		}

	case *unaryExpr:
		return -i.eval(t.expr)

	case *astRoot:
		result := i.eval(t.expr)
		if !i.evaluationFailed {
			fmt.Println(result)
		}
		return result

	case *parenExpr:
		return i.eval(t.expr)
	case *variable:
		val, ok := i.vars[t.name]
		if !ok {
			i.Error(fmt.Sprintf("Variable undefined: %s\n", t.name))
		}
		return val

	case *number:
		var err error
		val, err := strconv.ParseFloat(t.val, 64)
		if err != nil {
			i.Error(err.Error())
		}
		return val
	case *assignment:
		result := i.eval(t.expr)
		if !i.evaluationFailed {
			i.vars[t.variable] = result
		}
		return result

	default:
		panic("invalid node type")
	}

}
