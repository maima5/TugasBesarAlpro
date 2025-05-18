package main

import "fmt"

type Aktivitas struct {
	tanggal    string
	nama       string
	kategori   string
	skorDampak int
}

type Kategori struct {
	nama       string
	skorDampak int
}

const MaksData = 100

var daftarAktivitas [MaksData]Aktivitas
var jumlahData int = 0

var daftarKategori = [10]Kategori{
	{"Naik sepeda / angkutan umum", 15},
	{"Menghemat listrik atau air", 20},
	{"Membuang sampah pada tempatnya", 5},
	{"Membawa tumbler", 8},
	{"Mengikuti kegiatan bersih-bersih lingkungan", 18},
	{"Mendaur ulang sampah", 15},
	{"Menyumbang barang layak pakai", 12},
	{"Menggunakan energi terbarukan", 25},
	{"Menanam hidroponik", 18},
	{"Menanam pohon", 25},
}

func main() {
	var pilihan int
	selesai := false

	for !selesai {
		tampilanAwal()
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			menuUtama()
			pilihan = 0
		} else if pilihan == 2 {
			tentangEcoSteps()
		} else if pilihan == 3 {
			fmt.Printf("\nTerima kasih telah menggunakan EcoSteps 🌿")
			selesai = true
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
func tampilanAwal() {
	fmt.Println("==================================================")
	fmt.Println("              🌱  SELAMAT DATANG  🌱             ")
	fmt.Println("              Ｅ ｃ ｏ Ｓ ｔ ｅ ｐ ｓ               ")
	fmt.Println("         Pelacak Gaya Hidup Ramah Lingkungan       ")
	fmt.Println("---------------------------------------------------")
	fmt.Println("          Algoritma & Pemrograman 2 - 2025         ")
	fmt.Println("   Dibuat oleh: Husnul Khotimah & Nailah Dhiyah")
	fmt.Println("==================================================")
	fmt.Println("1. Mulai Aplikasi")
	fmt.Println("2. Tentang EcoSteps")
	fmt.Println("3. Keluar")
	fmt.Println("---------------------------------------------------")
	fmt.Print("Pilihan anda: ")
}
func menuUtama() {
	var pilihan int
	var selesai bool = false

	for !selesai {
		fmt.Println("\n=========== MENU UTAMA ECOSTEPS ===========")
		fmt.Println("1. Tambah Aktivitas")
		fmt.Println("2. Tampilkan Semua Aktivitas")
		fmt.Println("3. Cari Aktivitas")
		fmt.Println("4. Urutkan Aktivitas berdasarkan tanggal")
		fmt.Println("5. Laporan Bulanan")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahAktivitas(&daftarAktivitas, &jumlahData, &daftarKategori)
			pilihan = 0
		} else if pilihan == 2 {
			tampilkanAktivitas(&daftarAktivitas, &jumlahData)
		} else if pilihan == 3 {
			cariAktivitas(&daftarAktivitas, &jumlahData)
		} else if pilihan == 4 {
			urutkanAktivitasTanggal(&daftarAktivitas, jumlahData)
		} else if pilihan == 5 {
			laporanBulanan(&daftarAktivitas, jumlahData)
		} else if pilihan == 6 {
			selesai = true
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
func pilihMenu(pilihan int) {
	switch pilihan {
	case 1:
		tambahAktivitas(&daftarAktivitas, &jumlahData, &daftarKategori)
	case 2:
		tampilkanAktivitas(&daftarAktivitas, &jumlahData)
	case 3:
		cariAktivitas(&daftarAktivitas, &jumlahData)
	case 4:
		urutkanAktivitasTanggal(&daftarAktivitas, jumlahData)
		tampilkanAktivitas(&daftarAktivitas, &jumlahData)
	case 5:
		laporanBulanan(&daftarAktivitas, jumlahData)
	case 6:
		tampilanAwal()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
func tentangEcoSteps() {
	fmt.Println("\n========== Tentang EcoSteps ==========")
	fmt.Println("EcoSteps adalah aplikasi pelacak gaya hidup ramah lingkungan.")
	fmt.Println("Dengan aplikasi ini, Anda dapat memantau aktivitas sehari-hari,")
	fmt.Println("menghitung skor keberlanjutan, dan melihat laporan bulanan 🌿")
}

func tambahAktivitas(daftar *[MaksData]Aktivitas, jumlah *int, kategori *[10]Kategori) {
	var lanjutTanggal string = "y"

	for lanjutTanggal == "y" && *jumlah < MaksData {
		var tgl string
		fmt.Println("\n--- Input Tanggal Aktivitas ---")
		fmt.Print("Tanggal (YYYY-MM-DD): ")
		fmt.Scan(&tgl)

		var lanjutAktivitas string = "y"
		for lanjutAktivitas == "y" && *jumlah < MaksData {
			fmt.Println("\nPilih aktivitas dari daftar berikut:")
			for i := 0; i < len(*kategori); i++ {
				fmt.Printf("%d. %s (Skor: %d)\n", i+1, (*kategori)[i].nama, (*kategori)[i].skorDampak)
			}

			var pilihan int
			fmt.Print("Masukkan nomor aktivitas: ")
			fmt.Scan(&pilihan)

			if pilihan >= 1 && pilihan <= len(*kategori) {
				(*daftar)[*jumlah].tanggal = tgl
				(*daftar)[*jumlah].nama = (*kategori)[pilihan-1].nama
				(*daftar)[*jumlah].kategori = (*kategori)[pilihan-1].nama
				(*daftar)[*jumlah].skorDampak = (*kategori)[pilihan-1].skorDampak
				*jumlah = *jumlah + 1
				fmt.Println("✅ Aktivitas berhasil ditambahkan.")
			} else {
				fmt.Println("⚠️ Nomor kategori tidak valid.")
			}

			if *jumlah < MaksData {
				fmt.Print("Tambah aktivitas lain untuk tanggal ini? (y/n): ")
				fmt.Scan(&lanjutAktivitas)
			} else {
				fmt.Println("⚠️ Kapasitas data sudah penuh.")
				lanjutAktivitas = "n"
			}
		}

		if *jumlah < MaksData {
			fmt.Print("Input aktivitas untuk tanggal lain? (y/n): ")
			fmt.Scan(&lanjutTanggal)
		} else {
			fmt.Println("⚠️ Data sudah penuh, tidak bisa menambahkan tanggal baru.")
			lanjutTanggal = "n"
		}
	}
}

func tampilkanAktivitas(daftar *[MaksData]Aktivitas, jumlah *int) {
	fmt.Println("\n=== Daftar Semua Aktivitas ===")
	if *jumlah == 0 {
		fmt.Println("⚠️ Belum ada aktivitas yang tercatat.")
	} else {
		for i := 0; i < *jumlah; i++ {
			fmt.Printf("%d. Tanggal: %s | Aktivitas: %s | Skor: %d\n",
				i+1, (*daftar)[i].tanggal, (*daftar)[i].nama, (*daftar)[i].skorDampak)
		}
	}
}
func cariAktivitas(daftar *[MaksData]Aktivitas, jumlah *int) {
	var keyword string
	var ditemukan bool = false

	fmt.Println("\n=== Cari Aktivitas ===")
	fmt.Print("Masukkan kata kunci (tanggal aktivitas): ")
	fmt.Scanln(&keyword)
	fmt.Scanln(&keyword)

	for i := 0; i < *jumlah; i++ {
		if daftar[i].tanggal == keyword {
			if ditemukan == false {
				fmt.Println("\nHasil pencarian:")
			}
			ditemukan = true
			fmt.Printf("- %s | %s | Skor: %d\n",
				(*daftar)[i].tanggal, (*daftar)[i].nama, (*daftar)[i].skorDampak)
		}
	}

	if ditemukan == false {
		fmt.Println("⚠️ Tidak ditemukan aktivitas yang sesuai.")
	}
}
func urutkanAktivitasTanggal(daftar *[MaksData]Aktivitas, jumlah int) {
	var i, j, idxMin int
	var temp Aktivitas

	i = 0
	for i < jumlah-1 {
		idxMin = i
		j = i + 1
		for j < jumlah {
			if (*daftar)[j].tanggal < (*daftar)[idxMin].tanggal {
				idxMin = j
			}
			j++
		}
		if idxMin != i {
			temp = (*daftar)[i]
			(*daftar)[i] = (*daftar)[idxMin]
			(*daftar)[idxMin] = temp
		}
		i++
	}
	fmt.Println("\n=== Daftar Semua Aktivitas ===")
	if jumlah == 0 {
		fmt.Println("⚠️ Belum ada aktivitas yang tercatat.")
	} else {
		for i := 0; i < jumlah; i++ {
			fmt.Printf("%d. Tanggal: %s | Aktivitas: %s | Skor: %d\n",
				i+1, (*daftar)[i].tanggal, (*daftar)[i].nama, (*daftar)[i].skorDampak)
		}
	}
	fmt.Println("✅ Aktivitas berhasil diurutkan berdasarkan tanggal.")
}

func laporanBulanan(daftar *[MaksData]Aktivitas, jumlah int) {
	if jumlah == 0 {
		fmt.Println("⚠️ Belum ada aktivitas untuk dihitung.")
		return
	}

	var i int
	var bulanSebelumnya string
	var totalSkor, totalAktivitas int

	fmt.Println("\n===== LAPORAN BULANAN =====")
	i = 0
	for i < jumlah {
		bulanSekarang := (*daftar)[i].tanggal[:7]
		if bulanSekarang != bulanSebelumnya && i != 0 {
			fmt.Printf("📅 Bulan: %s - Total Aktivitas: %d - Total Skor: %d\n",
				bulanSebelumnya, totalAktivitas, totalSkor)
			totalSkor = 0
			totalAktivitas = 0
		}

		totalSkor += (*daftar)[i].skorDampak
		totalAktivitas++
		bulanSebelumnya = bulanSekarang
		i++
	}

	fmt.Printf("📅 Bulan: %s - Total Aktivitas: %d - Total Skor: %d\n",
		bulanSebelumnya, totalAktivitas, totalSkor)
}
