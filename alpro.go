//HALO INI HASIL KOMENTAR

package main

import "fmt"

const NMAX int = 100

type Time struct {
	menit int //menyimpan jumlah menit
	detik int
}

type Content struct {
	judul  string
	ukuran float64
	durasi Time
	sfx    string
	likes  int
}

type Creator struct {
	nama         string
	umur         int
	jumlahKonten int
	konten       [NMAX]Content
}

type DaftarCreator [NMAX]Creator

var semuaCreator DaftarCreator
var jumlahCreator int = 0

func main() {
	isiDataAwal(&semuaCreator, &jumlahCreator)
	tampilkanMenuUtama()
}

func tampilkanMenuUtama() {
	var pilihan int
	fmt.Println("Selamat datang di Platform Konten Kreator!")
	fmt.Println("Siapa Anda?")
	fmt.Println("1. Kreator")
	fmt.Println("2. Penonton")
	fmt.Println("0. Keluar")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		menuKreator()
	} else if pilihan == 2 {
		menuPenonton()
	} else if pilihan == 0 {
		fmt.Println("Terima kasih!")
		return
	} else {
		fmt.Println("Pilihan tidak valid")
		tampilkanMenuUtama()
	}
}

func menuKreator() {
	var pilihan int
	fmt.Println("Menu Kreator")
	fmt.Println("1. Input Data Kreator dan Konten")
	fmt.Println("0. Kembali")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		inputDataKreator(&semuaCreator, &jumlahCreator)
		tampilkanMenuUtama()
	} else if pilihan == 0 {
		tampilkanMenuUtama()
	} else {
		fmt.Println("Pilihan tidak valid")
		menuKreator()
	}
}

func menuPenonton() {
	var pilihan int
	fmt.Println("Menu Penonton")
	fmt.Println("1. Tampilkan Semua Konten")
	fmt.Println("2. Cari Konten Berdasarkan Judul")
	fmt.Println("3. Ukuran Terbesar dan Terkecil")
	fmt.Println("4. Urutkan Berdasarkan Likes")
	fmt.Println("5. Tonton Konten")
	fmt.Println("0. Kembali")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		tampilkanSemuaKonten(semuaCreator, jumlahCreator)
		menuPenonton()
	} else if pilihan == 2 {
		cariKontenDariSemuaCreator(semuaCreator, jumlahCreator)
		menuPenonton()
	} else if pilihan == 3 {
		temukanUkuranMinMax(semuaCreator, jumlahCreator)
		menuPenonton()
	} else if pilihan == 4 {
		urutkanKontenBerdasarkanLikes(&semuaCreator, jumlahCreator)
		menuPenonton()
	} else if pilihan == 5 {
		tontonKonten(semuaCreator, jumlahCreator)
		menuPenonton()
	} else if pilihan == 0 {
		tampilkanMenuUtama()
	} else {
		fmt.Println("Pilihan tidak valid")
		menuPenonton()
	}
}

//input
func isiDataAwal(A *DaftarCreator, jumlah *int) {
	A[0].nama = "Alice"
	A[0].umur = 30
	A[0].jumlahKonten = 2
	A[0].konten[0] = Content{"Vlog Bali", 120.5, Time{10, 30}, "Ocean Waves", 150}
	A[0].konten[1] = Content{"Tutorial Masak", 98.0, Time{15, 45}, "Kitchen Sound", 200}

	A[1].nama = "Bob"
	A[1].umur = 25
	A[1].jumlahKonten = 2
	A[1].konten[0] = Content{"Review Game", 150.75, Time{20, 0}, "Game Music", 300}
	A[1].konten[1] = Content{"Unboxing", 85.25, Time{8, 15}, "Box Sound", 180}

	A[2].nama = "System"
	A[2].umur = 0
	A[2].jumlahKonten = 2
	A[2].konten[0] = Content{"Travel Documentary", 132.0, Time{12, 50}, "Nature Ambience", 120}
	A[2].konten[1] = Content{"Science Explainer", 111.7, Time{14, 10}, "Lab Sound", 210}

	*jumlah = 3
}

func inputDataKreator(A *DaftarCreator, jumlah *int) int {
	var n, j int
	fmt.Print("Masukkan jumlah kreator yang ingin diinput: ")
	fmt.Scan(&n)

	for i := *jumlah; i < *jumlah+n; i++ {
		fmt.Print("Nama Kreator: ")
		fmt.Scan(&A[i].nama)
		fmt.Print("Umur Kreator: ")
		fmt.Scan(&A[i].umur)
		fmt.Print("Jumlah konten yang ingin dimasukkan: ")
		fmt.Scan(&A[i].jumlahKonten)

		for j = 0; j < A[i].jumlahKonten; j++ {
			fmt.Print("Judul Konten: ")
			fmt.Scan(&A[i].konten[j].judul)
			fmt.Print("Ukuran (MB): ")
			fmt.Scan(&A[i].konten[j].ukuran)
			fmt.Print("Durasi Menit: ")
			fmt.Scan(&A[i].konten[j].durasi.menit)
			fmt.Print("Durasi Detik: ")
			fmt.Scan(&A[i].konten[j].durasi.detik)
			fmt.Print("SFX: ")
			fmt.Scan(&A[i].konten[j].sfx)
			fmt.Print("Jumlah Likes: ")
			fmt.Scan(&A[i].konten[j].likes)
			fmt.Println("----------------------")
		}
		fmt.Println("==========================")
	}
	*jumlah += n
	return *jumlah
}

