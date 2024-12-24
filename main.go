// main.go
package main

import (
	"fmt"
)

// Array statis untuk data pasien dan pertanyaan
var pasienArr [100]Pasien
var pertanyaanArr [100]Pertanyaan

// Variabel untuk menghitung jumlah data
var pasienCount, pertanyaanCount, tanggapanCount int

func main() {
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Daftar pasien")
		fmt.Println("2. Ajukan pertanyaan")
		fmt.Println("3. Berikan tanggapan")
		fmt.Println("4. Lihat data pasien")
		fmt.Println("5. Lihat data pertanyaan dan tanggapan")
		fmt.Println("6. Keluar")

		var input int
		fmt.Print("Masukkan pilihan: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			var nama string
			var usia int
			fmt.Print("Masukkan nama: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan usia: ")
			fmt.Scanln(&usia)
			daftarPasien(nama, usia)
		case 2:
			var pasienID int
			var pertanyaan string
			fmt.Print("Masukkan ID pasien: ")
			fmt.Scanln(&pasienID)
			fmt.Print("Masukkan pertanyaan: ")
			fmt.Scanln(&pertanyaan)
			ajukanPertanyaan(pasienID, pertanyaan)
		case 3:
			var pertanyaanID int
			var dari, tanggapan string
			fmt.Print("Masukkan ID pertanyaan: ")
			fmt.Scanln(&pertanyaanID)
			fmt.Print("Masukkan nama yang memberikan tanggapan: ")
			fmt.Scanln(&dari)
			fmt.Print("Masukkan tanggapan: ")
			fmt.Scanln(&tanggapan)
			berikanTanggapan(pertanyaanID, dari, tanggapan)
		case 4:
			fmt.Println("Data Pasien:")
			for _, pasien := range pasienArr[:pasienCount] {
				fmt.Printf("ID: %d, Nama: %s, Usia: %d\n", pasien.ID, pasien.Nama, pasien.Usia)
			}
		case 5:
			fmt.Println("\nData Pertanyaan dan Tanggapan:")
			for _, pertanyaan := range pertanyaanArr[:pertanyaanCount] {
				fmt.Printf("ID: %d, Pasien ID: %d, Pertanyaan: %s\n", pertanyaan.ID, pertanyaan.PasienID, pertanyaan.Pertanyaan)
				for _, tanggapan := range pertanyaan.Tanggapan {
					fmt.Printf("   Tanggapan dari %s: %s\n", tanggapan.Dari, tanggapan.Teks)
				}
			}
		case 6:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
