package main

import "fmt"

const NMAX int = 10
const ADMIN_USERNAME string = "admin"
const ADMIN_PASSWORD string = "12345"

type Kandidat struct {
	noUrut   int
	nama     string
	suara    int
	prodi    string
	fakultas string
	angkatan int
}

type Konfigurasi struct {
	fakultas string
}

type DataKandidat struct {
	data [NMAX]Kandidat
	n    int
}

// TAMPILAN
func garisAtas(simbol string) {
	fmt.Println(simbol)
}

func garisBawah(simbol string) {
	fmt.Println(simbol)
}

func judulProgram(judul string) {
	garisAtas("==========================================================================================================")
	fmt.Println(judul)
	garisBawah("==========================================================================================================")
}

func tampilkanHeaderKandidat(judulKolom string) {
	garisAtas("==========================================================================================================")
	fmt.Println(judulKolom)
	garisBawah("==========================================================================================================")
}

func pilihFakultas(config *Konfigurasi) {
	var pilih int

	judulProgram("PEMILIHAN ORGANISASI MAHASISWA TELKOM UNIVERSITY")

	fmt.Println("1. Fakultas Teknik Elektro")
	fmt.Println("2. Fakultas Rekayasa Industri")
	fmt.Println("3. Fakultas Informatika")
	fmt.Println("4. Fakultas Ekonomi dan Bisnis")
	fmt.Println("5. Fakultas Komunikasi dan Ilmu Sosial")
	fmt.Println("6. Fakultas Industri Kreatif")
	fmt.Println("7. Fakultas Ilmu Terapan")

	fmt.Print("Pilih Fakultas : ")
	fmt.Scan(&pilih)

	switch pilih {

	case 1:
		config.fakultas = "Fakultas Teknik Elektro"

	case 2:
		config.fakultas = "Fakultas Rekayasa Industri"

	case 3:
		config.fakultas = "Fakultas Informatika"

	case 4:
		config.fakultas = "Fakultas Ekonomi dan Bisnis"

	case 5:
		config.fakultas = "Fakultas Komunikasi dan Ilmu Sosial"

	case 6:
		config.fakultas = "Fakultas Industri Kreatif"

	case 7:
		config.fakultas = "Fakultas Ilmu Terapan"

	default:
		config.fakultas = "Fakultas Informatika"
	}
}

// SEARCHING
func sequentialSearchNama(data DataKandidat, nama string) int {
	var i int

	for i = 0; i < data.n; i++ {
		if data.data[i].nama == nama {
			return i
		}
	}
	return -1
}

func sequentialSearchNoUrut(data DataKandidat, noUrut int) int {
	var i int

	for i = 0; i < data.n; i++ {
		if data.data[i].noUrut == noUrut {
			return i
		}
	}
	return -1
}

