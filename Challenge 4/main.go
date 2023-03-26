package main

import (
	"fmt"
	"log"
)

type Peserta struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	var peserta = []Peserta{
		{Nama: "Nazhwa Ameera H.", Alamat: "Surabaya", Pekerjaan: "Mahasiswa", Alasan: "Ingin mendapatkan pengalaman dan ilmu"},
		{Nama: "Rosa Amalia", Alamat: "Surabaya", Pekerjaan: "Mahasiswa", Alasan: "Mengisi waktu luang"},
		{Nama: "Nurul Izzatil Ulum", Alamat: "Surabaya", Pekerjaan: "Mahasiswa", Alasan: "Mencari kesibukan lain"},
		{Nama: "Zahrotul Adillah", Alamat: "Surabaya", Pekerjaan: "Mahasiswa", Alasan: "Ingin belajar bahasa Go"},
	}

	fmt.Printf("Selamat datang di pusat data kelas Scalable Web Service with Golang.\n")

	// read number
	fmt.Println("Masukkan nomor ID yang ingin dicari:")
	var id int
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%d", id)

	// error handling jika data yang dimasukkan tidak ada dalam database.
	if id > len(peserta) || id == 0 {
		fmt.Println("Data tidak ditemukan. Silakan masukkan data yang valid.")
	} else {
		searchByID(id, peserta)
	}

}

func searchByID(id int, peserta []Peserta) {
	for i := 0; i < id; i++ {
		if i == id-1 {
			fmt.Printf("%+v \n", peserta[i])
		}
	}
}
