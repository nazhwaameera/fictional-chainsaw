package main

import "fmt"

func main() {
	// 1. Menampilkan nilai i: 21
	i := 21
	fmt.Printf("%v \n", i) // menggunakan function Printf dan flag %v yang akan mengeluarkan value dalam default formatnya.
	// 2. Menampilkan tipe data dari variabel i
	fmt.Printf("%T \n", i) // menggunakan function Printf dan flag %T untuk mengeluarkan tipe data variabel.
	// 3. Menampilkan tanda %
	fmt.Printf("%% \n") //menggunakan function Printf dan flag %% yang mengeluarkan tanda persen secara literal.
	// 4. Menampilkan nilai boolean j: true
	j := true
	fmt.Printf("%t \n", j) // menggunakan function Printf dan flag %t untuk menampilkan nilai boolean.

	fmt.Printf("\n")

	// 5. Menampilkan nilai base 2 dari k: 10101
	k := 21
	fmt.Printf("%b \n", k) // menggunakan function Printf dan flag %b untuk menampilkan base 2 dari k.
	// 6. Menampilkan unicode Rusia: Я (ya)
	fmt.Printf("%c \n", 1071) // menggunakan function Printf dan flag %c untuk menampilkan unicode "Я"
	// 7. Menampilkan nilai base 10: 21
	fmt.Printf("%d \n", 21) // menggunakan flag %d untuk menampilkan nilai base 10 21
	// 8. Menampilkan nilai base 8: 25
	fmt.Printf("%o \n", 21) // menggunakan flag %o untuk menampilkan nilai base 8 25
	// 9. Menampilkan nilai base 16: f
	fmt.Printf("%x \n", 15) // menggunakan flag %x untuk menampilkan nilai base 16 F
	// 10. Menampilkan nilai base 16: F
	fmt.Printf("%X \n", 15) // menggunakan flag %X untuk menampilkan nilai base 16 F
	// 11. Menampilkan unicode karakter Я: U+042F
	fmt.Printf("%U \n", 'Я') // menggunakan flag %U untuk menampilkan Unicode dari "Я"

	fmt.Printf("\n")

	// 12. Menampilkan float: 123.456000
	fmt.Printf("%f \n", 123.456) // menggunakan flag %f untuk menampilkan decimal point tanpa eksponen.
	// 13. Menampilkan float scientific: 1.234560E+02
	fmt.Printf("%E \n", 123.456) // menggunakan flag %E untuk scientific notation.
}
