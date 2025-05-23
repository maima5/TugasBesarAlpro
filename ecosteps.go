package main

import "fmt"

type Aktivitas struct {
	year, month, day int
	nama             string
	skorDampak       int
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
	var daftarAktivitas [MaksData]Aktivitas
	var jumlahData, pilihan int
	var selesai, dataSudahAda bool

	for !selesai {
		fmt.Println("\033[32m") // Warna hijau
		fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                      📋  M E N U   U T A M A  📋                     ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║  📌 1. Tambah Aktivitas                                              ║")
		fmt.Println("║  📄 2. Tampilkan Semua Aktivitas                                     ║")
		fmt.Println("║  ✏️  3. Edit Aktivitas Berdasarkan Tanggal                            ║")
		fmt.Println("║  ❌ 4. Hapus Aktivitas Berdasarkan Tanggal                           ║")
		fmt.Println("║  🔍 5. Cari Aktivitas                                                ║")
		fmt.Println("║  🔃 6. Urutkan Aktivitas                                             ║")
		fmt.Println("║  📅 7. Laporan Bulanan                                               ║")
		fmt.Println("║  🔙 8. Kembali ke Halaman Awal                                       ║")
		fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
		fmt.Print("🌱 Pilihan anda ➤ ")
		fmt.Println("\033[0m") // Reset warna

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
					editAktivitas(&daftarAktivitas, jumlahData, daftarKategori)
				} else if pilihan == 4 {
					hapusAktivitas(&daftarAktivitas, &jumlahData)
				} else if pilihan == 5 {
					cari(daftarAktivitas, jumlahData)
				} else if pilihan == 6 {
					sorting(&daftarAktivitas, jumlahData)
				} else if pilihan == 7 {
					laporanBulanan(daftarAktivitas, jumlahData)
				} else {
					fmt.Println("⚠️ Pilihan tidak valid.")
				}
			}
		} else if pilihan == 8 {
			selesai = true
		} else {
			fmt.Println("⚠️ Pilihan tidak valid.")
		}
	}
}

