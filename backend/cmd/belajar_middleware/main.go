package main

import (
	"fmt"
	"strings"
)

func main() {
	kata := "Halo saya aku"
	hitungA(kata)
	hitungKata(kata)
	hitungBadWord(kata)
	uniqueCharacter(kata)
}

func hitungA(content string) {
	count := 0
	for _, v := range content {
		if v == 'a' || v == 'A' {
			count++
		}
	}
	fmt.Println("Jumlah a =", count)
}

func hitungKata(content string) {
	kata := strings.Split(content, " ")
	fmt.Println("Jumlah kata =", len(kata))
}

func hitungBadWord(content string) {
	count := 0
	kata := strings.Split(content, " ")
	for _, v := range kata {
		// if v == "saya" || v == "aku" {
		// 	count++
		// }
		switch v {
		case "saya":
			count++
		case "aku":
			count++
		default:
		}
	}
	fmt.Println("Jumlah badword (saya dan aku) =", count)
}

func uniqueCharacter(content string) {
	lower := strings.ToLower(content)
	var tmp map[rune]bool = make(map[rune]bool)
	for _, v := range lower {
		if tmp[v] == true {
			continue
		}
		tmp[v] = true
	}

	fmt.Println("Jumlah unique karakter =", len(tmp))
}
