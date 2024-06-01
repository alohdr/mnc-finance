package main

import (
	"fmt"
	"sort"
)

func calculate(totalBelanja, uangDibayar int) (int, map[int]int, bool) {
	if uangDibayar < totalBelanja {
		return 0, nil, false
	}

	pecahan := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	kembalian := uangDibayar - totalBelanja

	kembalian = (kembalian / 100) * 100

	result := make(map[int]int)
	for _, p := range pecahan {
		if kembalian >= p {
			result[p] = kembalian / p
			kembalian = kembalian % p
		}
	}

	return (uangDibayar - totalBelanja) / 100 * 100, result, true
}

func main() {
	var totalBelanja int
	fmt.Print("Total belanja seorang customer: Rp ")
	fmt.Scanln(&totalBelanja)

	var uangDibayar int
	fmt.Print("Total belanja seorang customer: Rp ")
	fmt.Scanln(&uangDibayar)

	fmt.Println("inputan = ", totalBelanja)
	fmt.Println("inputan = ", uangDibayar)

	kembalian, pecahan, valid := calculate(totalBelanja, uangDibayar)
	if !valid {
		fmt.Println("False, kurang bayar")
	} else {
		fmt.Printf("Kembalian yang harus diberikan kasir: %d, dibulatkan menjadi %d\n", uangDibayar-totalBelanja, kembalian)
		fmt.Println("Pecahan uang:")
		keys := make([]int, 0, len(pecahan))
		for k := range pecahan {
			keys = append(keys, k)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
		for _, p := range keys {
			if p >= 1000 {
				fmt.Printf("%d lembar %d\n", pecahan[p], p)
			} else {
				fmt.Printf("%d koin %d\n", pecahan[p], p)
			}
		}
	}
}
