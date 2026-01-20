//go:build ignore

package main

import (
    "fmt"
    "strconv"
)

func isOperation(op string) bool {
    if op == "+" ||
       op == "-" ||
       op == "*" ||
       op == "/" {
        return true
    }
    return false
}

func evalOperation(a string, b string, op string) string {
    aInt, _ := strconv.Atoi(a)
    bInt, _ := strconv.Atoi(b)
    switch op {
        case "+":
            return strconv.Itoa(aInt + bInt)
        case "-":
            return strconv.Itoa(aInt - bInt)
        case "*":
            return strconv.Itoa(aInt * bInt)
        case "/":
            return strconv.Itoa(aInt / bInt)
    }
    return ""
}

func evalRPN(tokens []string) int {
    i := 0
	for {
        if i > len(tokens) - 1 {
            break
        }
		if isOperation(tokens[i]) {
			tokens[i-2] = evalOperation(tokens[i-2], 
                                        tokens[i-1], tokens[i])
            copy(tokens[i-1:], tokens[i+1:])
			tokens = tokens[:len(tokens)-2]
            i = i - 2
		} else {
            i++
        }
	}
    res, _ :=  strconv.Atoi(tokens[0])
    return res;
}

func main() {
    fmt.Println(evalRPN([]string{"4","13","5","/","+"}))
}
