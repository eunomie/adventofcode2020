package main

import (
	"fmt"
	"strconv"

	"github.com/eunomie/adventofcode2020/lib"
)

func main() {
	fomulas := lib.AsStringArray(lib.Input())
	sum := 0
	for _, formula := range fomulas {
		sum += Compute(formula)
	}

	fmt.Println("sum:", sum)
}

func Addition(left, right int) int {
	return left + right
}

func Multiplication(left, right int) int {
	return left * right
}

type Expr struct {
	Value    int
	Operator func(int, int) int
}

func NewExpr(value int) *Expr {
	return &Expr{
		Value: value,
	}
}

func NewIdentity() *Expr {
	return &Expr{
		Value:    1,
		Operator: Multiplication,
	}
}

func (e *Expr) Compute(v int) int {
	return e.Operator(e.Value, v)
}

func Compute(formula string) int {
	stack := []*Expr{
		NewIdentity(),
	}
	var p, l *Expr

	for _, c := range formula {
		switch c {
		case ' ':
			continue
		case '+':
			stack[len(stack)-1].Operator = Addition
		case '*':
			stack[len(stack)-1].Operator = Multiplication
		case '(':
			stack = append(stack, NewIdentity())
		case ')':
			p, l, stack = stack[len(stack)-2], stack[len(stack)-1], stack[:len(stack)-2]
			stack = append(stack, NewExpr(p.Compute(l.Value)))
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			v, _ := strconv.Atoi(string(c))
			l, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, NewExpr(l.Compute(v)))
		}
	}

	return stack[0].Value
}