func tentangEcoSteps() {
	fmt.Println("\033[32m") // Warna hijau
	fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    📖 TENTANG APLIKASI ECOSTEPS                      ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║  EcoSteps adalah aplikasi pelacak gaya hidup ramah lingkungan.       ║")
	fmt.Println("║  Dengan aplikasi ini, Anda dapat memantau aktivitas sehari-hari,     ║")
	fmt.Println("║  menghitung skor keberlanjutan, dan melihat laporan bulanan. 🌿      ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
	fmt.Println("\033[0m") // Reset warna
}

func listKategori(kategori [10]Kategori) {
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

func tambahAktivitas(daftar *[MaksData]Aktivitas, jumlah *int, kategori *[10]Kategori) {
	var lanjutTanggal string = "y"
	var year, month, day int

	for lanjutTanggal == "y" && *jumlah < MaksData {
		fmt.Println("\033[32m")
		fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                    📅 Input Tanggal Aktivitas                        ║")
		fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
		fmt.Print("Tanggal (YYYY/MM/DD): ")
		fmt.Println("\033[0m")
		fmt.Scanf("%d/%d/%d", year, month, day)

		var lanjutAktivitas string = "y"
		for lanjutAktivitas == "y" && *jumlah < MaksData {
			listKategori(*kategori)
			fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
			fmt.Printf("║ %-69s ║\n", "⚠️  Pilih hanya SATU aktivitas!")
			fmt.Printf("║ %-69s ║\n", "➡️  Masukkan nomor aktivitas:")
			fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
			fmt.Print("🌿 ➤ ")
			fmt.Println("\033[0m")

			var pilihan int
			fmt.Scan(&pilihan)

			if pilihan >= 1 && pilihan <= len(*kategori) {
				(*daftar)[*jumlah].year = year
				(*daftar)[*jumlah].month = month
				(*daftar)[*jumlah].day = day
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
				fmt.Println("║        ➕ Tambah aktivitas lain untuk tanggal ini? (y/n)             ║")
				fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
				fmt.Print("🌿 Jawaban anda ➤ ")
				fmt.Println("\033[0m")
				fmt.Scan(&lanjutAktivitas)
			} else {
				fmt.Println("⚠️ Kapasitas data sudah penuh.")
				lanjutAktivitas = "n"
			}
		}

		if *jumlah < MaksData {
			fmt.Println("\033[32m")
			fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
			fmt.Println("║        📅 Input aktivitas untuk tanggal lain? (y/n)                  ║")
			fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
			fmt.Print("🌿 Jawaban anda ➤ ")
			fmt.Println("\033[0m")
			fmt.Scan(&lanjutTanggal)
		}
	}
}

func tampilkanAktivitas(daftar [MaksData]Aktivitas, jumlah int) {
	var i int
	fmt.Println("\n=== Daftar Semua Aktivitas ===")
	for i = 0; i < jumlah; i++ {
		fmt.Printf("%2d. Tanggal: %-12d/%d/%d | Aktivitas: %-40s | Skor: %d\n",
			i+1, daftar[i].year, daftar[i].month, daftar[i].day, daftar[i].nama, daftar[i].skorDampak)
	}
}
func tanggal(daftar [MaksData]Aktivitas, i int) int {
	var year, month, day int
	year = daftar[i].year * 360
	month = daftar[i].month * 30
	day = daftar[i].day
	return year + month + day
}
func months(daftar [MaksData]Aktivitas, i int) int {
	var year, month int
	year = daftar[i].year * 24
	month = daftar[i].month
	return year + month
}
func tampilkanTanggalUnik(daftar [MaksData]Aktivitas, jumlah int) {
	var date int
	var arrdate [MaksData]int
	var tanggalUnik [MaksData]int
	var jumlahTanggalUnik int = 0
	var i, j int
	var sudahAda bool

	fmt.Println("\033[32m")
	fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║               📅 Tanggal-tanggal yang Sudah Ada Aktivitas            ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════════╣")

	for i = 0; i < jumlah; i++ {
		date = tanggal(daftar, i)
		arrdate[i] = date
		sudahAda = false
		for j = 0; !sudahAda && j < jumlahTanggalUnik; j++ {
			if arrdate[i] == tanggalUnik[j] {
				sudahAda = true
			}
		}
		if !sudahAda {
			tanggalUnik[jumlahTanggalUnik] = arrdate[i]
			jumlahTanggalUnik++
		}
	}

	if jumlahTanggalUnik == 0 {
		fmt.Printf("║ %-68s ║\n", "⚠️  Belum ada aktivitas yang tercatat.")
	} else {
		for i = 0; i < jumlahTanggalUnik; i++ {
			fmt.Printf("║ 🔹 %-65d/%d/%d ║\n", tanggalUnik[i]/360, (tanggalUnik[i]%360)/24, ((tanggalUnik[i] % 360) % 24))
		}
	}

	fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")
	fmt.Println("\033[0m")
}

func editAktivitas(daftar *[MaksData]Aktivitas, jumlah int, kategori [10]Kategori) {
	var year, month, day int
	var date, i, idx, tanggaledit int
	var arrdate [MaksData]int
	tampilkanTanggalUnik(*daftar, jumlah)
	fmt.Print("Tanggal (YYYY/MM/DD): ")
	fmt.Scanf("%d/%d/%d", year, month, day)
	tanggaledit = year*360 + month*30 + day
	var indeksAktivitas [MaksData]int
	var totalDitemukan int = 0

	for i = 0; i < jumlah; i++ {
		date = tanggal(*daftar, i)
		arrdate[i] = date
		if arrdate[i] == tanggaledit {
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
			fmt.Println("\nPilih aktivitas baru:")
			listKategori(kategori)
			fmt.Print("🌿 Jawaban anda ➤ ")
			fmt.Println("\033[0m")
			var pilihanKategori int
			fmt.Print("Masukkan pilihan kategori baru: ")
			fmt.Scan(&pilihanKategori)

			if pilihanKategori >= 1 && pilihanKategori <= len(kategori) {
				var idxEdit int = indeksAktivitas[pilihanAktivitas-1]
				(*daftar)[idxEdit].nama = kategori[pilihanKategori-1].nama
				(*daftar)[idxEdit].skorDampak = kategori[pilihanKategori-1].skorDampak
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
	var year, month, day, tanggaledit, date int
	var arrdate [MaksData]int
	tampilkanTanggalUnik(*daftar, *jumlah)
	fmt.Print("Tanggal (YYYY/MM/DD): ")
	fmt.Scanf("%d/%d/%d", year, month, day)
	tanggaledit = year*360 + month*30 + day
	var indeksAktivitas [MaksData]int
	var totalDitemukan int = 0
	var i, idx int
	for i = 0; i < *jumlah; i++ {
		date = tanggal(*daftar, i)
		arrdate[i] = date
		if arrdate[i] == tanggaledit {
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
func cari(daftar [MaksData]Aktivitas, jumlah int) {
	var cari string
	fmt.Println("\n\n📌 Pilih cari aktivitas berdasarkan tanggal (t) atau skor dampak (d).")
	fmt.Print("🔍 Pilih (t/d): ")
	fmt.Scan(&cari)
	for cari != "t" && cari != "d" {
		fmt.Println("⚠️ Input tidak valid. Coba lagi.")
		fmt.Print("🔍 Pilih (t/d): ")
		fmt.Scan(&cari)
	}
	if cari == "t" {
		seqAktivitasTanggal(daftar, jumlah)
	} else if cari == "d" {
		binarySearchPoin(daftar, jumlah)
	}
}

func seqAktivitasTanggal(daftar [MaksData]Aktivitas, jumlah int) {
	var ditemukan bool = false
	var i, year, month, day, tanggalsort, date int
	var arrdate [MaksData]int
	tampilkanTanggalUnik(daftar, jumlah)
	fmt.Println("\n=== Cari Aktivitas ===")
	fmt.Print("Masukkan kata kunci tanggal (YYYY/MM/DD):  ")
	fmt.Scanf("%d/%d/%d", year, month, day)
	tanggalsort = year*360 + month*30 + day

	fmt.Println("\nHasil pencarian:")
	for i = 0; i < jumlah; i++ {
		date = tanggal(daftar, i)
		arrdate[i] = date
		if arrdate[i] == tanggalsort {
			ditemukan = true
			fmt.Printf("- %d/%d/%d | %s | Skor: %d\n",
				daftar[i].year, daftar[i].month, daftar[i].day, daftar[i].nama, daftar[i].skorDampak)
		}
	}

	if ditemukan == false {
		fmt.Println("⚠️ Tidak ditemukan aktivitas yang sesuai.")
	}
}

func binarySearchPoin(daftar [MaksData]Aktivitas, jumlah int) {
	var x, i, left, right, mid, found int
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
		fmt.Println("⚠️ Skor dampak tidak ditemukan.")
	} else {
		fmt.Println("✅ Ditemukan aktivitas dengan skor dampak", x, ":")

		fmt.Printf("- %d/%d/%d | %s | Skor: %d\n",
			daftar[i].year, daftar[i].month, daftar[i].day, daftar[i].nama, daftar[i].skorDampak)

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

func sorting(daftar *[MaksData]Aktivitas, jumlah int) {
	var sort string
	fmt.Println("\n\n🔃 Pilih pengurutan berdasarkan:")
	fmt.Println("    📅 Tanggal (t), 🔁 Frekuensi (f), atau 🟢 Skor Dampak (d)")
	fmt.Print("🔽 Pilih (t/f/d): ")
	fmt.Scan(&sort)
	for sort != "t" && sort != "f" && sort != "d" {
		fmt.Println("⚠️ Input tidak valid. Coba lagi.")
		fmt.Print("🔽 Pilih (t/f/d): ")
		fmt.Scan(&sort)
	}
	if sort == "t" {
		sortTanggal(daftar, jumlah)
	} else if sort == "f" {
		sortFrekuensi(daftar, jumlah)
	} else if sort == "d" {
		sortSkor(daftar, jumlah)
	}
}
func sortTanggal(daftar *[MaksData]Aktivitas, jumlah int) {
	var urut string
	fmt.Println("Pilih daftarkan tanggal secara decending(d) atau ascending(a)")
	fmt.Print("Pilih salah satu (d/a):")
	fmt.Scan(&urut)
	if urut == "a" {
		sortTanggalAscend(daftar, jumlah)
	} else if urut == "d" {
		sortFrekuensiDescend(daftar, jumlah)
	}
	tampilkanAktivitas(*daftar, jumlah)
	fmt.Println("✅ Aktivitas berhasil didaftarkan berdasarkan tanggal.")
}
func sortTanggalAscend(daftar *[MaksData]Aktivitas, jumlah int) {
	var i, j, idx, date int
	var arrdate [MaksData]int
	var temp Aktivitas
	for i = 0; i < jumlah; i++ {
		date = tanggal(*daftar, i)
		arrdate[i] = date
	}
	for i = 0; i < jumlah-1; i++ {
		idx = i
		for j = i + 1; j < jumlah; j++ {
			if arrdate[j] < arrdate[idx] {
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
func sortTanggalDescend(daftar *[MaksData]Aktivitas, jumlah int) {
	var date int
	var arrdate [MaksData]int
	var i, j, idx int
	var temp Aktivitas
	for i = 0; i < jumlah; i++ {
		date = tanggal(*daftar, i)
		arrdate[i] = date
	}
	for i = 0; i < jumlah-1; i++ {
		idx = i
		for j = i + 1; j < jumlah; j++ {
			if arrdate[j] > arrdate[idx] {
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
func sortFrekuensi(daftar *[MaksData]Aktivitas, jumlah int) {
	var urut string
	fmt.Println("Pilih daftarkan frekuensi secara decending(d) atau ascending(a)")
	fmt.Print("Pilih salah satu (d/a):")
	fmt.Scan(&urut)
	if urut == "d" {
		sortFrekuensiDescend(daftar, jumlah)
	} else if urut == "a" {
		sortFrekuensiAscend(daftar, jumlah)
	}
	tampilkanAktivitas(*daftar, jumlah)
	fmt.Println("✅ Aktivitas berhasil didaftarkan berdasarkan frekuensi aktivitas.")
}
func hitungFrekuensi(daftar [MaksData]Aktivitas, jumlah int, nama string) int {
	var count, i int
	for i = 0; i < jumlah; i++ {
		if daftar[i].nama == nama {
			count++
		}
	}
	return count
}
func sortFrekuensiDescend(daftar *[MaksData]Aktivitas, jumlah int) {
	var i, j, freqTemp, freqJ int
	var temp Aktivitas
	var geser bool
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
}
func sortFrekuensiAscend(daftar *[MaksData]Aktivitas, jumlah int) {
	var i, j, freqTemp, freqJ int
	var temp Aktivitas
	var geser bool
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

func sortSkor(daftar *[MaksData]Aktivitas, jumlah int) {
	var urut string
	fmt.Println("Pilih daftarkan tanggal secara decending(d) atau ascending(a)")
	fmt.Print("Pilih salah satu (d/a):")
	fmt.Scan(&urut)
	if urut == "a" {
		sortSkorAscend(daftar, jumlah)
	} else if urut == "d" {
		sortSkorDescend(daftar, jumlah)
	}
	tampilkanAktivitas(*daftar, jumlah)
	fmt.Println("✅ Aktivitas berhasil didaftarkan berdasarkan skor dampak aktivitas.")
}
func sortSkorAscend(daftar *[MaksData]Aktivitas, jumlah int) {
	var i, j, idx int
	var temp Aktivitas
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
}
func sortSkorDescend(daftar *[MaksData]Aktivitas, jumlah int) {
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
func laporanBulanan(daftar [MaksData]Aktivitas, jumlah int) {
	sortTanggalAscend(&daftar, jumlah)
	var i, totalSkor, totalAktivitas int
	var bulanSekarang, bulanSebelumnya int
	if jumlah != 0 {
		fmt.Println("\n===== LAPORAN BULANAN =====")
		for i = 0; i < jumlah; i++ {
			bulanSekarang = months(daftar, i)
			if bulanSekarang != bulanSebelumnya && i != 0 {
				fmt.Printf("📅 Bulan: %d/%d - Total Aktivitas: %d - Total Skor: %d\n",
					bulanSebelumnya/24, bulanSebelumnya%24, totalAktivitas, totalSkor)
				totalSkor = 0
				totalAktivitas = 0
			}
			totalSkor += daftar[i].skorDampak
			totalAktivitas++
			bulanSebelumnya = bulanSekarang
		}
		fmt.Printf("📅 Bulan: %d/%d - Total Aktivitas: %d - Total Skor: %d\n",
			bulanSebelumnya/24, bulanSebelumnya%24, totalAktivitas, totalSkor)
	} else {
		fmt.Println("⚠️ Belum ada aktivitas untuk dihitung.")
	}
}