func binarySearchNama(data DataKandidat, nama string) int {
	var left, right, mid int

	left = 0
	right = data.n - 1

	for left <= right {
		mid = (left + right) / 2
		if data.data[mid].nama == nama {
			return mid
		} else if data.data[mid].nama < nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// SORTING
func insertionSortNamaAscending(data *DataKandidat) {
	var i, j int
	var temp Kandidat

	for i = 1; i < data.n; i++ {
		temp = data.data[i]
		j = i
		for j > 0 && temp.nama < data.data[j-1].nama {
			data.data[j] = data.data[j-1]
			j--
		}
		data.data[j] = temp
	}
}

func insertionSortNoUrutAscending(data *DataKandidat) {
	var i, j int
	var temp Kandidat

	for i = 1; i < data.n; i++ {
		temp = data.data[i]
		j = i
		for j > 0 && temp.noUrut < data.data[j-1].noUrut {
			data.data[j] = data.data[j-1]
			j--
		}
		data.data[j] = temp
	}
}

func selectionSortSuaraDescending(data *DataKandidat) {
	var i, j, maxIdx int
	var temp Kandidat

	for i = 0; i < data.n-1; i++ {
		maxIdx = i
		for j = i + 1; j < data.n; j++ {
			if data.data[j].suara > data.data[maxIdx].suara {
				maxIdx = j
			}
		}
		temp = data.data[i]
		data.data[i] = data.data[maxIdx]
		data.data[maxIdx] = temp
	}
}

// NILAI EKSTRIM
func cariPemenang(data DataKandidat) Kandidat {
	var i, maxIdx int

	maxIdx = 0

	for i = 1; i < data.n; i++ {
		if data.data[i].suara > data.data[maxIdx].suara {
			maxIdx = i
		}
	}
	return data.data[maxIdx]
}

func cariTerendah(data DataKandidat) Kandidat {
	var i, minIdx int

	minIdx = 0

	for i = 1; i < data.n; i++ {
		if data.data[i].suara < data.data[minIdx].suara {
			minIdx = i
		}
	}
	return data.data[minIdx]
}

func hitungTotalSuara(data DataKandidat) int {
	var i, total int

	total = 0

	for i = 0; i < data.n; i++ {
		total += data.data[i].suara
	}
	return total
}

// PROCEDURE ADMIN
func tambahKandidat(data *DataKandidat, config Konfigurasi, pesan string) {
	var jumlah, i int
	var nama string
	var idx int

	fmt.Println(pesan)

	if data.n == NMAX {

		fmt.Println(">> Data kandidat sudah penuh.")

	} else {

		fmt.Print("Masukkan jumlah kandidat : ")
		fmt.Scan(&jumlah)

		if jumlah < 1 || jumlah > 10 - data.n {

			fmt.Println(">> Jumlah kandidat maksimal 10.")

		} else {

			fmt.Println("Gunakan underscore (_) jika nama mengandung spasi")

			for i = 0; i < jumlah; i++ {

				fmt.Printf("Masukkan nama kandidat ke-%d : ", i+1)
				fmt.Scan(&nama)

				idx = sequentialSearchNama(*data, nama)

				if idx == -1 {

					data.data[data.n].noUrut = data.n + 1
					data.data[data.n].nama = nama

					fmt.Print("Masukkan program studi : ")
					fmt.Scan(&data.data[data.n].prodi)

					fmt.Print("Masukkan angkatan (2022-2024) : ")
					fmt.Scan(&data.data[data.n].angkatan)

					if data.data[data.n].angkatan < 2022 ||
						data.data[data.n].angkatan > 2024 {

						fmt.Println(">> Angkatan harus antara 2022 - 2024")
						i--
						continue
					}

					data.data[data.n].fakultas = config.fakultas

					data.data[data.n].suara = 0

					data.n++

				} else {

					fmt.Println(">> Kandidat sudah ada.")
					i--
				}
			}

			fmt.Println(">> Kandidat berhasil ditambahkan.")
		}
	}
}

func hapusKandidat(data *DataKandidat, pesan string) {
	var nama string
	var idx, i int
	
	fmt.Println(pesan)

	fmt.Print("Masukkan nama kandidat yang dihapus : ")
	fmt.Scan(&nama)

	idx = sequentialSearchNama(*data, nama)

	if idx != -1 {
		for i = idx; i < data.n-1; i++ {
			data.data[i] = data.data[i+1]
		}
		data.n--
		for i = 0; i < data.n; i++ {
			data.data[i].noUrut = i + 1
		}
		fmt.Println(">> Kandidat berhasil dihapus.")
	} else {
		fmt.Println(">> Kandidat tidak ditemukan.")
	}
}

func editKandidat(data *DataKandidat, pesan string) {
	var namaLama, namaBaru, prodiBaru string
	var idx int
	var angkatanBaru int 
	

	fmt.Println(pesan)
	fmt.Println("Gunakan underscore (_) jika nama mengandung spasi")

	fmt.Print("Masukkan nama kandidat lama : ")
	fmt.Scan(&namaLama)

	idx = sequentialSearchNama(*data, namaLama)

	if idx != -1 {
		fmt.Print("Masukkan nama kandidat baru : ")
		fmt.Scan(&namaBaru)
		
		fmt.Print("Masukkan program studi kandidat baru : ")
		fmt.Scan(&prodiBaru)
		
		fmt.Print("Masukkan angkatan kandidat baru : ")
		fmt.Scan(&angkatanBaru)
		
		if angkatanBaru >= 2022 && angkatanBaru <= 2024 {
			data.data[idx].nama = namaBaru
			data.data[idx].prodi = prodiBaru
			data.data[idx].angkatan = angkatanBaru
			fmt.Println(">> Kandidat berhasil diperbarui.")
		} else {
			fmt.Println("Kandidat harus angkatan 2022-2024")
		}

	} else {
		fmt.Println(">> Kandidat tidak ditemukan.")
	}
}

func resetVoting(data *DataKandidat, pesan string) {
	var i int

	fmt.Println(pesan)
	for i = 0; i < data.n; i++ {
		data.data[i].suara = 0
	}
	fmt.Println(">> Semua voting berhasil direset.")
}

//publik yg voting
func votingKandidat(data *DataKandidat, pesan string) {
	var noUrut int
	var idx int

	fmt.Println(pesan)

	if data.n == 0 {
		fmt.Println(">> Belum ada kandidat.")
	} else {
		tampilkanKandidat(*data, "DAFTAR KANDIDAT")

		fmt.Print("Masukkan nomor urut kandidat pilihan : ")
		fmt.Scan(&noUrut)

		idx = sequentialSearchNoUrut(*data, noUrut)
		if idx != -1 {
			data.data[idx].suara++
			fmt.Println(">> Voting berhasil.")
		} else {
			fmt.Println(">> Kandidat tidak ditemukan.")
		}
	}
}

// TAMPILKAN DATA
func tampilkanKandidat(data DataKandidat, judul string) {
	var i int
	var temp DataKandidat
	
	temp = data
	fmt.Println(judul)
	insertionSortNoUrutAscending(&temp)
	if data.n == 0 {
		fmt.Println(">> Belum ada kandidat.")
	} else {
		tampilkanHeaderKandidat("No | No Urut | Nama Kandidat        | Prodi                     | Angkatan   ")
		for i = 0; i < data.n; i++ {
			fmt.Printf("%-2d | %-7d | %-20s | %-25s | %-10d \n", i+1, temp.data[i].noUrut, temp.data[i].nama, temp.data[i].prodi, temp.data[i].angkatan)
		}
		garisBawah("==========================================================================================================")
	}
}


func cariDetailKandidat(data *DataKandidat, pesan string) {
	var nama string
	var idx int
	fmt.Println(pesan)
	if data.n == 0 {
		fmt.Println(">> Belum ada kandidat.")
	} else {
		insertionSortNamaAscending(data)

		fmt.Println("Gunakan underscore (_) jika nama mengandung spasi")

		fmt.Print("Masukkan nama kandidat : ")
		fmt.Scan(&nama)

		idx = binarySearchNama(*data, nama)
		if idx != -1 {
			garisAtas("==========================================================================================================")

			fmt.Println("DETAIL KANDIDAT")

			garisBawah("==========================================================================================================")

			fmt.Println("Nama Kandidat  :", data.data[idx].nama)
			fmt.Println("No Urut        :", data.data[idx].noUrut)
			fmt.Println("Program Studi  :", data.data[idx].prodi)
			fmt.Println("Fakultas       :", data.data[idx].fakultas)
			fmt.Println("Angkatan       :", data.data[idx].angkatan)
			fmt.Println("Jumlah Suara   :", data.data[idx].suara)
		} else {
			fmt.Println(">> Kandidat tidak ditemukan.")
		}
	}
}

func tampilkanHasilVoting(data *DataKandidat, judul string) {
	var i int
	fmt.Println(judul)
	if data.n == 0 {
		fmt.Println(">> Belum ada kandidat.")
	} else {
		selectionSortSuaraDescending(data)
		tampilkanHeaderKandidat("No | No Urut | Nama Kandidat        | Prodi                     | Angkatan   | Jumlah Suara")
		for i = 0; i < data.n; i++ {
			fmt.Printf("%-2d | %-7d | %-20s | %-25s | %-10d | %d\n", i+1, data.data[i].noUrut, data.data[i].nama, data.data[i].prodi, data.data[i].angkatan, data.data[i].suara)
		}
	}
}

func tampilkanPemenang(data DataKandidat, judul string) {
	var i int
	var total int
	var maxSuara int
	var jumlahPemenang int
	var semuaNol bool

	fmt.Println(judul)

	if data.n == 0 {
		fmt.Println("Belum ada kandidat")
		return
	}
	total = hitungTotalSuara(data)
	if total == 0 {
		fmt.Println("Voting belum dimulai")
		return
	}

	// cari suara tertinggi
	maxSuara = data.data[0].suara
	for i = 1; i < data.n; i++ {
		if data.data[i].suara > maxSuara {
			maxSuara = data.data[i].suara
		}
	}

	// hitung berapa kandidat dengan suara tertinggi
	for i = 0; i < data.n; i++ {
		if data.data[i].suara == maxSuara {
			jumlahPemenang++
		}
		if data.data[i].suara > 0 {
			semuaNol = true
		}
	}

	if !semuaNol {
		fmt.Println("Voting belum dimulai")
		return
	}
	fmt.Println()

	// hanya satu pemenang
	if jumlahPemenang == 1 {
		fmt.Println("=== PEMENANG VOTING ===")
		for i = 0; i < data.n; i++ {
			if data.data[i].suara == maxSuara {
				fmt.Println("Nama Kandidat :", data.data[i].nama)
				fmt.Println("No Urut       :", data.data[i].noUrut)
				fmt.Println("Jumlah Suara  :", data.data[i].suara)
			}
		}
	} else {
		// seri
		fmt.Println("=== BELUM ADA PEMENANG ===")
		fmt.Printf(
			"Terdapat %d kandidat dengan suara tertinggi\n", jumlahPemenang)
		fmt.Println()
		fmt.Println("Kandidat Terunggul:")

		for i = 0; i < data.n; i++ {
			if data.data[i].suara == maxSuara {
				fmt.Printf(
					"- %s (No.%d) : %d suara\n",
					data.data[i].nama,
					data.data[i].noUrut,
					data.data[i].suara,
				)
			}
		}
	}

}

func tampilkanStatistik(data DataKandidat, judul string) {
	var i, total, maxSuara, minSuara, jumlahMax, jumlahMin int
	var rata, persen float64
	var semuaNol bool
	
	fmt.Println(judul)
	
	if data.n == 0 {
		fmt.Println(">> Belum ada kandidat.")
		return
	}

	total = hitungTotalSuara(data)
	fmt.Println("Jumlah Kandidat   :", data.n)
	fmt.Println("Total Suara       :", total)

	// cek voting belum dimulai
	semuaNol = true

	for i = 0; i < data.n; i++ {
		if data.data[i].suara > 0 {
			semuaNol = false
		}
	}

	if semuaNol {
		fmt.Println()
		fmt.Println("Status Voting :")
		fmt.Println("Voting belum dimulai.")
		return
	}

	rata = float64(total) / float64(data.n)
	fmt.Printf("Rata-rata Suara : %.2f\n", rata)

	maxSuara = data.data[0].suara
	minSuara = data.data[0].suara

	// cari max min
	for i = 1; i < data.n; i++ {
		if data.data[i].suara > maxSuara {
			maxSuara = data.data[i].suara
		}
		if data.data[i].suara < minSuara {
			minSuara = data.data[i].suara
		}
	}
	fmt.Println()

	// SUARA TERTINGGI
	fmt.Println("Suara Tertinggi :")

	for i = 0; i < data.n; i++ {
		if data.data[i].suara == maxSuara {
			jumlahMax++
			fmt.Printf("- %s (%d suara)\n",
				data.data[i].nama,
				data.data[i].suara)
		}
	}
	fmt.Println()

	// SUARA TERENDAH
	fmt.Println("Suara Terendah :")
	for i = 0; i < data.n; i++ {
		if data.data[i].suara == minSuara {
			jumlahMin++
			fmt.Printf("- %s (%d suara)\n",
				data.data[i].nama,
				data.data[i].suara)
		}
	}
	fmt.Println()

	// RENTANG
	fmt.Println("Rentang Suara :")
	fmt.Printf("%d - %d\n", minSuara, maxSuara)

	fmt.Printf("(Selisih : %d suara)\n",
		maxSuara-minSuara)

	fmt.Println()

	// DISTRIBUSI
	fmt.Println("Distribusi Suara :")

	for i = 0; i < data.n; i++ {
		persen =(float64(data.data[i].suara) / float64(total)) * 100
		fmt.Printf("%s : %d suara (%.2f%%)\n", data.data[i].nama, data.data[i].suara, persen)
	}
	fmt.Println()

	// PERSENTASE PEMENANG
	fmt.Println("Persentase Pemenang :")
	for i = 0; i < data.n; i++ {
		if data.data[i].suara == maxSuara {
			persen = (float64(data.data[i].suara) / float64(total)) * 100
			fmt.Printf("- %s : %.2f%%\n", data.data[i].nama, persen)
		}
	}
	fmt.Println()

	// STATUS
	fmt.Println("Status Voting :")
	if jumlahMax > 1 {
		fmt.Println("Terdapat lebih dari satu kandidat dengan suara tertinggi (seri)")
	} else {
		fmt.Println("Terdapat satu kandidat unggul.")
	}
	fmt.Println()

	// KESIMPULAN
	fmt.Println("Kesimpulan :")
	if jumlahMax > 1 {
		for i = 0; i < data.n; i++ {
			if data.data[i].suara == maxSuara {
				fmt.Print(data.data[i].nama, " ")
			}
		}
		fmt.Printf("memimpin dengan %d suara\n", maxSuara)
	} else {
		for i = 0; i < data.n; i++ {
			if data.data[i].suara == maxSuara {
				fmt.Printf("%s menjadi kandidat dengan perolehan suara terbanyak (%d suara)\n", data.data[i].nama, maxSuara)
			}
		}
	}
}

// LOGIN
func loginAdmin(usernameBenar string, passwordBenar string) bool {
	var username, password string

	fmt.Print("Username : ")
	fmt.Scan(&username)

	fmt.Print("Password : ")
	fmt.Scan(&password)
	if username == usernameBenar && password == passwordBenar {
		fmt.Println(">> Login berhasil.")
		return true
	}
	fmt.Println(">> Login gagal.")
	return false
}

// MENU
func menuAdmin(data *DataKandidat, config Konfigurasi, judul string) {
	var pilihan int
	var logout bool

	for !logout {

		judulProgram(judul)

		fmt.Println("1. Tambah Kandidat")
		fmt.Println("2. Hapus Kandidat")
		fmt.Println("3. Edit Kandidat")
		fmt.Println("4. Tampilkan Kandidat")
		fmt.Println("5. Cari Kandidat")
		fmt.Println("6. Hasil Voting")
		fmt.Println("7. Statistik Voting")
		fmt.Println("8. Tampilkan Pemenang")
		fmt.Println("9. Reset Voting")
		fmt.Println("10. Logout")

		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahKandidat(data, config, "========================== TAMBAH KANDIDAT ==========================")

		case 2:
			hapusKandidat(data, "========================== HAPUS KANDIDAT ==========================")

		case 3:
			editKandidat(data, "========================== EDIT KANDIDAT ==========================")

		case 4:
			tampilkanKandidat(*data, "========================== DAFTAR KANDIDAT ==========================")

		case 5:
			cariDetailKandidat(data, "========================== CARI KANDIDAT ==========================")

		case 6:
			tampilkanHasilVoting(data, "========================== HASIL VOTING ==========================")

		case 7:
			tampilkanStatistik(*data, "========================== STATISTIK VOTING ==========================")

		case 8:
			tampilkanPemenang(*data, "=============================== PEMENANG ===============================")

		case 9:
			resetVoting(data, "========================== RESET VOTING ==========================")

		case 10:
			logout = true

		default:
			fmt.Println(">> Pilihan tidak valid.")
		}
	}
}

func menuPublik(data *DataKandidat, judul string) {
	var pilihan int
	var keluar bool

	for !keluar {

		judulProgram(judul)

		fmt.Println("1. Lihat Kandidat")
		fmt.Println("2. Voting")
		fmt.Println("3. Cari Kandidat")
		fmt.Println("4. Hasil Voting")
		fmt.Println("5. Pemenang")
		fmt.Println("6. Kembali")

		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanKandidat(*data, "===== DAFTAR KANDIDAT =====")

		case 2:
			votingKandidat(data, "===== MENU VOTING =====")

		case 3:
			cariDetailKandidat(data, "===== CARI KANDIDAT =====")

		case 4:
			tampilkanHasilVoting(data, "===== HASIL VOTING =====")

		case 5:
			tampilkanPemenang(*data, "===== PEMENANG =====")

		case 6:
			keluar = true

		default:
			fmt.Println(">> Pilihan tidak valid.")
		}
	}
}