func tampilkanSemuaKonten(A DaftarCreator, jumlah int) {
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Kreator: %s\n", A[i].nama)
		for j := 0; j < A[i].jumlahKonten; j++ {
			k := A[i].konten[j]
			fmt.Printf("  Judul: %s | Ukuran: %.2f MB | Durasi: %dm %ds | SFX: %s | Likes: %d\n",
				k.judul, k.ukuran, k.durasi.menit, k.durasi.detik, k.sfx, k.likes)
		}
	}
}

func cariKontenDariSemuaCreator(A DaftarCreator, jumlah int) {
	var judul1, judul2 string
	fmt.Print("Masukkan judul konten yang dicari: ")
	fmt.Scan(&judul1, &judul2)
	judul := judul1 + " " + judul2

	ketemu := false
	for i := 0; i < jumlah; i++ {
		for j := 0; j < A[i].jumlahKonten; j++ {
			if A[i].konten[j].judul == judul {
				fmt.Printf("Ditemukan! Kreator: %s, Likes: %d\n", A[i].nama, A[i].konten[j].likes)
				ketemu = true
			}
		}
	}
	if !ketemu {
		fmt.Println("Konten tidak ditemukan.")
	}
}

func tontonKonten(A DaftarCreator, jumlah int) {
	var judul1, judul2 string
	fmt.Print("Masukkan judul konten yang ingin ditonton: ")
	fmt.Scan(&judul1, &judul2)
	judulDicari := judul1 + " " + judul2

	var semuaKonten [NMAX]Content
	var semuaKreator [NMAX]string
	var nKonten int = 0

	// Gabungkan semua konten ke dalam satu array dan simpan kreatornya
	for i := 0; i < jumlah; i++ {
		for j := 0; j < A[i].jumlahKonten; j++ {
			semuaKonten[nKonten] = A[i].konten[j]
			semuaKreator[nKonten] = A[i].nama
			nKonten++
		}
	}

	// Insertion Sort berdasarkan judul konten (ascending)
	for i := 1; i < nKonten; i++ {
		tempKonten := semuaKonten[i]
		tempKreator := semuaKreator[i]
		j := i - 1
		for j >= 0 && semuaKonten[j].judul > tempKonten.judul {
			semuaKonten[j+1] = semuaKonten[j]
			semuaKreator[j+1] = semuaKreator[j]
			j--
		}
		semuaKonten[j+1] = tempKonten
		semuaKreator[j+1] = tempKreator
	}

	// Binary Search untuk mencari judul konten
	low := 0
	high := nKonten - 1
	ketemu := false

	for low <= high {
		mid := (low + high) / 2
		if semuaKonten[mid].judul == judulDicari {
			fmt.Println("Menonton Konten:")
			fmt.Printf("Judul   : %s\n", semuaKonten[mid].judul)
			fmt.Printf("Kreator : %s\n", semuaKreator[mid])
			fmt.Printf("Ukuran  : %.2f MB\n", semuaKonten[mid].ukuran)
			fmt.Printf("Durasi  : %d menit %d detik\n", semuaKonten[mid].durasi.menit, semuaKonten[mid].durasi.detik)
			fmt.Printf("SFX     : %s\n", semuaKonten[mid].sfx)
			fmt.Printf("Likes   : %d\n", semuaKonten[mid].likes)
			ketemu = true
			break
		} else if semuaKonten[mid].judul < judulDicari {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !ketemu {
		fmt.Println("Konten tidak ditemukan.")
	}
}

func temukanUkuranMinMax(A DaftarCreator, jumlah int) {
	cariUkuranTerkecil(A, jumlah)
	cariUkuranTerbesar(A, jumlah)
}

func cariUkuranTerkecil(A DaftarCreator, jumlah int) {
	var min float64 = 999999
	var judul string
	for i := 0; i < jumlah; i++ {
		for j := 0; j < A[i].jumlahKonten; j++ {
			if A[i].konten[j].ukuran < min {
				min = A[i].konten[j].ukuran
				judul = A[i].konten[j].judul
			}
		}
	}
	fmt.Printf("Ukuran terkecil: %.2f MB (%s)\n", min, judul)
}

func cariUkuranTerbesar(A DaftarCreator, jumlah int) {
	var max float64 = -1
	var judul string
	for i := 0; i < jumlah; i++ {
		for j := 0; j < A[i].jumlahKonten; j++ {
			if A[i].konten[j].ukuran > max {
				max = A[i].konten[j].ukuran
				judul = A[i].konten[j].judul
			}
		}
	}
	fmt.Printf("Ukuran terbesar: %.2f MB (%s)\n", max, judul)
}

func urutkanKontenBerdasarkanLikes(A *DaftarCreator, jumlah int) {
	var semuaKonten [NMAX]Content
	var nKonten int = 0

	for i := 0; i < jumlah; i++ {
		for j := 0; j < A[i].jumlahKonten; j++ {
			semuaKonten[nKonten] = A[i].konten[j]
			nKonten++
		}
	}

	for i := 0; i < nKonten-1; i++ {
		maxIdx := i
		for j := i + 1; j < nKonten; j++ {
			if semuaKonten[j].likes > semuaKonten[maxIdx].likes {
				maxIdx = j
			}
		}
		semuaKonten[i], semuaKonten[maxIdx] = semuaKonten[maxIdx], semuaKonten[i]
	}

	fmt.Println("Konten diurutkan berdasarkan Likes (menurun):")
	for i := 0; i < nKonten; i++ {
		k := semuaKonten[i]
		fmt.Printf("Judul: %s | Ukuran: %.2f MB | Durasi: %dm %ds | SFX: %s | Likes: %d\n",
			k.judul, k.ukuran, k.durasi.menit, k.durasi.detik, k.sfx, k.likes)
	}
}
