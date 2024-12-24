// pertanyaan.go
package main

// Struktur data Pertanyaan dan Tanggapan
type Pertanyaan struct {
	ID         int
	PasienID   int
	Pertanyaan string
	Tanggapan  []Tanggapan
}

type Tanggapan struct {
	ID   int
	Dari string
	Teks string
}

// Fungsi untuk mengajukan pertanyaan
func ajukanPertanyaan(pasienID int, pertanyaan string) Pertanyaan {
	pertanyaanBaru := Pertanyaan{
		ID:         pertanyaanCount + 1,
		PasienID:   pasienID,
		Pertanyaan: pertanyaan,
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
