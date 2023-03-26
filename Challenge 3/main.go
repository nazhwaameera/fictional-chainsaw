package main

import "fmt"

func main() {
	word := "selamat malam"
	char := make(map[string]int)

	for _, v := range word {
		fmt.Printf("%c\n", v)
		char[string(v)] = char[string(v)] + 1
	}

	fmt.Println(char)
}
