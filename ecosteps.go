package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Aktivitas struct {
	year, month, day int
	nama             string
	skorDampak       int
}
type Kategori struct {
	nama       string
	skorDampak int
}

const MaksData int = 100

type arraktivitas [MaksData]Aktivitas
type arrkat [10]Kategori

var daftarKategori = arrkat{
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

func refreshTampilan(daftar arraktivitas, jumlah int, showMessage string) {
	cls()
	tampilkanAktivitas(daftar, jumlah)
	if showMessage != "" {
		fmt.Println("\n", showMessage)
	}
}
func cls() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func main() {
	var pilihan int
	var selesai bool = false
	for !selesai {
		tampilanAwal()
		fmt.Scan(&pilihan)
		cls()
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
	cls()
	fmt.Println("\033[32m") // Warna hijau
	fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                         🌱 SELAMAT DATANG 🌱                         ║")
	fmt.Println("║                         Ｅ ｃ ｏ Ｓ ｔ ｅ ｐ ｓ                      ║")
	fmt.Println("║                Pelacak Gaya Hidup Ramah Lingkungan 🌿                ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║  Algoritma & Pemrograman 2 - 2025                                    ║")
	fmt.Println("║  Dibuat oleh: Husnul Khotimah & Nailah Dhiya                         ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║  > 1. Mulai Aplikasi                                                 ║")
	fmt.Println("║  > 2. Tentang EcoSteps                                               ║")
	fmt.Println("║  > 3. Keluar                                                         ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
	fmt.Print("🌿 Pilihan anda ➤ ")
	fmt.Println("\033[0m") // Reset warna
}
func menuUtama() {
	var daftarAktivitas arraktivitas
	var jumlahData, pilihan int
	var selesai bool
	dummyData(&daftarAktivitas, &jumlahData)
	tampilkanAktivitas(daftarAktivitas, jumlahData)
	for !selesai {
		fmt.Println("\033[32m") // Warna hijau
		fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                      📋  M E N U   U T A M A  📋                     ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║  📌 1. Tambah Aktivitas                                              ║")
		fmt.Println("║  ✏️ 2. Edit Aktivitas Berdasarkan Tanggal                            ║")
		fmt.Println("║  ❌ 3. Hapus Aktivitas Berdasarkan Tanggal                           ║")
		fmt.Println("║  🔍 4. Cari Aktivitas                                                ║")
		fmt.Println("║  🔃 5. Urutkan Aktivitas                                             ║")
		fmt.Println("║  📅 6. Laporan Bulanan                                               ║")
		fmt.Println("║  🔙 7. Kembali ke Halaman Awal                                       ║")
		fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
		fmt.Print("🌱 Pilihan anda ➤ ")
		fmt.Println("\033[0m") // Reset warna

		fmt.Scan(&pilihan)
		cls()
		if pilihan == 1 {
			tambahAktivitas(&daftarAktivitas, &jumlahData, &daftarKategori)
		} else if pilihan == 2 {
			editAktivitas(&daftarAktivitas, jumlahData, daftarKategori)
		} else if pilihan == 3 {
			hapusAktivitas(&daftarAktivitas, &jumlahData)
		} else if pilihan == 4 {
			cari(daftarAktivitas, jumlahData)
		} else if pilihan == 5 {
			sorting(&daftarAktivitas, jumlahData)
		} else if pilihan == 6 {
			laporanBulanan(daftarAktivitas, jumlahData)
		} else if pilihan == 7 {
			selesai = true
		} else {
			fmt.Println("⚠️ Pilihan tidak valid.")
		}
	}
}

func tentangEcoSteps() {
	var kembali string
	fmt.Println("\033[32m") // Warna hijau
	fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    📖 TENTANG APLIKASI ECOSTEPS                      ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║  EcoSteps adalah aplikasi pelacak gaya hidup ramah lingkungan.       ║")
	fmt.Println("║  Dengan aplikasi ini, Anda dapat memantau aktivitas sehari-hari,     ║")
	fmt.Println("║  menghitung skor keberlanjutan, dan melihat laporan bulanan. 🌿      ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")

	fmt.Print("Ketik x untuk kembali ➤ ")
	fmt.Println("\033[0m") // Reset warna
	fmt.Scan(&kembali)
	if kembali == "x" {
		cls()
	} else {
		for kembali != "x" {
			fmt.Println("\033[0m") // Reset warna
			fmt.Print("Ketik x untuk kembali ➤ ")
			fmt.Scan(&kembali)
			fmt.Println("\033[0m") // Reset warna
		}
	}
}

func listKategori(kategori arrkat) {
	var baris string
	fmt.Println("\033[32m")
	fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║ 📌 Pilih Aktivitas dari Daftar Berikut                               ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════════╣")
	for i := 0; i < len(kategori); i++ {
		baris = fmt.Sprintf("%2d. %s (Skor: %d)", i+1, kategori[i].nama, kategori[i].skorDampak)
		fmt.Printf("║ %-68s ║\n", baris)
	}
	fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
}

func tambahAktivitas(daftar *arraktivitas, jumlah *int, kategori *arrkat) {
	var lanjutTanggal string = "y"
	var inputTanggal string
	var tahun, bulan, hari int
	var selesai bool

	fmt.Println("\nKetik x untuk kembali.")
	for !selesai {
		for lanjutTanggal == "y" && *jumlah < MaksData && !selesai {
			fmt.Println("\033[32m")

			fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("║                    📅 Input Tanggal Aktivitas                        ║")
			fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
			fmt.Print("Tanggal (YYYY/MM/DD): ")
			fmt.Println("\033[0m")
			fmt.Scan(&inputTanggal)

			if inputTanggal == "x" {
				selesai = true
				refreshTampilan(*daftar, *jumlah, " ")
			} else {
				fmt.Sscanf(inputTanggal, "%d/%d/%d", &tahun, &bulan, &hari)
				var lanjutAktivitas string = "y"
				for lanjutAktivitas == "y" && *jumlah < MaksData && !selesai {
					listKategori(*kategori)
					fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
					fmt.Printf("║ %-69s ║\n", "⚠️  Pilih hanya SATU aktivitas!")
					fmt.Printf("║ %-69s ║\n", "➡️  Masukkan nomor aktivitas:")
					fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
					fmt.Print("🌿 ➤ ")
					fmt.Println("\033[0m")

					var pilihan int
					fmt.Scanf("\n%d", &pilihan)

					if pilihan >= 1 && pilihan <= len(*kategori) {
						(*daftar)[*jumlah].year = tahun
						(*daftar)[*jumlah].month = bulan
						(*daftar)[*jumlah].day = hari
						(*daftar)[*jumlah].nama = (*kategori)[pilihan-1].nama
						(*daftar)[*jumlah].skorDampak = (*kategori)[pilihan-1].skorDampak
						*jumlah++
						fmt.Println("✅ Aktivitas berhasil ditambahkan.")
					} else {
						fmt.Println("⚠️ Nomor kategori tidak valid.")
					}

					if *jumlah < MaksData {
						fmt.Println("\033[32m")
						fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
						fmt.Println("║        ➕ Tambah aktivitas lain untuk tanggal ini? (y/n/x)             ║")
						fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
						fmt.Print("🌿 Jawaban anda ➤ ")
						fmt.Println("\033[0m")
						fmt.Scanf("\n%s", &lanjutAktivitas)
						if lanjutAktivitas == "x" {
							selesai = true
							refreshTampilan(*daftar, *jumlah, "")
						}
					} else {
						fmt.Println("⚠️ Kapasitas data sudah penuh.")
						lanjutAktivitas = "n"
					}
				}

				if *jumlah < MaksData && !selesai {
					fmt.Println("\033[32m")
					fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
					fmt.Println("║        📅 Input aktivitas untuk tanggal lain? (y/n/x)                  ║")
					fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
					fmt.Print("🌿 Jawaban anda ➤ ")
					fmt.Println("\033[0m")
					fmt.Scanf("\n%s", &lanjutTanggal)
					if lanjutTanggal == "x" {
						selesai = true
						refreshTampilan(*daftar, *jumlah, "")
					}
				}
			}
		}
		selesai = true
		refreshTampilan(*daftar, *jumlah, "")
	}

}

func tampilkanAktivitas(daftar arraktivitas, jumlah int) {
	var i int
	fmt.Println("\033[32m")
	fmt.Println("╔══════╦══════════════╦══════════════════════════════════════════════════════════╦════════════╗")
	fmt.Println("║  No  ║   Tanggal    ║                          Aktivitas                       ║    Skor    ║")
	fmt.Println("╠══════╬══════════════╬══════════════════════════════════════════════════════════╬════════════╣")

	for i = 0; i < jumlah; i++ {
		fmt.Printf("║ %2d   ║  %04d/%02d/%02d  ║ %-56s ║    %3d     ║\n",
			i+1, daftar[i].year, daftar[i].month, daftar[i].day,
			daftar[i].nama, daftar[i].skorDampak)
	}

	fmt.Println("╚══════╩══════════════╩══════════════════════════════════════════════════════════╩════════════╝")
	fmt.Println("\033[0m")
}
func tanggal(daftar arraktivitas, i int) int {
	var year, month, day int
	year = daftar[i].year * 10000
	month = daftar[i].month * 100
	day = daftar[i].day
	return year + month + day
}

func months(daftar arraktivitas, i int) int {
	var year, month int
	year = daftar[i].year * 100
	month = daftar[i].month
	return year + month
}
func tampilkanTanggalUnik(daftar arraktivitas, jumlah int) {
	var arrdate [MaksData]int
	var tanggalUnik [MaksData]int
	var jumlahTanggalUnik int = 0
	var i, j int
	var sudahAda bool

	fmt.Println("\033[32m")
	fmt.Println("╔══════════════════════════╗")
	fmt.Println("║  📅 Tanggal-tanggal 📅   ║")
	fmt.Println("║ yang Sudah Ada Aktivitas ║")
	fmt.Println("╠══════════════════════════╣")

	for i = 0; i < jumlah; i++ {
		arrdate[i] = tanggal(daftar, i) // Mengonversi tanggal ke angka
	}
	for i = 0; i < jumlah; i++ {
		sudahAda = false
		for j = 0; j < jumlahTanggalUnik; j++ {
			if arrdate[i] == tanggalUnik[j] { // Cek apakah tanggal sudah ada di tanggalUnik
				sudahAda = true
			}
		}

		// Jika tanggal belum ada, simpan ke tanggalUnik
		if !sudahAda {
			tanggalUnik[jumlahTanggalUnik] = arrdate[i]
			jumlahTanggalUnik++
		}
	}

	if jumlahTanggalUnik == 0 {
		fmt.Printf("║ %-68s ║\n", "⚠️  Belum ada aktivitas yang tercatat.")
	} else {
		// Menampilkan tanggal unik yang ditemukan
		for i = 0; i < jumlahTanggalUnik; i++ {
			fmt.Printf("║ 🔹 %d/%d/%d ║\n",
				(tanggalUnik[i] / 10000),   // Menampilkan Tahun
				(tanggalUnik[i]%10000)/100, // Menampilkan Bulan
				tanggalUnik[i]%100)
		}
	}

	fmt.Println("╚══════════════════════════╝")
	fmt.Println("\033[0m")
}

func editAktivitas(daftar *arraktivitas, jumlah int, kategori arrkat) {
	tampilkanAktivitas(*daftar, jumlah)
	var selesai bool
	var year, month, day, tanggaledit int
	var arrdate [MaksData]int
	var i, idx int

	fmt.Println("\033[32m")
	fmt.Println("╔═════════════════════════════════════════╗")
	fmt.Println("║ ✏️  Ubah Aktivitas Berdasarkan Tanggal   ║")
	fmt.Println("╚═════════════════════════════════════════╝")
	fmt.Println("\033[0m")

	tampilkanTanggalUnik(*daftar, jumlah)
	fmt.Println("\nKetik x untuk kembali.\n")
	for !selesai {
		fmt.Println("\033[32m")
		fmt.Println("╔════════════════════════════════════════╗")
		fmt.Println("║ 📅 Silakan masukkan tanggal aktivitas  ║")
		fmt.Println("╚════════════════════════════════════════╝")
		fmt.Println("\033[0m")
		fmt.Print("➡️  Tanggal (YYYY/MM/DD): ")
		var inputTanggal string
		fmt.Scan(&inputTanggal)
		if inputTanggal == "x" {
			selesai = true
			refreshTampilan(*daftar, jumlah, " ")
		} else {
			fmt.Sscanf(inputTanggal, "%d/%d/%d", &year, &month, &day)
			tanggaledit = year*10000 + month*100 + day
			var indeksAktivitas [MaksData]int
			var totalDitemukan int = 0

			for i = 0; i < jumlah; i++ {
				arrdate[i] = tanggal(*daftar, i)
				if arrdate[i] == tanggaledit {
					indeksAktivitas[totalDitemukan] = i
					totalDitemukan++
				}
			}

			if totalDitemukan == 0 {
				selesai = true
				refreshTampilan(*daftar, jumlah, "⚠️ Tidak ada aktivitas ditemukan pada tanggal tersebut.")
			} else {
				fmt.Println("\033[32m")
				fmt.Println("╔══════════════════════════════════════════════════╗")
				fmt.Println("║ 📋 Daftar aktivitas pada tanggal tersebut:       ║")
				fmt.Println("║══════════════════════════════════════════════════║")
				for i = 0; i < totalDitemukan; i++ {
					idx = indeksAktivitas[i]
					fmt.Printf("║ %2d. %-30.30s (Skor: %2d)    ║\n", i+1, (*daftar)[idx].nama, (*daftar)[idx].skorDampak)
				}
				fmt.Println("╚══════════════════════════════════════════════════╝")
				fmt.Println("\033[0m")

				var pilihanAktivitas int
				fmt.Print("Pilih nomor aktivitas yang ingin diedit: ")
				fmt.Scan(&pilihanAktivitas)

				if pilihanAktivitas >= 1 && pilihanAktivitas <= totalDitemukan {
					fmt.Println("\033[32m")
					fmt.Println("╔════════════════════════════════════════════════════╗")
					fmt.Println("║ 📌 Pilih aktivitas baru dari daftar kategori:      ║")
					fmt.Println("╚════════════════════════════════════════════════════╝")
					fmt.Println("\033[0m")
					listKategori(kategori)
					fmt.Print("🌿 Jawaban anda ➤ ")
					fmt.Println("\033[0m")
					var pilihanKategori int
					fmt.Print("🌿 Masukkan pilihan kategori baru: ")
					fmt.Scan(&pilihanKategori)

					if pilihanKategori >= 1 && pilihanKategori <= len(kategori) {
						var idxEdit int = indeksAktivitas[pilihanAktivitas-1]
						(*daftar)[idxEdit].nama = kategori[pilihanKategori-1].nama
						(*daftar)[idxEdit].skorDampak = kategori[pilihanKategori-1].skorDampak
						selesai = true
						refreshTampilan(*daftar, jumlah, "✅ Aktivitas berhasil diedit.")
					} else {
						selesai = true
						refreshTampilan(*daftar, jumlah, "❌ Pilihan kategori tidak valid.")

					}
				} else {
					selesai = true
					refreshTampilan(*daftar, jumlah, "❌ Pilihan aktivitas tidak valid.")
				}
			}
		}

	}

}

func hapusAktivitas(daftar *arraktivitas, jumlah *int) {
	tampilkanAktivitas(*daftar, *jumlah)
	var selesai bool
	var year, month, day, tanggalhapus int
	var arrdate [MaksData]int

	fmt.Println("\033[32m")
	fmt.Println("╔═════════════════════════════════════════╗")
	fmt.Println("║ ❌ Hapus Aktivitas Berdasarkan Tanggal  ║")
	fmt.Println("╚═════════════════════════════════════════╝")
	fmt.Println("\033[0m")

	tampilkanTanggalUnik(*daftar, *jumlah)
	fmt.Println("Ketik x untuk kembali.")
	for !selesai {
		fmt.Println("\033[32m")
		fmt.Println("╔════════════════════════════════════════════════════╗")
		fmt.Println("║ 📅 Masukkan tanggal aktivitas yang ingin dihapus   ║")
		fmt.Println("╚════════════════════════════════════════════════════╝")
		fmt.Println("\033[0m")
		fmt.Print("➡️  Tanggal (YYYY/MM/DD): ")
		var inputTanggal string
		fmt.Scan(&inputTanggal)
		if inputTanggal == "x" {
			selesai = true
			refreshTampilan(*daftar, *jumlah, " ")
		} else {
			fmt.Sscanf(inputTanggal, "%d/%d/%d", &year, &month, &day)
			tanggalhapus = year*10000 + month*100 + day
			var indeksAktivitas [MaksData]int
			var totalDitemukan int = 0
			var i int
			for i = 0; i < *jumlah; i++ {
				arrdate[i] = tanggal(*daftar, i)
				if arrdate[i] == tanggalhapus {
					indeksAktivitas[totalDitemukan] = i
					totalDitemukan++
				}
			}
			if totalDitemukan == 0 {
				fmt.Println("⚠️ Tidak ada aktivitas ditemukan pada tanggal tersebut.")
			} else {
				fmt.Println("\033[32m")
				fmt.Println("╔══════════════════════════════════════════════════╗")
				fmt.Println("║ 🗑️  Daftar aktivitas pada tanggal tersebut:       ║")
				fmt.Println("║══════════════════════════════════════════════════║")
				for i = 0; i < totalDitemukan; i++ {
					idx := indeksAktivitas[i]
					fmt.Printf("║ %2d. %-30.30s (Skor: %2d)    ║\n", i+1, (*daftar)[idx].nama, (*daftar)[idx].skorDampak)
				}
				fmt.Println("╚══════════════════════════════════════════════════╝")
				fmt.Println("\033[0m")
				var pilihanAktivitas int
				fmt.Print("🗂️  Pilih nomor aktivitas yang ingin dihapus: ")
				fmt.Scan(&pilihanAktivitas)

				if pilihanAktivitas == 0 {
					selesai = true
				}

				if pilihanAktivitas >= 1 && pilihanAktivitas <= totalDitemukan {
					var idxHapus int = indeksAktivitas[pilihanAktivitas-1]
					for i = idxHapus; i < *jumlah-1; i++ {
						(*daftar)[i] = (*daftar)[i+1]
					}
					*jumlah--
					selesai = true
					refreshTampilan(*daftar, *jumlah, "✅ Aktivitas berhasil dihapus.")
				} else {
					selesai = true
					refreshTampilan(*daftar, *jumlah, "❌ Pilihan aktivitas tidak valid.")
				}
			}
		}

	}

}

func cari(daftar arraktivitas, jumlah int) {
	tampilkanAktivitas(daftar, jumlah)
	var cari string
	var selesai bool
	fmt.Println("\n\nKetik x untuk kembali.")
	fmt.Println("\033[32m")
	fmt.Println("╔════════════════════════════════════════════════════╗")
	fmt.Println("║ 🔍 Cari Aktivitas Berdasarkan:                    ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")
	fmt.Println("\033[0m")
	for !selesai {
		fmt.Println("📌 Pilih pencarian berdasarkan:")
		fmt.Println("   📅 Tanggal (t) atau 🔢 Skor Dampak (d)")
		fmt.Print("🔍 Pilih (t/d): ")
		fmt.Scan(&cari)
		for cari != "t" && cari != "d" && cari != "x" {
			fmt.Println("⚠️ Input tidak valid. Coba lagi.")
			fmt.Print("🔍 Pilih (t/d): ")
			fmt.Scan(&cari)
		}
		if cari == "t" {
			seqAktivitasTanggal(daftar, jumlah)
			selesai = true
		} else if cari == "d" {
			binarySearchPoin(daftar, jumlah)
			selesai = true
		} else if cari == "x" {
			selesai = true
			refreshTampilan(daftar, jumlah, " ")
		}
	}
}

func seqAktivitasTanggal(daftar arraktivitas, jumlah int) {
	var ditemukan bool = false
	var i int
	var year, month, day, tanggalsort int
	var arrdate [MaksData]int
	tampilkanTanggalUnik(daftar, jumlah)
	fmt.Println("\n=== Cari Aktivitas ===")
	fmt.Print("Masukkan kata kunci tanggal (YYYY/MM/DD):  ")
	fmt.Scanf("\n%d/%d/%d", &year, &month, &day)
	tanggalsort = year*10000 + month*100 + day
	refreshTampilan(daftar, jumlah, " ")
	fmt.Println("\nHasil pencarian:")
	for i = 0; i < jumlah; i++ {
		arrdate[i] = tanggal(daftar, i)
		if arrdate[i] == tanggalsort {
			ditemukan = true
			fmt.Printf("- %d/%d/%d | %s | Skor: %d\n",
				daftar[i].year, daftar[i].month, daftar[i].day, daftar[i].nama, daftar[i].skorDampak)
		}
	}
	if ditemukan == false {
		refreshTampilan(daftar, jumlah, "⚠️ Tidak ditemukan aktivitas yang sesuai.")
	}
}

func binarySearchPoin(daftar arraktivitas, jumlah int) {
	var x, left, right, mid, found int
	fmt.Print("Cari aktivitas berdasarkan skor dampak (6–15): ")
	fmt.Scan(&x)
	sortSkorAscend(&daftar, jumlah)
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
		refreshTampilan(daftar, jumlah, "⚠️ Skor dampak tidak ditemukan.")
	} else {
		refreshTampilan(daftar, jumlah, " ")
		var i int = 0
		fmt.Println("✅ Ditemukan aktivitas dengan skor dampak", x, ":")
		fmt.Printf("- %d/%d/%d | %s | Skor: %d\n",
			daftar[found].year, daftar[found].month, daftar[found].day, daftar[found].nama, daftar[found].skorDampak)

		for i = found - 1; i >= 0 && daftar[i].skorDampak == x; i-- {
			fmt.Printf("- %d/%d/%d | %s | Skor: %d\n",
				daftar[i].year, daftar[i].month, daftar[i].day, daftar[i].nama, daftar[i].skorDampak)
		}
		for i = found + 1; i < jumlah && daftar[i].skorDampak == x; i++ {
			fmt.Printf("- %d/%d/%d | %s | Skor: %d\n",
				daftar[i].year, daftar[i].month, daftar[i].day, daftar[i].nama, daftar[i].skorDampak)
		}
	}
}

func sorting(daftar *arraktivitas, jumlah int) {
	tampilkanAktivitas(*daftar, jumlah)
	var selesai bool
	var sort string
	fmt.Println("\n\nKetik x untuk kembali.")
	for !selesai {
		fmt.Println("🔃 Pilih pengurutan berdasarkan:")
		fmt.Println("    📅 Tanggal (t) , Frekuensi (f), atau 🟢 Skor Dampak (d)")
		fmt.Print("🔽 Pilih (t/f/d): ")
		fmt.Scan(&sort)
		for sort != "t" && sort != "f" && sort != "d" && sort != "x" {
			fmt.Println("⚠️ Input tidak valid. Coba lagi.")
			fmt.Print("🔽 Pilih (t/f/d): ")
			fmt.Scan(&sort)
		}
		if sort == "t" {
			sortTanggal(daftar, jumlah)
			selesai = true
		} else if sort == "f" {
			sortFrekuensi(daftar, jumlah)
			selesai = true
		} else if sort == "d" {
			sortSkor(daftar, jumlah)
			selesai = true
		} else if sort == "x" {
			selesai = true
			refreshTampilan(*daftar, jumlah, " ")
		}
	}

}
func sortFrekuensi(daftar *arraktivitas, jumlah int) {
	var urut string
	fmt.Println("Pilih daftarkan frekuensi secara decending(d) atau ascending(a)")
	fmt.Print("Pilih salah satu (d/a):")
	fmt.Scan(&urut)
	if urut == "d" {
		sortFrekuensiDescend(daftar, jumlah)
	} else if urut == "a" {
		sortFrekuensiAscend(daftar, jumlah)
	}
	refreshTampilan(*daftar, jumlah, "✅ Aktivitas berhasil diurutkan berdasarkan frekuensi aktivitas.")
}
func hitungFrekuensi(daftar arraktivitas, jumlah int, nama string) int {
	var count, i int
	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			count++
		}
	}
	return count
}
func sortFrekuensiDescend(daftar *arraktivitas, jumlah int) {
	var namaUnik [MaksData]string
	var frekuensi [MaksData]int
	var jumlahUnik int = 0
	var i, j int
	var sudahAda bool

	// 1. Kumpulkan nama unik & frekuensinya
	for i = 0; i < jumlah; i++ {
		sudahAda = false
		for j = 0; j < jumlahUnik; j++ {
			if daftar[i].nama == namaUnik[j] {
				sudahAda = true
			}
		}
		if !sudahAda {
			namaUnik[jumlahUnik] = daftar[i].nama
			frekuensi[jumlahUnik] = hitungFrekuensi(*daftar, jumlah, daftar[i].nama)
			jumlahUnik++
		}
	}

	// 2. Urutkan namaUnik berdasarkan frekuensi descending
	var tempStr string
	var tempInt int
	for i = 0; i < jumlahUnik-1; i++ {
		for j = i + 1; j < jumlahUnik; j++ {
			if frekuensi[i] < frekuensi[j] {
				// Tukar
				tempStr = namaUnik[i]
				namaUnik[i] = namaUnik[j]
				namaUnik[j] = tempStr

				tempInt = frekuensi[i]
				frekuensi[i] = frekuensi[j]
				frekuensi[j] = tempInt
			}
		}
	}

	// 3. Susun daftar baru berdasarkan urutan nama unik
	var hasil arraktivitas
	var idx int = 0
	for i = 0; i < jumlahUnik; i++ {
		for j = 0; j < jumlah; j++ {
			if daftar[j].nama == namaUnik[i] {
				hasil[idx] = daftar[j]
				idx++
			}
		}
	}

	// 4. Copy kembali ke daftar asli
	for i = 0; i < jumlah; i++ {
		(*daftar)[i] = hasil[i]
	}
}

