package main

import "fmt"

func main() {
	// buatlah sebuah program go dengan implementasi perulangan for dan kombinasi if-else dengan expected output seperti yang telah diberikan.
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)

		if i == 4 {
			for j := 0; j < 11; j++ {
				if j == 5 {
					word := "САШАРВО"
					for index, char := range word {
						fmt.Printf("character %U %q starts at byte positions %d\n", char, char, index)
					}
					continue
				}
				fmt.Println("Nilai j =", j)

			}
		}
	}
}
