package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var daftarTeman = map[int]Teman{
	1: {"santi", "tangerang", "Software Engineer", "Ingin belajar lebih banyak tentang pemrograman"},
	2: {"hans", "serpong", "Data Analyst", "Meningkatkan keterampilan dalam analisis data"},
	// input data teman lain di sini
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go <nomor_absen>")
		os.Exit(1)
	}

	absen := os.Args[1]
	var nomorAbsen int
	_, err := fmt.Sscanf(absen, "%d", &nomorAbsen)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		os.Exit(1)
	}

	teman, found := daftarTeman[nomorAbsen]
	if !found {
		fmt.Println("Teman dengan nomor absen tersebut tidak ditemukan mungkin salah orang")
		os.Exit(1)
	}

	tampilkanDataTeman(teman)
}

func tampilkanDataTeman(teman Teman) {
	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}