func sortFrekuensiAscend(daftar *arraktivitas, jumlah int) {
	var namaUnik [MaksData]string
	var frekuensi [MaksData]int
	var jumlahUnik int = 0
	var i, j int
	var sudahAda bool

	// 1. Kumpulkan nama unik & frekuensinya
	for i = 0; i < jumlah; i++ {
		sudahAda = false
		for j = 0; j < jumlahUnik; j++ {
			if daftar[i].nama == namaUnik[j] {
				sudahAda = true
			}
		}
		if !sudahAda {
			namaUnik[jumlahUnik] = daftar[i].nama
			frekuensi[jumlahUnik] = hitungFrekuensi(*daftar, jumlah, daftar[i].nama)
			jumlahUnik++
		}
	}

	// 2. Urutkan namaUnik berdasarkan frekuensi descending
	var tempStr string
	var tempInt int
	for i = 0; i < jumlahUnik-1; i++ {
		for j = i + 1; j < jumlahUnik; j++ {
			if frekuensi[i] > frekuensi[j] {
				// Tukar
				tempStr = namaUnik[i]
				namaUnik[i] = namaUnik[j]
				namaUnik[j] = tempStr

				tempInt = frekuensi[i]
				frekuensi[i] = frekuensi[j]
				frekuensi[j] = tempInt
			}
		}
	}

	// 3. Susun daftar baru berdasarkan urutan nama unik
	var hasil arraktivitas
	var idx int = 0
	for i = 0; i < jumlahUnik; i++ {
		for j = 0; j < jumlah; j++ {
			if daftar[j].nama == namaUnik[i] {
				hasil[idx] = daftar[j]
				idx++
			}
		}
	}

	// 4. Copy kembali ke daftar asli
	for i = 0; i < jumlah; i++ {
		(*daftar)[i] = hasil[i]
	}
}

