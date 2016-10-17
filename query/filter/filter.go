package filter

import (
	"fmt"
)

const (
	LHS = iota
	RHS
	OP
	CLOSURE
)

var closures = map[string]string {
	"and": ";",
	"or" : ",",
	"("  : "(",
	")"  : ")",
}

func isClosure(s string) bool {
	if closures[s] == "" {
		return false
	}

	return true
}

// hasClosure returns true if and only if
// the input strings begins or ends with a 
// closing parenthesis.
func hasClosure(s string) bool {
	//fmt.Print("Has")
	if len(s) == 0 {
		return false
	}

	b := []byte(s)

	if b[0] != '(' && b[len(b)-1] != ')' {
		return false
	}

	return true
}

var operators = map[string]string {
	"eq":  "==",
	"ne":  "!=",
	"and": ";",
	"or":  ",",
	"gt":  "=gt=",
	"ge":  "=ge=",
	"lt":  "=lt=",
	"le":  "=le=",
}

func isOperator(s string) bool {
	if operators[s] == "" {
		return false
	}

	return true
}


func translateOP(s string) (string, error){
	tr := operators[s]

	if tr == "" {
		return "", fmt.Errorf("Filter OPERATOR: \"%s\" isn't an operator.")
	}

	return tr, nil
}

