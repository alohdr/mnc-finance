package main

import (
	"bufio"
	"fmt"
	"os"
)

func isValid(input string) bool {
	if len(input) < 1 || len(input) > 4096 {
		return false
	}

	stack := []rune{}
	brackets := map[rune]rune{
		'>': '<',
		'}': '{',
		']': '[',
	}

	for _, char := range input {
		switch char {
		case '<', '{', '[':
			stack = append(stack, char)
		case '>', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != brackets[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan string untuk divalidasi (hanya karakter <>{}[]):")

	for scanner.Scan() {
		input := scanner.Text()
		if isValid(input) {
			fmt.Println("TRUE")
		} else {
			fmt.Println("FALSE")
		}
		fmt.Println("Masukkan string lain untuk divalidasi (atau tekan Ctrl+C untuk keluar):")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error membaca input:", err)
	}
}