func sortTanggal(daftar *arraktivitas, jumlah int) {
	var urut string
	fmt.Println("Pilih daftarkan tanggal secara decending(d) atau ascending(a)")
	fmt.Print("Pilih salah satu (d/a):")
	fmt.Scan(&urut)
	if urut == "a" {
		sortTanggalAscend(daftar, jumlah)
	} else if urut == "d" {
		sortTanggalDescend(daftar, jumlah)
	}
	refreshTampilan(*daftar, jumlah, "✅ Aktivitas berhasil didaftarkan berdasarkan tanggal.")
}
func sortTanggalAscend(daftar *arraktivitas, jumlah int) {
	var i, j, idx int
	var temp Aktivitas
	for i = 0; i < jumlah-1; i++ {
		idx = i
		for j = i + 1; j < jumlah; j++ {
			if tanggal(*daftar, j) < tanggal(*daftar, idx) {
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

func sortTanggalDescend(daftar *arraktivitas, jumlah int) {
	var i, j, idx int
	var temp Aktivitas
	for i = 0; i < jumlah-1; i++ {
		idx = i
		for j = i + 1; j < jumlah; j++ {
			if tanggal(*daftar, j) > tanggal(*daftar, idx) {
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

func sortSkor(daftar *arraktivitas, jumlah int) {
	var urut string
	fmt.Println("Pilih daftarkan tanggal secara decending(d) atau ascending(a)")
	fmt.Print("Pilih salah satu (d/a):")
	fmt.Scan(&urut)
	if urut == "a" {
		sortSkorAscend(daftar, jumlah)
	} else if urut == "d" {
		sortSkorDescend(daftar, jumlah)
	}
	refreshTampilan(*daftar, jumlah, "✅ Aktivitas berhasil didaftarkan berdasarkan skor dampak aktivitas.")
}
func sortSkorAscend(daftar *arraktivitas, jumlah int) {
	var i, j int
	var temp Aktivitas
	for i = 1; i < jumlah; i++ {
		temp = (*daftar)[i]
		j = i - 1
		for j >= 0 && (*daftar)[j].skorDampak > temp.skorDampak {
			(*daftar)[j+1] = (*daftar)[j]
			j--
		}
		(*daftar)[j+1] = temp
	}
}
func sortSkorDescend(daftar *arraktivitas, jumlah int) {
	var i, j, idx int
	var temp Aktivitas
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
func laporanBulanan(daftar arraktivitas, jumlah int) {
	tampilkanAktivitas(daftar, jumlah)
	sortTanggalAscend(&daftar, jumlah)
	var i, totalSkor, totalAktivitas int
	var bulanSekarang, bulanSebelumnya int
	if jumlah != 0 {
		refreshTampilan(daftar, jumlah, " ")
		fmt.Println("\n===== LAPORAN BULANAN =====")
		for i = 0; i < jumlah; i++ {
			bulanSekarang = months(daftar, i)
			if bulanSekarang != bulanSebelumnya && i != 0 {
				fmt.Printf("📅 Bulan: %d/%d - Total Aktivitas: %d - Total Skor: %d\n",
					bulanSebelumnya/100, bulanSebelumnya%100, totalAktivitas, totalSkor)
				totalSkor = 0
				totalAktivitas = 0
			}
			totalSkor += daftar[i].skorDampak
			totalAktivitas++
			bulanSebelumnya = bulanSekarang
		}
		fmt.Printf("📅 Bulan: %d/%d - Total Aktivitas: %d - Total Skor: %d\n",
			bulanSebelumnya/100, bulanSebelumnya%100, totalAktivitas, totalSkor)

	} else {
		refreshTampilan(daftar, jumlah, "⚠️ Belum ada aktivitas untuk dihitung.")
	}

}
func dummyData(daftar *arraktivitas, jumlah *int) {
	daftar[*jumlah] = Aktivitas{2025, 5, 21, "Membuang sampah pada tempatnya", 6}
	*jumlah++
	daftar[*jumlah] = Aktivitas{2025, 5, 21, "Menanam pohon", 15}
	*jumlah++
	daftar[*jumlah] = Aktivitas{2025, 5, 21, "Membawa tumbler", 7}
	*jumlah++
	daftar[*jumlah] = Aktivitas{2025, 6, 2, "Menggunakan energi terbarukan", 14}
	*jumlah++
	daftar[*jumlah] = Aktivitas{2025, 6, 2, "Mengikuti kegiatan bersih-bersih lingkungan", 12}
	*jumlah++
	daftar[*jumlah] = Aktivitas{2025, 6, 2, "Menanam hidroponik", 11}
	*jumlah++
	daftar[*jumlah] = Aktivitas{2025, 6, 3, "Mendaur ulang sampah", 10}
	*jumlah++
	daftar[*jumlah] = Aktivitas{2025, 7, 1, "Naik sepeda / angkutan umum", 9}
	*jumlah++
	daftar[*jumlah] = Aktivitas{2025, 7, 1, "Membuang sampah pada tempatnya", 6}
	*jumlah++
	daftar[*jumlah] = Aktivitas{2025, 7, 1, "Menanam pohon", 15}
	*jumlah++
	daftar[*jumlah] = Aktivitas{2025, 7, 3, "Membawa tumbler", 7}
	*jumlah++
	daftar[*jumlah] = Aktivitas{2025, 7, 3, "Membuang sampah pada tempatnya", 6}
	*jumlah++
	daftar[*jumlah] = Aktivitas{2025, 7, 3, "Membawa tumbler", 7}
	*jumlah++
}
