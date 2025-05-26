package main

import (
	"fmt"
)

const MAX = 100

type Proyek struct {
	ID       string
	Nama     string
	Klien    string
	Prioritas string // "tinggi", "sedang", "rendah"
	Tanggal  string // Format: YYYY-MM-DD
	Status   string // "sudah" atau "belum"
}

var data [MAX]Proyek
var n int


//function tambahan


// Utility
func swap(a, b *Proyek) {
	temp := *a
	*a = *b
	*b = temp
}

// CRUD

func tambahProyek(p Proyek) {
	if n < MAX {
		data[n] = p
		n++
		fmt.Println("Proyek berhasil ditambahkan.")
	} else {
		fmt.Println("Data penuh!")
	}
}

func tampilkanData() {
	var pilihan1 int
	var inputID, inputNama string
	fmt.Println("Data Proyek:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. ID: %s | Nama Proyek: %s | Klien: %s | Prioritas: %s | Tanggal: %s | Status: %s\n", i+1, data[i].ID, data[i].Nama, data[i].Klien, data[i].Prioritas, data[i].Tanggal, data[i].Status)
	}
	fmt.Println("----------------------------------------------")
	fmt.Println("\n============= Cari/Urutkan =============")
	fmt.Println("\n--------------------------------------------")
		fmt.Println("1. Cari berdasarkan ID")
		fmt.Println("2. Cari berdasarkan nama proyek")
		fmt.Println("3. Urutkan berdasarkan prioritas")
		fmt.Println("4. Urutkan berdasarkan tanggal")
		fmt.Println("Pilih: ")
		fmt.Scan(&pilihan1)

	switch pilihan1{
	case 1:
		fmt.Print("Masukkan ID: ")
		fmt.Scan(&inputID)
		idx := sequentialSearchID(inputID)
		if idx != -1 {
			fmt.Printf("Ditemukan: %+v\n", data[idx])
		} else {
			fmt.Println("Tidak ditemukan.")
		}
	case 2:
		fmt.Print("Masukkan nama proyek: ")
		fmt.Scan(&inputNama)
		idx := binarySearchNama(inputNama)
		if idx != -1 {
			fmt.Printf("Ditemukan: %+v\n", data[idx])
		} else {
			fmt.Println("Tidak ditemukan.")
		}
	case 3:
		var asc string
		fmt.Print("Urutkan naik? (y/n): ")
		fmt.Scan(&asc)
		selectionSortPrioritas(asc == "y")
		tampilkanData()
	case 4:
		var asc string
		fmt.Print("Urutkan naik? (y/n): ")
		fmt.Scan(&asc)
		insertionSortTanggal(asc == "y")
		tampilkanData()
	}
}

func editProyek(id string) {
	idx := sequentialSearchID(id)
	if idx != -1 {
		fmt.Println("Masukkan data baru:")
		fmt.Print("Nama Proyek: ")
		fmt.Scan(&data[idx].Nama)
		fmt.Print("Klien: ")
		fmt.Scan(&data[idx].Klien)
		fmt.Print("Prioritas: ")
		fmt.Scan(&data[idx].Prioritas)
		fmt.Print("Tanggal: ")
		fmt.Scan(&data[idx].Tanggal)
		fmt.Print("Status: ")
		fmt.Scan(&data[idx].Status)
		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("ID tidak ditemukan.")
	}
}

func hapusProyek(id string) {
	idx := sequentialSearchID(id)
	if idx != -1 {
		for i := idx; i < n-1; i++ {
			data[i] = data[i+1]
		}
		n--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("ID tidak ditemukan.")
	}
}

// Search
func sequentialSearchID(id string) int {
	for i := 0; i < n; i++ {
		if data[i].ID == id {
			return i
		}
	}
	return -1
}

func binarySearchNama(nama string) int {
	selectionSortNama(true) // pastikan urut
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if data[mid].Nama == nama {
			return mid
		} else if data[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Sort
func selectionSortPrioritas(asc bool) {
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (asc && data[j].Prioritas < data[idx].Prioritas) || (!asc && data[j].Prioritas > data[idx].Prioritas) {
				idx = j
			}
		}
		swap(&data[i], &data[idx])
	}
}

func insertionSortTanggal(asc bool) {
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && ((asc && data[j].Tanggal > key.Tanggal) || (!asc && data[j].Tanggal < key.Tanggal)) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func selectionSortNama(asc bool) {
	
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (asc && data[j].Nama < data[idx].Nama) || (!asc && data[j].Nama > data[idx].Nama) {
				idx = j
			}
		}
		swap(&data[i], &data[idx])
	}
}

func menu() {
	var pilihan int
	var inputID string
	for pilihan != 9 {
		fmt.Println("----------------------------------------------")
		fmt.Println("\n============= Aplikasi Tugasin =============")
		fmt.Println("\n--------------------------------------------")
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah proyek")
		fmt.Println("2. Tampilkan semua proyek")
		fmt.Println("3. Edit proyek")
		fmt.Println("4. Hapus proyek")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var p Proyek
			fmt.Print("ID: ")
			fmt.Scan(&p.ID)
			fmt.Print("Nama Proyek: ")
			fmt.Scan(&p.Nama)
			fmt.Print("Klien: ")
			fmt.Scan(&p.Klien)
			fmt.Print("Prioritas: ")
			fmt.Scan(&p.Prioritas)
			fmt.Print("Tanggal (YYYY-MM-DD): ")
			fmt.Scan(&p.Tanggal)
			fmt.Print("Status (sudah/belum): ")
			fmt.Scan(&p.Status)
			tambahProyek(p)
		case 2:
			tampilkanData()
		case 3:
			fmt.Print("Masukkan ID untuk diedit: ")
			fmt.Scan(&inputID)
			editProyek(inputID)
		case 4:
			fmt.Print("Masukkan ID untuk dihapus: ")
			fmt.Scan(&inputID)
			hapusProyek(inputID)
		case 5:
			fmt.Println("Thank You")
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func main() {
	menu()
}

