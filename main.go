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
	// Simulasi pendaftaran pasien
	pasien1 := daftarPasien("John Doe", 30)
	pasien2 := daftarPasien("Jane Smith", 25)

	// Simulasi pengajuan pertanyaan
	pertanyaan1 := ajukanPertanyaan(pasien1.ID, "Apa itu penyakit diabetes?")
	pertanyaan2 := ajukanPertanyaan(pasien2.ID, "Bagaimana cara menjaga kesehatan jantung?")

	// Simulasi memberikan tanggapan
	berikanTanggapan(pertanyaan1.ID, "Dr. A", "Diabetes adalah penyakit yang disebabkan oleh kadar gula darah yang tinggi.")
	berikanTanggapan(pertanyaan2.ID, "Dr. B", "Olahraga teratur dan pola makan sehat dapat membantu menjaga kesehatan jantung.")

	// Menampilkan pasien yang terdaftar
	fmt.Println("Data Pasien:")
	for _, pasien := range pasienArr[:pasienCount] {
		fmt.Printf("ID: %d, Nama: %s, Usia: %d\n", pasien.ID, pasien.Nama, pasien.Usia)
	}

	// Menampilkan pertanyaan dan tanggapan
	fmt.Println("\nData Pertanyaan dan Tanggapan:")
	for _, pertanyaan := range pertanyaanArr[:pertanyaanCount] {
		fmt.Printf("ID: %d, Pasien ID: %d, Pertanyaan: %s\n", pertanyaan.ID, pertanyaan.PasienID, pertanyaan.Pertanyaan)
		for _, tanggapan := range pertanyaan.Tanggapan {
			fmt.Printf("   Tanggapan dari %s: %s\n", tanggapan.Dari, tanggapan.Teks)
		}
	}
}
