package main
import "fmt"

const NMAX int = 10
const ADMIN_USERNAME string = "admin"
const ADMIN_PASSWORD string = "12345"

type Kandidat struct {
	noUrut int
	nama   string
	suara  int
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
	garisAtas("========================================================================================")
	fmt.Println(judul)
	garisBawah("========================================================================================")
}

func tampilkanHeaderKandidat(judulKolom string) {
	garisAtas("========================================================================================")
	fmt.Println(judulKolom)
	garisBawah("========================================================================================")
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
func tambahKandidat(data *DataKandidat, pesan string) {
	var jumlah, i int
	var nama string
	var idx int

	fmt.Println(pesan)

	if data.n == NMAX {
		fmt.Println(">> Data kandidat sudah penuh.")
	} else {
		fmt.Print("Masukkan jumlah kandidat : ")
		fmt.Scan(&jumlah)
		if jumlah < 1 || jumlah > 10 {
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
	var namaLama, namaBaru string
	var idx int

	fmt.Println(pesan)
	fmt.Println("Gunakan underscore (_) jika nama mengandung spasi")

	fmt.Print("Masukkan nama kandidat lama : ")
	fmt.Scan(&namaLama)

	idx = sequentialSearchNama(*data, namaLama)

	if idx != -1 {
		fmt.Print("Masukkan nama kandidat baru : ")
		fmt.Scan(&namaBaru)

		data.data[idx].nama = namaBaru

		fmt.Println(">> Kandidat berhasil diperbarui.")
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

	fmt.Println(judul)
	if data.n == 0 {
		fmt.Println(">> Belum ada kandidat.")
	} else {
		tampilkanHeaderKandidat("No | No Urut | Nama Kandidat             | Jumlah Suara")
		for i = 0; i < data.n; i++ {
			fmt.Printf("%-2d | %-7d | %-25s | %d suara\n",
				i+1,
				data.data[i].noUrut,
				data.data[i].nama,
				data.data[i].suara)
		}
		garisBawah("========================================================================================")
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
			garisAtas("========================================================================================")

			fmt.Println("DETAIL KANDIDAT")

			garisBawah("========================================================================================")

			fmt.Println("Nama Kandidat:", data.data[idx].nama)
			fmt.Println("No Urut:", data.data[idx].noUrut)
			fmt.Println("Jumlah Suara:", data.data[idx].suara)
		} else {
			fmt.Println(">> Kandidat tidak ditemukan.")
		}
	}
}

func tampilkanHasilVoting(data *DataKandidat, judul string) {
	fmt.Println(judul)
	if data.n == 0 {
		fmt.Println(">> Belum ada kandidat.")
	} else {
		selectionSortSuaraDescending(data)

		tampilkanKandidat(*data, "HASIL VOTING")
	}
}

func tampilkanPemenang(data DataKandidat, judul string) {
	var total int
	var pemenang Kandidat

	fmt.Println(judul)
	if data.n == 0 {
		fmt.Println("Belum ada kandidat")
	} else {
		total = hitungTotalSuara(data)
		if total == 0 {
			fmt.Println("Voting belum dimulai")
		} else {
			pemenang = cariPemenang(data)

			fmt.Println("Nama Kandidat:", pemenang.nama)
			fmt.Println("No Urut:", pemenang.noUrut)
			fmt.Println("Jumlah Suara:", pemenang.suara)
		}
	}
}

func tampilkanStatistik(data DataKandidat, judul string) {
	var total int
	var tertinggi, terendah Kandidat

	fmt.Println(judul)

	if data.n == 0 {
		fmt.Println(">> Belum ada kandidat.")
	} else {
		total = hitungTotalSuara(data)

		tertinggi = cariPemenang(data)
		terendah = cariTerendah(data)

		fmt.Println("Total Suara      :", total)
		fmt.Println("Suara Tertinggi  :", tertinggi.nama)
		fmt.Println("Suara Terendah   :", terendah.nama)
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
func menuAdmin(data *DataKandidat, judul string) {
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
			tambahKandidat(data, "===== TAMBAH KANDIDAT =====")

		case 2:
			hapusKandidat(data, "===== HAPUS KANDIDAT =====")

		case 3:
			editKandidat(data, "===== EDIT KANDIDAT =====")

		case 4:
			tampilkanKandidat(*data, "===== DAFTAR KANDIDAT =====")

		case 5:
			cariDetailKandidat(data, "===== CARI KANDIDAT =====")

		case 6:
			tampilkanHasilVoting(data, "===== HASIL VOTING =====")

		case 7:
			tampilkanStatistik(*data, "===== STATISTIK =====")

		case 8:
			tampilkanPemenang(*data, "===== PEMENANG =====")

		case 9:
			resetVoting(data, "===== RESET VOTING =====")

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

func menuUtama(data *DataKandidat, judul string) {
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
				menuAdmin(data, "APLIKASI VOTING PEMILIHAN KETUA BEM")
			}

		case 2:
			menuPublik(data, "APLIKASI VOTING PEMILIHAN KETUA BEM")

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

	data.n = 3

	data.data[0].noUrut = 1
	data.data[0].nama = "rava_amesta"
	data.data[0].suara = 0

	data.data[1].noUrut = 2
	data.data[1].nama = "dimas_satyabhakti"
	data.data[1].suara = 0

	data.data[2].noUrut = 3
	data.data[2].nama = "eren_yeager"
	data.data[2].suara = 0

	menuUtama(&data, "APLIKASI VOTING PEMILIHAN KETUA BEM")
}