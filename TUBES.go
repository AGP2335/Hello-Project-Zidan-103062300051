package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mahasiswa struct {
	Nama     string
	Umur     int
	Jurusan  string
	Alamat   string
	Nilai    int
}

var mahasiswas []Mahasiswa
var scanner = bufio.NewScanner(os.Stdin)
var jurusanList = []string{"Sistem Informasi", "Teknologi Informasi"}

func main() {
	for {
		displayMenu()
		pilihan := scanInput("Masukan pilihan: ")
		switch pilihan {
		case "1":
			daftar()
		case "2":
			editMahasiswa()
		case "3":
			cekHasil()
		case "4":
			middleware()
		case "5":
			fmt.Println("Keluar dari program...")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func cekHasil() {
	if len(mahasiswas) == 0 {
		fmt.Println("Belum ada data calon mahasiswa yang terdaftar.")
		return
	}

	nama := scanInput("Masukan nama calon mahasiswa yang ingin dicari: ")

	ditemukan := false
	for _, mhs := range mahasiswas {
		if mhs.Nama == nama {
			if mhs.Nilai > 75 {
				fmt.Println("Selamat! Kamu diterima di universitas kita!")
			} else {
				fmt.Println("Maaf, kamu belum diterima. Tetap semangat ya!")
			}
			ditemukan = true
			break
		}
	}

	if !ditemukan {
		fmt.Println("Calon mahasiswa dengan nama tersebut tidak ditemukan.")
	}
}

func daftar() {
	var umur, nilai int

	fmt.Println("<< Menu Pendaftaran Mahasiswa >>")

	nama := scanInput("Masukan Nama: ")

	fmt.Print("Masukan umur: ")
	fmt.Scanln(&umur)

	fmt.Println("Pilih Jurusan: ")
	for i, jurusan := range jurusanList {
		fmt.Printf("%d. %s\n", i+1, jurusan)
	}
	jurusanPilihan := scanInput("Masukan pilihan jurusan (1/2): ")
	jurusanIndex := -1

	if jurusanPilihan == "1" || jurusanPilihan == "2" {
		jurusanIndex = strings.Index("12", jurusanPilihan)
	} else {
		fmt.Println("Jurusan tidak valid.")
		return
	}

	alamat := scanInput("Masukan Alamat: ")

	fmt.Print("Masukan Nilai: ")
	fmt.Scan(&nilai)

	if nama == "" || jurusanIndex < 0 || umur <= 0 || alamat == "" {
		fmt.Println("Error: Data harus diisi dengan benar!")
		return
	}

	mahasiswaBaru := Mahasiswa{
		Nama:    nama,
		Umur:    umur,
		Jurusan: jurusanList[jurusanIndex],
		Alamat:  alamat,
		Nilai:   nilai,
	}

	mahasiswas = append(mahasiswas, mahasiswaBaru)

	fmt.Print("Data Mahasiswa Berhasil Disimpan!\n")
	fmt.Printf("Jumlah Mahasiswa Terdaftar: %d\n", len(mahasiswas))
}

func data() {
	if len(mahasiswas) == 0 {
		fmt.Println("Belum ada data mahasiswa yang terdaftar.")
		return
	}

	fmt.Println("<< Data Mahasiswa >>")
	for i, mhs := range mahasiswas {
		fmt.Printf("Mahasiswa %d\n", i+1)
		fmt.Printf("Nama    : %s\n", mhs.Nama)
		fmt.Printf("Umur    : %d\n", mhs.Umur)
		fmt.Printf("Jurusan : %s\n", mhs.Jurusan)
		fmt.Printf("Alamat  : %s\n", mhs.Alamat)
		fmt.Printf("Nilai   : %d\n", mhs.Nilai)
		fmt.Println("-----------------------")
	}
}

func admin() {
	for {
		var pilihan int

		fmt.Println("<< Menu Admin >>")
		fmt.Println(" 1. Lihat Data")
		fmt.Println(" 2. Cari Data Calon Mahasiswa")
		fmt.Println(" 3. Urutkan Data Calon Mahasiswa")
		fmt.Println(" 4. Edit Data Calon Mahasiswa")
		fmt.Println(" 5. Hapus Data Calon Mahasiswa")
		fmt.Println(" 6. Tampilkan Mahasiswa yang Keterima")
		fmt.Println(" 7. Tampilkan Mahasiswa yang Tidak Keterima")
		fmt.Println(" 8. Kembali")

		fmt.Print("Masukan Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			data()
		case 2:
			cariCalonMahasiswa()
		case 3:
			urutMahasiswa()
		case 4:
			editMahasiswa()
		case 5:
			hapusMahasiswa()
		case 6:
			tampilkanMahasiswaKeterima()
		case 7:
			tampilkanMahasiswaTidakKeterima()
		case 8:
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func urutMahasiswa() {
	for {
		var pilihan int
		fmt.Println("<< Menu Urutkan Calon Mahasiswa >>")
		fmt.Println(" 1. Urutkan Data Calon Mahasiswa Berdasarkan Nilai")
		fmt.Println(" 2. Urutkan Data Calon Mahasiswa Berdasarkan Nama")
		fmt.Println(" 3. Urutkan Data Calon Mahasiswa Berdasarkan Jurusan")
		fmt.Println(" 4. Kembali")

		fmt.Print("Masukan Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			sortedStudents := urutkanCalonMahasiswa(mahasiswas)
			fmt.Println("Data mahasiswa setelah diurutkan berdasarkan nilai (Descending):")
			for i, mhs := range sortedStudents {
				fmt.Printf("Mahasiswa %d\n", i+1)
				fmt.Printf("Nama    : %s\n", mhs.Nama)
				fmt.Printf("Umur    : %d\n", mhs.Umur)
				fmt.Printf("Jurusan : %s\n", mhs.Jurusan)
				fmt.Printf("Alamat  : %s\n", mhs.Alamat)
				fmt.Printf("Nilai   : %d\n", mhs.Nilai)
				fmt.Println("-----------------------")
			}
		case 2:
			sortedStudents := urutkanCalonMahasiswaNama(mahasiswas)
			fmt.Println("Data mahasiswa setelah diurutkan berdasarkan nama:")
			for i, mhs := range sortedStudents {
				fmt.Printf("Mahasiswa %d\n", i+1)
				fmt.Printf("Nama    : %s\n", mhs.Nama)
				fmt.Printf("Umur    : %d\n", mhs.Umur)
				fmt.Printf("Jurusan : %s\n", mhs.Jurusan)
				fmt.Printf("Alamat  : %s\n", mhs.Alamat)
				fmt.Printf("Nilai   : %d\n", mhs.Nilai)
				fmt.Println("-----------------------")
			}
		case 3:
			sortedStudents := urutkanCalonMahasiswaJurusan(mahasiswas)
			fmt.Println("Data mahasiswa setelah diurutkan berdasarkan jurusan:")
			for i, mhs := range sortedStudents {
				fmt.Printf("Mahasiswa %d\n", i+1)
				fmt.Printf("Nama    : %s\n", mhs.Nama)
				fmt.Printf("Umur    : %d\n", mhs.Umur)
				fmt.Printf("Jurusan : %s\n", mhs.Jurusan)
				fmt.Printf("Alamat  : %s\n", mhs.Alamat)
				fmt.Printf("Nilai   : %d\n", mhs.Nilai)
				fmt.Println("-----------------------")
			}
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func middleware() {
	var pin int

	fmt.Print("Masukan pin: ")
	fmt.Scan(&pin)

	if pin == 2323 {
		admin()
	} else {
		fmt.Println("Pin salah")
	}
}

func cariCalonMahasiswa() {
	if len(mahasiswas) == 0 {
		fmt.Println("Belum ada data calon mahasiswa yang terdaftar.")
		return
	}

	nama := scanInput("Masukan nama calon mahasiswa yang ingin dicari: ")

	ditemukan := false
	for i, mhs := range mahasiswas {
		if mhs.Nama == nama {
			fmt.Printf("Mahasiswa %d\n", i+1)
			fmt.Printf("Nama    : %s\n", mhs.Nama)
			fmt.Printf("Umur    : %d\n", mhs.Umur)
			fmt.Printf("Jurusan : %s\n", mhs.Jurusan)
			fmt.Printf("Alamat  : %s\n", mhs.Alamat)
			fmt.Printf("Nilai   : %d\n", mhs.Nilai)
			fmt.Println("-----------------------")
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Calon mahasiswa dengan nama tersebut tidak ditemukan.")
	}
}

func editMahasiswa() {
	if len(mahasiswas) == 0 {
		fmt.Println("Belum ada data calon mahasiswa yang terdaftar.")
		return
	}

	nama := scanInput("Masukan nama calon mahasiswa yang ingin diedit: ")

	ditemukan := false
	for i, mhs := range mahasiswas {
		if mhs.Nama == nama {
			fmt.Printf("Mengedit data mahasiswa %d\n", i+1)
			mahasiswas[i] = inputMahasiswaData(mahasiswas[i])
			fmt.Println("Data mahasiswa berhasil diedit.")
			ditemukan = true
			break
		}
	}

	if !ditemukan {
		fmt.Println("Calon mahasiswa dengan nama tersebut tidak ditemukan.")
	}
}

func hapusMahasiswa() {
	if len(mahasiswas) == 0 {
		fmt.Println("Belum ada data calon mahasiswa yang terdaftar.")
		return
	}

	nama := scanInput("Masukan nama calon mahasiswa yang ingin dihapus: ")

	ditemukan := false
	for i, mhs := range mahasiswas {
		if mhs.Nama == nama {
			mahasiswas = append(mahasiswas[:i], mahasiswas[i+1:]...)
			fmt.Println("Data mahasiswa berhasil dihapus.")
			ditemukan = true
			break
		}
	}

	if !ditemukan {
		fmt.Println("Calon mahasiswa dengan nama tersebut tidak ditemukan.")
	}
}

func tampilkanMahasiswaKeterima() {
	fmt.Println("<< Mahasiswa yang Keterima >>")
	ditemukan := false
	for i, mhs := range mahasiswas {
		if mhs.Nilai > 75 {
			fmt.Printf("Mahasiswa %d\n", i+1)
			fmt.Printf("Nama    : %s\n", mhs.Nama)
			fmt.Printf("Umur    : %d\n", mhs.Umur)
			fmt.Printf("Jurusan : %s\n", mhs.Jurusan)
			fmt.Printf("Alamat  : %s\n", mhs.Alamat)
			fmt.Printf("Nilai   : %d\n", mhs.Nilai)
			fmt.Println("-----------------------")
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada mahasiswa yang keterima.")
	}
}

func tampilkanMahasiswaTidakKeterima() {
	fmt.Println("<< Mahasiswa yang Tidak Keterima >>")
	ditemukan := false
	for i, mhs := range mahasiswas {
		if mhs.Nilai <= 75 {
			fmt.Printf("Mahasiswa %d\n", i+1)
			fmt.Printf("Nama    : %s\n", mhs.Nama)
			fmt.Printf("Umur    : %d\n", mhs.Umur)
			fmt.Printf("Jurusan : %s\n", mhs.Jurusan)
			fmt.Printf("Alamat  : %s\n", mhs.Alamat)
			fmt.Printf("Nilai   : %d\n", mhs.Nilai)
			fmt.Println("-----------------------")
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada mahasiswa yang tidak keterima.")
	}
}

func urutkanCalonMahasiswa(data []Mahasiswa) []Mahasiswa {
	sortedData := make([]Mahasiswa, len(data))
	copy(sortedData, data)
	
	for i := 0; i < len(sortedData)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(sortedData); j++ {
			if sortedData[j].Nilai > sortedData[minIndex].Nilai {
				minIndex = j
			}
		}
		sortedData[i], sortedData[minIndex] = sortedData[minIndex], sortedData[i]
	}
	
	return sortedData
}

func urutkanCalonMahasiswaNama(data []Mahasiswa) []Mahasiswa {
	sortedData := make([]Mahasiswa, len(data))
	copy(sortedData, data)
	
	for i := 0; i < len(sortedData)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(sortedData); j++ {
			if sortedData[j].Nama < sortedData[minIndex].Nama {
				minIndex = j
			}
		}
		sortedData[i], sortedData[minIndex] = sortedData[minIndex], sortedData[i]
	}
	
	return sortedData
}

func urutkanCalonMahasiswaJurusan(data []Mahasiswa) []Mahasiswa {
	sortedData := make([]Mahasiswa, len(data))
	copy(sortedData, data)
	
	for i := 0; i < len(sortedData)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(sortedData); j++ {
			if sortedData[j].Jurusan < sortedData[minIndex].Jurusan {
				minIndex = j
			}
		}
		sortedData[i], sortedData[minIndex] = sortedData[minIndex], sortedData[i]
	}
	
	return sortedData
}

func inputMahasiswaData(existing Mahasiswa) Mahasiswa {
	var umur, nilai int
	var jurusanIndex int

	nama := scanInput(fmt.Sprintf("Masukan Nama (%s): ", existing.Nama))
	if nama == "" {
		nama = existing.Nama
	}

	fmt.Printf("Masukan umur (%d): ", existing.Umur)
	fmt.Scanln(&umur)
	if umur == 0 {
		umur = existing.Umur
	}

	fmt.Println("Pilih Jurusan: ")
	for i, jurusan := range jurusanList {
		fmt.Printf("%d. %s\n", i+1, jurusan)
	}
	fmt.Printf("Masukan pilihan jurusan (1/2) [%s]: ", existing.Jurusan)
	fmt.Scan(&jurusanIndex)
	if jurusanIndex == 0 {
		jurusanIndex = strings.Index("12", existing.Jurusan)
	}

	alamat := scanInput(fmt.Sprintf("Masukan Alamat (%s): ", existing.Alamat))
	if alamat == "" {
		alamat = existing.Alamat
	}

	fmt.Printf("Masukan Nilai (%d): ", existing.Nilai)
	fmt.Scan(&nilai)
	if nilai == 0 {
		nilai = existing.Nilai
	}

	return Mahasiswa{
		Nama:    nama,
		Umur:    umur,
		Jurusan: jurusanList[jurusanIndex-1],
		Alamat:  alamat,
		Nilai:   nilai,
	}
}

func scanInput(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func displayMenu() {
	fmt.Println("\n<< PENDAFTARAN MAHASISWA >>")
	fmt.Println(" 1. Daftar")
	fmt.Println(" 2. Edit Data Calon Mahasiswa")
	fmt.Println(" 3. Cek Hasil")
	fmt.Println(" 4. Admin")
	fmt.Println(" 5. Keluar")
}