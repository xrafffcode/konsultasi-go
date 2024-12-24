// pasien.go
package main

// Struktur data Pasien
type Pasien struct {
	ID   int
	Nama string
	Usia int
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
