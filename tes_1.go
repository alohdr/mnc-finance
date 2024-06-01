package main

import (
	"bufio"
	"fmt"
	"os"
)

func isEqual(s1, s2 string) bool {
	return s1 == s2
}

func main() {
	var n int
	fmt.Println("Masukkan jumlah string:")
	fmt.Scanln(&n)

	fmt.Println("Masukkan string-string:")
	scanner := bufio.NewScanner(os.Stdin)
	s := make([]string, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		s[i] = scanner.Text()
	}

	foundMap := make(map[string]bool)
	found := false

	for i := 0; i < n; i++ {
		if foundMap[s[i]] {
			continue
		}

		for j := i + 1; j < n; j++ {
			if isEqual(s[i], s[j]) {
				if !found {
					found = true
					fmt.Printf("%d", i+1)
				}
				fmt.Printf(" %d", j+1)
				foundMap[s[i]] = true
				break
			}
		}
	}

	if !found {
		fmt.Println("false")
	} else {
		fmt.Println()
	}
}
