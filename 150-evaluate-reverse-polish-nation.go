//go:build ignore

package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, tok := range tokens {
		switch tok {
		case "+", "-", "*", "/":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch tok {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			}
		default:
			num, _ := strconv.Atoi(tok)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func main() {
	fmt.Println(evalRPN([]string{
		"10","6", "9", "3", "+",
		"-11", "*", "/", "*", "17", "+", "5", "+"}))
}
