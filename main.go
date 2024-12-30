package main

import (
	"fmt"
	"sort"
)

// Struktur data Pasien
type Pasien struct {
	ID   int
	Nama string
	Usia int
}

// Struktur data Pertanyaan dan Tanggapan
type Pertanyaan struct {
	ID         int
	PasienID   int
	Pertanyaan string
	Tanggapan  []Tanggapan
	Tags       []string // Menambahkan tag untuk pertanyaan
}

type Tanggapan struct {
	ID   int
	Dari string
	Teks string
}

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
		fmt.Println("6. Cari pertanyaan berdasarkan tag")
		fmt.Println("7. Lihat topik terurut berdasarkan jumlah pertanyaan")
		fmt.Println("8. Hitung total tanggapan")
		fmt.Println("9. Keluar")

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
			var tags string
			fmt.Print("Masukkan ID pasien: ")
			fmt.Scanln(&pasienID)
			fmt.Print("Masukkan pertanyaan: ")
			fmt.Scanln(&pertanyaan)
			fmt.Print("Masukkan tag (pisahkan dengan koma): ")
			fmt.Scanln(&tags)
			tagList := parseTags(tags)
			ajukanPertanyaan(pasienID, pertanyaan, tagList)
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
			lihatForum()
		case 6:
			var tag string
			fmt.Print("Masukkan tag yang ingin dicari: ")
			fmt.Scanln(&tag)
			cariPertanyaan(tag)
		case 7:
			tampilkanTopikTerurut()
		case 8:
			totalTanggapan := hitungTanggapan(pertanyaanArr[:pertanyaanCount], 0)
			fmt.Printf("Total tanggapan: %d\n", totalTanggapan)
		case 9:
			fmt.Println("Program selesai. Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// Fungsi untuk mendaftar pasien
func daftarPasien(nama string, usia int) Pasien {
	pasien := Pasien{
		ID:   pasienCount + 1,
		Nama: nama,
		Usia: usia,
	}

	pasienArr[pasienCount] = pasien
	pasienCount++
	return pasien
}

// Fungsi untuk mengajukan pertanyaan
func ajukanPertanyaan(pasienID int, pertanyaan string, tags []string) Pertanyaan {
	pertanyaanBaru := Pertanyaan{
		ID:         pertanyaanCount + 1,
		PasienID:   pasienID,
		Pertanyaan: pertanyaan,
		Tags:       tags,
	}

	pertanyaanArr[pertanyaanCount] = pertanyaanBaru
	pertanyaanCount++
	return pertanyaanBaru
}

// Fungsi untuk memberikan tanggapan
func berikanTanggapan(pertanyaanID int, dari, tanggapan string) {
	tanggapanBaru := Tanggapan{
		ID:   tanggapanCount + 1,
		Dari: dari,
		Teks: tanggapan,
	}

	for i := 0; i < pertanyaanCount; i++ {
		if pertanyaanArr[i].ID == pertanyaanID {
			pertanyaanArr[i].Tanggapan = append(pertanyaanArr[i].Tanggapan, tanggapanBaru)
			tanggapanCount++
			break
		}
	}
}

// Fungsi untuk melihat forum konsultasi
func lihatForum() {
	fmt.Println("\nData Pertanyaan dan Tanggapan:")
	for _, pertanyaan := range pertanyaanArr[:pertanyaanCount] {
		fmt.Printf("ID: %d, Pasien ID: %d, Pertanyaan: %s\n", pertanyaan.ID, pertanyaan.PasienID, pertanyaan.Pertanyaan)
		for _, tanggapan := range pertanyaan.Tanggapan {
			fmt.Printf("   Tanggapan dari %s: %s\n", tanggapan.Dari, tanggapan.Teks)
		}
	}
}

// Fungsi untuk mencari pertanyaan berdasarkan tag
func cariPertanyaan(tag string) {
	fmt.Printf("\nPertanyaan dengan tag '%s':\n", tag)
	for _, pertanyaan := range pertanyaanArr[:pertanyaanCount] {
		if contains(pertanyaan.Tags, tag) {
			fmt.Printf("ID: %d, Pertanyaan: %s\n", pertanyaan.ID, pertanyaan.Pertanyaan)
		}
	}
}

// Fungsi untuk menampilkan topik terurut berdasarkan jumlah pertanyaan
func tampilkanTopikTerurut() {
	tagCount := make(map[string]int)
	for _, pertanyaan := range pertanyaanArr[:pertanyaanCount] {
		for _, tag := range pertanyaan.Tags {
			tagCount[tag]++
		}
	}

	tags := make([]string, 0, len(tagCount))
	for tag := range tagCount {
		tags = append(tags, tag)
	}

	sort.Slice(tags, func(i, j int) bool {
		return tagCount[tags[i]] > tagCount[tags[j]]
	})

	fmt.Println("\nTopik terurut berdasarkan jumlah pertanyaan:")
	for _, tag := range tags {
		fmt.Printf("Tag: %s, Jumlah: %d\n", tag, tagCount[tag])
	}
}

// Fungsi utilitas untuk memeriksa keberadaan elemen di array
func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// Fungsi utilitas untuk memparse tag
func parseTags(tagStr string) []string {
	tags := []string{}
	current := ""
	for _, char := range tagStr {
		if char == ',' {
			tags = append(tags, current)
			current = ""
		} else {
			current += string(char)
		}
	}
	if current != "" {
		tags = append(tags, current)
	}
	return tags
}

// Rekursif untuk menghitung total tanggapan
func hitungTanggapan(pertanyaanArr []Pertanyaan, index int) int {
	if index == len(pertanyaanArr) {
		return 0
	}
	return len(pertanyaanArr[index].Tanggapan) + hitungTanggapan(pertanyaanArr, index+1)
}