func menuUtama(data *DataKandidat, config Konfigurasi, judul string) {
	var pilihan int
	var selesai bool
	var berhasil bool

	for !selesai {

		judulProgram(judul)

		fmt.Println("1. Login Admin")
		fmt.Println("2. Masuk Sebagai Publik")
		fmt.Println("3. Keluar")

		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			berhasil = loginAdmin(ADMIN_USERNAME, ADMIN_PASSWORD)

			if berhasil {
				menuAdmin(data, config, judul)
			}

		case 2:
			menuPublik(data, judul)

		case 3:
			selesai = true
			fmt.Println(">> Program selesai.")

		default:
			fmt.Println(">> Pilihan tidak valid.")
		}
	}
}

// MAIN
func main() {
	var data DataKandidat
	var config Konfigurasi

	pilihFakultas(&config)

	switch config.fakultas {

	case "Fakultas Teknik Elektro":

		data.n = 2
		data.data[0] = Kandidat{
			noUrut:   1,
			nama:     "Naruto_Uzumaki",
			prodi:    "Teknik_Elektro",
			fakultas: config.fakultas,
			angkatan: 2023,
			suara:    0,
		}

		data.data[1] = Kandidat{
			noUrut:   2,
			nama:     "Sakura_Haruno",
			prodi:    "Teknik_Biomedis",
			fakultas: config.fakultas,
			angkatan: 2022,
			suara:    0,
		}

	case "Fakultas Rekayasa Industri":

		data.n = 2
		data.data[0] = Kandidat{
			noUrut:   1,
			nama:     "Armin_Arlert",
			prodi:    "Teknik_Industri",
			fakultas: config.fakultas,
			angkatan: 2024,
			suara:    0,
		}

		data.data[1] = Kandidat{
			noUrut:   2,
			nama:     "Mikasa_Ackerman",
			prodi:    "Sistem_Informasi",
			fakultas: config.fakultas,
			angkatan: 2023,
			suara:    0,
		}

	case "Fakultas Informatika":

		data.n = 2
		data.data[0] = Kandidat{
			noUrut:   1,
			nama:     "Eren_Yeager",
			prodi:    "Teknologi_Informasi",
			fakultas: config.fakultas,
			angkatan: 2023,
			suara:    0,
		}

		data.data[1] = Kandidat{
			noUrut:   2,
			nama:     "Jean_Kirstein",
			prodi:    "Informatika",
			fakultas: config.fakultas,
			angkatan: 2022,
			suara:    0,
		}

	case "Fakultas Ekonomi dan Bisnis":

		data.n = 2
		data.data[0] = Kandidat{
			noUrut:   1,
			nama:     "Sasha_Blouse",
			prodi:    "Manajemen",
			fakultas: config.fakultas,
			angkatan: 2024,
			suara:    0,
		}

		data.data[1] = Kandidat{
			noUrut:   2,
			nama:     "Connie_Springer",
			prodi:    "Akuntansi",
			fakultas: config.fakultas,
			angkatan: 2022,
			suara:    0,
		}

	case "Fakultas Komunikasi dan Ilmu Sosial":

		data.n = 2
		data.data[0] = Kandidat{
			noUrut:   1,
			nama:     "Nabila_Riska",
			prodi:    "Ilmu_Komunikasi",
			fakultas: config.fakultas,
			angkatan: 2024,
			suara:    0,
		}

		data.data[1] = Kandidat{
			noUrut:   2,
			nama:     "Indah_Wulansari",
			prodi:    "Public_Relations",
			fakultas: config.fakultas,
			angkatan: 2023,
			suara:    0,
		}

	case "Fakultas Industri Kreatif":

		data.n = 2
		data.data[0] = Kandidat{
			noUrut:   1,
			nama:     "Aliya_Mahisa",
			prodi:    "Seni_Rupa",
			fakultas: config.fakultas,
			angkatan: 2022,
			suara:    0,
		}

		data.data[1] = Kandidat{
			noUrut:   2,
			nama:     "Rava_Amesta",
			prodi:    "Film_dan_Animasi",
			fakultas: config.fakultas,
			angkatan: 2024,
			suara:    0,
		}

	case "Fakultas Ilmu Terapan":

		data.n = 2
		data.data[0] = Kandidat{
			noUrut:   1,
			nama:     "Dimas_Satyabhakti",
			prodi:    "Teknologi_Komputer",
			fakultas: config.fakultas,
			angkatan: 2023,
			suara:    0,
		}
		
		data.data[1] = Kandidat{
			noUrut:   2,
			nama:     "Sasuke_Uchiha",
			prodi:    "Teknik_Telekomunikasi",
			fakultas: config.fakultas,
			angkatan: 2022,
			suara:    0,
		}
	}
	menuUtama(&data, config, "APLIKASI VOTING PEMILIHAN KETUA BEM "+config.fakultas)
}