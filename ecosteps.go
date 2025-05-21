package main

import "fmt"

type Aktivitas struct {
	tanggal    string
	nama       string
	skorDampak int
}
type Kategori struct {
	nama       string
	skorDampak int
}

const MaksData = 100

var daftarKategori = [10]Kategori{
	{"Membuang sampah pada tempatnya", 6},
	{"Membawa tumbler", 7},
	{"Menyumbang barang layak pakai", 8},
	{"Naik sepeda / angkutan umum", 9},
	{"Mendaur ulang sampah", 10},
	{"Menanam hidroponik", 11},
	{"Mengikuti kegiatan bersih-bersih lingkungan", 12},
	{"Menghemat listrik atau air", 13},
	{"Menggunakan energi terbarukan", 14},
	{"Menanam pohon", 15},
}

func main() {
	var pilihan int
	var selesai bool = false
	for !selesai {
		tampilanAwal()
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			menuUtama()
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
	fmt.Println("\n\n==================================================")
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
	var daftarAktivitas [MaksData]Aktivitas
	var jumlahData, pilihan int
	var selesai, dataSudahAda bool
	var sort, cari string
	for !selesai {
		fmt.Println("\n\n=========== MENU UTAMA ECOSTEPS ===========")
		fmt.Println("1. Tambah Aktivitas")
		fmt.Println("2. Tampilkan Semua Aktivitas")
		fmt.Println("3. Edit Aktivitas Berdasarkan Tanggal")
		fmt.Println("4. Hapus Aktivitas Berdasarkan Tanggal")
		fmt.Println("5. Cari Aktivitas")
		fmt.Println("6. daftarkan Aktivitas")
		fmt.Println("7. Laporan Bulanan")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahAktivitas(&daftarAktivitas, &jumlahData, &daftarKategori)
			dataSudahAda = true
		} else if pilihan >= 2 && pilihan <= 7 {
			if !dataSudahAda {
				fmt.Println("⚠️  Anda harus menambahkan aktivitas terlebih dahulu.")
			} else if dataSudahAda {
				if pilihan == 2 {
					tampilkanAktivitas(daftarAktivitas, jumlahData)
				} else if pilihan == 3 {
					editAktivitas(&daftarAktivitas, jumlahData, &daftarKategori)
				} else if pilihan == 4 {
					hapusAktivitas(&daftarAktivitas, &jumlahData)
				} else if pilihan == 5 {
					fmt.Println("\n\nPilih cari aktivitas berdasarkan tanggal(t) atau skor dampak(d).")
					fmt.Print("Pilih (t/d) : ")
					fmt.Scan(&cari)
					for cari != "t" && cari != "d" {
						fmt.Println("Input tidak valid. Coba input kembali.")
						fmt.Print("Pilih (t/d) : ")
						fmt.Scan(&cari)
					}
					if cari == "t" {
						cariAktivitas(daftarAktivitas, jumlahData)
					} else if cari == "d" {
						binarySearchPoin(daftarAktivitas, jumlahData)
					}
				} else if pilihan == 6 {
					fmt.Println("\n\nPilih daftarkan berdasarkan tanggal(t), frekuensi(f), atau skor dampak(d).")
					fmt.Print("Pilih (t/f/d) : ")
					fmt.Scan(&sort)
					for sort != "t" && sort != "f" && sort != "d" {
						fmt.Println("Input tidak valid. Coba input kembali.")
						fmt.Print("Pilih (t/f/d) : ")
						fmt.Scan(&sort)
					}
					if sort == "t" {
						sortTanggal(&daftarAktivitas, jumlahData)
					} else if sort == "f" {
						sortFrekuensi(&daftarAktivitas, jumlahData)
					} else if sort == "d" {
						sortSkor(&daftarAktivitas, jumlahData)
					}
				} else if pilihan == 7 {
					laporanBulanan(daftarAktivitas, jumlahData)
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}
		} else if pilihan == 8 {
			selesai = true
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tentangEcoSteps() {
	fmt.Println("\n\n========== Tentang EcoSteps ==========")
	fmt.Println("EcoSteps adalah aplikasi pelacak gaya hidup ramah lingkungan.")
	fmt.Println("Dengan aplikasi ini, Anda dapat memantau aktivitas sehari-hari,")
	fmt.Println("menghitung skor keberlanjutan, dan melihat laporan bulanan 🌿")
}
func aktivitas(kategori [10]Kategori) {
	var i int
	for i = 0; i < len(kategori); i++ {
		fmt.Printf("%d. %s (Skor: %d)\n", i+1, kategori[i].nama, kategori[i].skorDampak)
	}
}

func tambahAktivitas(daftar *[MaksData]Aktivitas, jumlah *int, kategori *[10]Kategori) {
	var lanjutTanggal string = "y"
	var tgl string
	for lanjutTanggal == "y" && *jumlah < MaksData {
		fmt.Println("\n--- Input Tanggal Aktivitas ---")
		fmt.Print("Tanggal (YYYY-MM-DD): ")
		fmt.Scan(&tgl)

		var lanjutAktivitas string = "y"
		for lanjutAktivitas == "y" && *jumlah < MaksData {
			fmt.Println("\nPilih aktivitas dari daftar berikut:")
			aktivitas(*kategori)
			var pilihan int
			fmt.Println("Pilih hanya SATU aktivitas!")
			fmt.Print("Masukkan nomor aktivitas: ")
			fmt.Scan(&pilihan)
			if pilihan >= 1 && pilihan <= len(*kategori) {
				(*daftar)[*jumlah].tanggal = tgl
				(*daftar)[*jumlah].nama = (*kategori)[pilihan-1].nama
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
		fmt.Print("Input aktivitas untuk tanggal lain? (y/n): ")
		fmt.Scan(&lanjutTanggal)
	}
}

func tampilkanAktivitas(daftar [MaksData]Aktivitas, jumlah int) {
	var i int
	fmt.Println("\n=== Daftar Semua Aktivitas ===")
	for i = 0; i < jumlah; i++ {
		fmt.Printf("%d. Tanggal: %s | Aktivitas: %s | Skor: %d\n",
			i+1, daftar[i].tanggal, daftar[i].nama, daftar[i].skorDampak)
	}
}
func tampilkanTanggalUnik(daftar [MaksData]Aktivitas, jumlah int) {
	var tanggalUnik [MaksData]string
	var jumlahTanggalUnik int = 0
	var i, j int
	var sudahAda bool

	fmt.Println("\n📅 Tanggal-tanggal yang sudah ada aktivitas:")

	for i = 0; i < jumlah; i++ {
		sudahAda = false
		for j = 0; !sudahAda && j < jumlahTanggalUnik; j++ {
			if daftar[i].tanggal == tanggalUnik[j] {
				sudahAda = true
			}
		}
		if !sudahAda {
			tanggalUnik[jumlahTanggalUnik] = daftar[i].tanggal
			jumlahTanggalUnik++
		}
	}
	for i = 0; i < jumlahTanggalUnik; i++ {
		fmt.Printf("- %s\n", tanggalUnik[i])
	}
}

func editAktivitas(daftar *[MaksData]Aktivitas, jumlah int, kategori *[10]Kategori) {
	var tanggal string
	tampilkanTanggalUnik(*daftar, jumlah)
	fmt.Print("Masukkan tanggal aktivitas yang ingin diedit (format: YYYY-MM-DD): ")
	fmt.Scan(&tanggal)
	var i, idx int
	var indeksAktivitas [MaksData]int
	var totalDitemukan int = 0

	for i = 0; i < jumlah; i++ {
		if (*daftar)[i].tanggal == tanggal {
			indeksAktivitas[totalDitemukan] = i
			totalDitemukan++
		}
	}

	if totalDitemukan == 0 {
		fmt.Println("⚠️ Tidak ada aktivitas ditemukan pada tanggal tersebut.")
	} else {
		fmt.Println("Daftar aktivitas pada tanggal tersebut:")
		for i = 0; i < totalDitemukan; i++ {
			idx = indeksAktivitas[i]
			fmt.Printf("%d. %s (Skor: %d)\n", i+1, (*daftar)[idx].nama, (*daftar)[idx].skorDampak)
		}

		var pilihanAktivitas int
		fmt.Print("Pilih nomor aktivitas yang ingin diedit: ")
		fmt.Scan(&pilihanAktivitas)

		if pilihanAktivitas >= 1 && pilihanAktivitas <= totalDitemukan {
			fmt.Println("Pilih aktivitas baru:")
			aktivitas(*kategori)

			var pilihanKategori int
			fmt.Print("Masukkan pilihan kategori baru: ")
			fmt.Scan(&pilihanKategori)

			if pilihanKategori >= 1 && pilihanKategori <= len(*kategori) {
				var idxEdit int = indeksAktivitas[pilihanAktivitas-1]
				(*daftar)[idxEdit].nama = (*kategori)[pilihanKategori-1].nama
				(*daftar)[idxEdit].skorDampak = (*kategori)[pilihanKategori-1].skorDampak
				fmt.Println("✅ Aktivitas berhasil diedit.")
			} else {
				fmt.Println("❌ Pilihan kategori tidak valid.")
			}
		} else {
			fmt.Println("❌ Pilihan aktivitas tidak valid.")
		}
	}
}

func hapusAktivitas(daftar *[MaksData]Aktivitas, jumlah *int) {
	var tanggal string
	fmt.Print("Masukkan tanggal aktivitas yang ingin dihapus (format: YYYY-MM-DD): ")
	fmt.Scan(&tanggal)
	var indeksAktivitas [MaksData]int
	var totalDitemukan int = 0
	var i, idx int
	for i = 0; i < *jumlah; i++ {
		if (*daftar)[i].tanggal == tanggal {
			indeksAktivitas[totalDitemukan] = i
			totalDitemukan++
		}
	}
	if totalDitemukan == 0 {
		fmt.Println("⚠️ Tidak ada aktivitas ditemukan pada tanggal tersebut.")
	} else {
		fmt.Println("Daftar aktivitas pada tanggal tersebut:")
		for i = 0; i < totalDitemukan; i++ {
			idx = indeksAktivitas[i]
			fmt.Printf("%d. %s (Skor: %d)\n", i+1, (*daftar)[idx].nama, (*daftar)[idx].skorDampak)
		}
		var pilihanAktivitas int
		fmt.Print("Pilih nomor aktivitas yang ingin dihapus: ")
		fmt.Scan(&pilihanAktivitas)

		if pilihanAktivitas >= 1 && pilihanAktivitas <= totalDitemukan {
			var idxHapus int = indeksAktivitas[pilihanAktivitas-1]
			for i = idxHapus; i < *jumlah-1; i++ {
				(*daftar)[i] = (*daftar)[i+1]
			}
			*jumlah--
			fmt.Println("✅ Aktivitas berhasil dihapus.")
		} else {
			fmt.Println("❌ Pilihan aktivitas tidak valid.")
		}
	}
}

func cariAktivitas(daftar [MaksData]Aktivitas, jumlah int) {
	var keyword string
	var ditemukan bool = false
	var i int
	fmt.Println("\n=== Cari Aktivitas ===")
	fmt.Print("Masukkan kata kunci (tanggal aktivitas): ")
	fmt.Scanln(&keyword)
	fmt.Scanln(&keyword)

	for i = 0; i < jumlah; i++ {
		fmt.Println("\nHasil pencarian:")
		if daftar[i].tanggal == keyword {
			ditemukan = true
			fmt.Printf("- %s | %s | Skor: %d\n",
				daftar[i].tanggal, daftar[i].nama, daftar[i].skorDampak)
		}
	}

	if ditemukan == false {
		fmt.Println("⚠️ Tidak ditemukan aktivitas yang sesuai.")
	}
}
func sortTanggal(daftar *[MaksData]Aktivitas, jumlah int) {
	var i, j, idx int
	var temp Aktivitas
	var urut string
	fmt.Println("Pilih daftarkan tanggal secara decending(d) atau ascending(a)")
	fmt.Print("Pilih salah satu (d/a):")
	fmt.Scan(&urut)

	if urut == "a" {
		for i = 0; i < jumlah-1; i++ {
			idx = i
			for j = i + 1; j < jumlah; j++ {
				if (*daftar)[j].tanggal < (*daftar)[idx].tanggal {
					idx = j
				}
			}
			if idx != i {
				temp = (*daftar)[i]
				(*daftar)[i] = (*daftar)[idx]
				(*daftar)[idx] = temp
			}
		}
	} else if urut == "d" {
		for i = 0; i < jumlah-1; i++ {
			idx = i
			for j = i + 1; j < jumlah; j++ {
				if (*daftar)[j].tanggal > (*daftar)[idx].tanggal {
					idx = j
				}
			}
			if idx != i {
				temp = (*daftar)[i]
				(*daftar)[i] = (*daftar)[idx]
				(*daftar)[idx] = temp
			}
		}
	}
	tampilkanAktivitas(*daftar, jumlah)
	fmt.Println("✅ Aktivitas berhasil didaftarkan berdasarkan tanggal.")
}

func binarySearchPoin(daftar [MaksData]Aktivitas, jumlah int) {

	var temp Aktivitas
	var x, i, j, idx int
	fmt.Print("Cari aktivitas berdasarkan skor dampak (6–15): ")
	fmt.Scan(&x)
	for i = 0; i < jumlah-1; i++ {
		idx = i
		for j = i + 1; j < jumlah; j++ {
			if daftar[j].skorDampak < daftar[idx].skorDampak {
				idx = j
			}
		}
		temp = daftar[i]
		daftar[i] = daftar[idx]
		daftar[idx] = temp
	}

	var left, right, mid, found int
	left = 0
	right = jumlah - 1
	found = -1

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if x < daftar[mid].skorDampak {
			right = mid - 1
		} else if x > daftar[mid].skorDampak {
			left = mid + 1
		} else {
			found = mid
		}
	}

	if found == -1 {
		fmt.Println("⚠️ Skor dampak tidak ditemukan.")
	} else {
		fmt.Println("✅ Ditemukan aktivitas dengan skor dampak", x, ":")

		fmt.Printf("- %s | %s | Skor: %d\n",
			daftar[found].tanggal, daftar[found].nama, daftar[found].skorDampak)

		for i = found - 1; i >= 0 && daftar[i].skorDampak == x; i-- {
			fmt.Printf("- %s | %s | Skor: %d\n",
				daftar[i].tanggal, daftar[i].nama, daftar[i].skorDampak)
		}
		for i = found + 1; i < jumlah && daftar[i].skorDampak == x; i++ {
			fmt.Printf("- %s | %s | Skor: %d\n",
				daftar[i].tanggal, daftar[i].nama, daftar[i].skorDampak)
		}
	}
}

func hitungFrekuensi(daftar [MaksData]Aktivitas, jumlah int, nama string) int {
	var count int = 0
	var i int
	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			count++
		}
	}
	return count
}
func sortFrekuensi(daftar *[MaksData]Aktivitas, jumlah int) {
	var i, j int
	var temp Aktivitas
	var freqTemp, freqJ int
	var geser bool
	var urut string
	fmt.Println("Pilih daftarkan frekuensi secara decending(d) atau ascending(a)")
	fmt.Print("Pilih salah satu (d/a):")
	fmt.Scan(&urut)
	if urut == "d" {
		for i = 1; i < jumlah; i++ {
			temp = (*daftar)[i]
			freqTemp = hitungFrekuensi(*daftar, jumlah, temp.nama)
			j = i - 1
			geser = true

			for geser && j >= 0 {
				freqJ = hitungFrekuensi(*daftar, jumlah, (*daftar)[j].nama)
				if freqJ < freqTemp {
					(*daftar)[j+1] = (*daftar)[j]
					j = j - 1
				} else {
					geser = false
				}
			}
			(*daftar)[j+1] = temp
		}
	} else if urut == "a" {
		for i = 1; i < jumlah; i++ {
			temp = (*daftar)[i]
			freqTemp = hitungFrekuensi(*daftar, jumlah, temp.nama)
			j = i - 1
			geser = true

			for geser && j >= 0 {
				freqJ = hitungFrekuensi(*daftar, jumlah, (*daftar)[j].nama)
				if freqJ > freqTemp {
					(*daftar)[j+1] = (*daftar)[j]
					j = j - 1
				} else {
					geser = false
				}
			}
			(*daftar)[j+1] = temp
		}
	}
	tampilkanAktivitas(*daftar, jumlah)
	fmt.Println("✅ Aktivitas berhasil didaftarkan berdasarkan frekuensi aktivitas.")

}
func sortSkor(daftar *[MaksData]Aktivitas, jumlah int) {
	var i, j, idx int
	var temp Aktivitas
	var urut string
	fmt.Println("Pilih daftarkan frekuensi secara decending(d) atau ascending(a)")
	fmt.Print("Pilih salah satu (d/a):")
	fmt.Scan(&urut)
	if urut == "a" {
		for i = 0; i < jumlah-1; i++ {
			idx = i
			for j = i + 1; j < jumlah; j++ {
				if (*daftar)[j].skorDampak < (*daftar)[idx].skorDampak {
					idx = j
				}
			}
			temp = (*daftar)[i]
			(*daftar)[i] = (*daftar)[idx]
			(*daftar)[idx] = temp
		}
	} else if urut == "d" {
		for i = 0; i < jumlah-1; i++ {
			idx = i
			for j = i + 1; j < jumlah; j++ {
				if (*daftar)[j].skorDampak > (*daftar)[idx].skorDampak {
					idx = j
				}
			}
			temp = (*daftar)[i]
			(*daftar)[i] = (*daftar)[idx]
			(*daftar)[idx] = temp
		}
	}
	tampilkanAktivitas(*daftar, jumlah)
	fmt.Println("✅ Aktivitas berhasil didaftarkan berdasarkan skor dampak aktivitas.")
}
func laporanBulanan(daftar [MaksData]Aktivitas, jumlah int) {
	var i int
	var bulanSekarang string = daftar[i].tanggal[:7]
	var bulanSebelumnya string
	var totalSkor, totalAktivitas int
	if jumlah != 0 {
		fmt.Println("\n===== LAPORAN BULANAN =====")
		for i = 0; i < jumlah; i++ {
			if bulanSekarang != bulanSebelumnya && i != 0 {
				fmt.Printf("📅 Bulan: %s - Total Aktivitas: %d - Total Skor: %d\n",
					bulanSebelumnya, totalAktivitas, totalSkor)
				totalSkor = 0
				totalAktivitas = 0
			}

			totalSkor += daftar[i].skorDampak
			totalAktivitas++
			bulanSebelumnya = bulanSekarang
		}

		fmt.Printf("📅 Bulan: %s - Total Aktivitas: %d - Total Skor: %d\n",
			bulanSebelumnya, totalAktivitas, totalSkor)
	} else {
		fmt.Println("⚠️ Belum ada aktivitas untuk dihitung.")
	}
}
