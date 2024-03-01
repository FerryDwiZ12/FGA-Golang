package main

import (
	"fmt"
	"os"
)

// Struct untuk merepresentasikan data teman
type Teman struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Fungsi untuk mendapatkan data teman berdasarkan absen
func getTemanByAbsen(absen int) Teman {
	// Dalam implementasi nyata, data teman dapat diambil dari database atau sumber data lainnya
	// Untuk keperluan contoh, kita buat beberapa data teman secara statis
	temanData := map[int]Teman{
		1: {1, "Jono", "Jakarta", "Developer", "Suka dengan pemrograman"},
		2: {2, "Jane", "Bandung", "Designer", "Ingin mengembangkan keterampilan desain"},
		3: {3, "Ngopo", "Surabaya", "Analyst", "Tertarik dengan analisis data"},
	}

	teman, found := temanData[absen]
	if !found {
		fmt.Println("Teman dengan absen", absen, "tidak ditemukan.")
		os.Exit(1)
	}

	return teman
}

// Fungsi untuk menampilkan data teman
func printTeman(teman Teman) {
	fmt.Println("Data Teman:")
	fmt.Println("Absen   :", teman.Absen)
	fmt.Println("Nama    :", teman.Nama)
	fmt.Println("Alamat  :", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan  :", teman.Alasan)
}

func main() {
	// Mengecek apakah argumen absen disediakan
	if len(os.Args) < 2 {
		fmt.Println("Gunakan: go run biodata.go <absen>")
		os.Exit(1)
	}

	// Mengambil argumen absen dari command line
	absen := os.Args[1]

	// Mengonversi argumen absen menjadi integer
	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Absen harus berupa angka.")
		os.Exit(1)
	}

	// Mendapatkan data teman berdasarkan absen
	teman := getTemanByAbsen(absenInt)

	// Menampilkan data teman
	printTeman(teman)
}
