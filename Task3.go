package main

import (
	"bufio"
	"fmt"
	"os"
)

func isBalancedBrackets(s string) string {
	stack := []rune{}
	bracketPairs := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[char] {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Output:", isBalancedBrackets(input))
}
