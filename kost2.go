package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// tampilan
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printCentered(text string, width int) {
	padding := (width - len(text)) / 2
	fmt.Print(strings.Repeat(" ", padding))
	fmt.Print(text)
	fmt.Println()
}

func setGreenText() {
	fmt.Print("\033[32m")
}

func resetTextColor() {
	fmt.Print("\033[0m")
}

const NMAX = 100

type Penghuni struct {
	Nama           string
	NomorKamar     int
	NomorIdentitas int
	NomorTelepon   string
	Alamat         string
	TanggalMasuk   string
	JenisKamar     string
	LamaSewa       int
	Pembayaran     int
	Total          int
	Lunas          bool
}

type Kamar struct {
	NomorKamar int
	Tipe       string
	Harga      float64
	Fasilitas  [NMAX]string
	Status     string
}

var penghuni [NMAX]Penghuni
var kamar [NMAX]Kamar

var numPenghuni int
var numKamar int

func main() {
	menu()
}

func menu() {
	var tampilkandataKamar bool
	var tampilkanpenghuni bool
	var tampilkanpembayaran bool
	for {
		if !tampilkandataKamar && !tampilkanpenghuni && !tampilkanpembayaran{
            clearScreen()
            setGreenText()
		
		// width := 80

		// fmt.Println(strings.Repeat("═", width))
		fmt.Printf("%25s╔═════════════════════════════════════════════════════════════════════════╗\n", "")
		fmt.Printf("%25s║%73s║\n", "", "")
		fmt.Printf("%25s║%53s%-20s║\n", "", "Aplikasi Pengelolaan Kost", "")
		fmt.Printf("%25s║%73s║\n", "", "")
		fmt.Printf("%25s║%55s %17s║\n", "", "Created by : • Relingga Aditya ", "")
		fmt.Printf("%25s║%56s %16s║\n", "", "• Nadia Aulia Salma", "")
		fmt.Printf("%25s║%56s %16s║\n", "", "• Deswita Syaharani", "")
		fmt.Printf("%25s║%73s║\n", "", "")
		fmt.Printf("%25s╠═════════════════════════════════════════════════════════════════════════╣\n", "")
		

		
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "1.Input Data Kamar")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "2.Pendataan Penghuni")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "3.Pencarian Data Kamar Dan Ketersedian Kamar")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "4.Edit data kamar")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "5.Edit data penghuni")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "6.Hapus data kamar")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "8.Tampilkan data penghuni")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "9.Tampilkan data kamar")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "10.Pembayaran penghuni")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "11. Keluar")
		fmt.Printf("%25s║%61s╔═══════════╣\n", "", "")
		fmt.Printf("%25s║%73s║\n", "", "║ Menu Utama")
		fmt.Printf("%25s╚═════════════════════════════════════════════════════════════════════════╝\n", "")
		fmt.Printf("%25sPilih Menu > ", "")
		}
		var pilihan int	
		fmt.Scan(&pilihan)
			
		switch pilihan {
		case 1:
			inputDataKamar()
		case 2:
			pendataanPenghuni()
		case 3:
			pencarianDataKamar()
		case 4:
			editDataKamar()
		case 5:
			editDataPenghuni()
		case 6:
			hapusDataKamar()
		case 7:
			hapusDataPenghuni()
		case 8:
			tampilkanDataPenghuni()
			tampilkanpenghuni=true
		case 9:
			tampilkanDataKamar()
			tampilkandataKamar = true 
		case 10:
			pembayaran()
			tampilkanpembayaran=true
		case 11:
			fmt.Println("Terima kasih. Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func pendataanPenghuni() {
	clearScreen()
	setGreenText()
	width := 80
	fmt.Println(strings.Repeat("═", width))
	printCentered("Pendataan Penghuni", width)
	fmt.Println(strings.Repeat("═", width))
	resetTextColor()

	var penghuniBaru Penghuni
	fmt.Print("Masukkan nama penghuni:")
	fmt.Scan(&penghuniBaru.Nama)
	fmt.Print("Masukkan nomor kamar penghuni:")
	fmt.Scan(&penghuniBaru.NomorKamar)
	if !cekKetersediaanKamar(penghuniBaru.NomorKamar) {
		fmt.Println("Kamar tidak tersedia. Silakan pilih kamar lain.")
		return
	}
	fmt.Print("Masukkan nomor identitas penghuni:")
	fmt.Scan(&penghuniBaru.NomorIdentitas)
	fmt.Print("Masukkan nomor telepon penghuni:")
	fmt.Scan(&penghuniBaru.NomorTelepon)
	fmt.Print("Masukkan alamat penghuni:")
	fmt.Scan(&penghuniBaru.Alamat)
	fmt.Print("Masukkan tanggal masuk penghuni (YYYY-MM-DD):")
	fmt.Scan(&penghuniBaru.TanggalMasuk)
	fmt.Print("Masukkan jenis kamar penghuni:")
	fmt.Scan(&penghuniBaru.JenisKamar)
	fmt.Print("Durasi sewa:")
	fmt.Scan(&penghuniBaru.LamaSewa)
	fmt.Print("Jumlah uang yang dibayarkan:")
	fmt.Scan(&penghuniBaru.Pembayaran)

	// Find the price of the room
	var hargakamar float64
	for i := 0; i < numKamar; i++ {
		if kamar[i].NomorKamar == penghuniBaru.NomorKamar {
			hargakamar = kamar[i].Harga
			kamar[i].Status =penghuniBaru.Nama
			break
		}
	}

	penghuniBaru.Total = int(hargakamar) * penghuniBaru.LamaSewa
	penghuniBaru.Lunas = false

	penghuni[numPenghuni] = penghuniBaru
	numPenghuni++
	fmt.Println("Penghuni berhasil didata.")
}

func inputDataKamar() {
	clearScreen()
	setGreenText()
	width := 80
	fmt.Println(strings.Repeat("═", width))
	printCentered("Input Data Kamar", width)
	fmt.Println(strings.Repeat("═", width))

	var kamarBaru Kamar
	var inputVasilitas string
	var nf int
	var i int

	fmt.Print("Masukkan nomor kamar: ")
	fmt.Scan(&kamarBaru.NomorKamar)

	// Periksa apakah nomor kamar sudah ada
	for j := 0; j < numKamar; j++ {
		if kamar[j].NomorKamar == kamarBaru.NomorKamar {
			fmt.Println("Maaf, nomor kamar sudah ada.")
			return
		}
	}

	fmt.Print("Masukkan tipe kamar: ")
	fmt.Scan(&kamarBaru.Tipe)
	fmt.Print("Masukkan harga sewa per bulan: ")
	fmt.Scan(&kamarBaru.Harga)
	fmt.Print("Masukkan fasilitas kamar (pisahkan dengan spasi, ketik '.' untuk selesai): ")

	i = 0
	fmt.Scan(&inputVasilitas)

	for inputVasilitas != "." {
		kamarBaru.Fasilitas[i] = inputVasilitas
		i++
		nf++
		fmt.Scan(&inputVasilitas)
	}

	kamarBaru.Status = "tersedia"

	kamar[numKamar] = kamarBaru
	numKamar++
	fmt.Println("Data kamar berhasil dimasukkan.")
}

func pencarianDataKamar() {
	clearScreen()
	setGreenText()
	width := 80
	fmt.Println(strings.Repeat("═", width))
	printCentered("Pencarian Data Kamar", width)
	fmt.Println(strings.Repeat("═", width))
		
	var nomorKamarDicari int
	ketemu := false

	fmt.Println("Masukkan nomor kamar yang ingin dicari:")
	fmt.Scan(&nomorKamarDicari)

	for i := 0; i < numKamar; i++ {
		if kamar[i].NomorKamar == nomorKamarDicari {
			ketemu = true
			fmt.Printf("Kamar dengan nomor %d ditemukan.\n", nomorKamarDicari)
			fmt.Printf("Tipe: %s, Harga: %.2f, Status: %s\n", kamar[i].Tipe, kamar[i].Harga, kamar[i].Status)
			fmt.Print("Fasilitas: ")
			for j := 0; j < 5; j++ {
				if kamar[i].Fasilitas[j] != "" {
					fmt.Printf("%s ", kamar[i].Fasilitas[j])
				}
			}
			fmt.Println()
			break
		}
	}

	if !ketemu {
		fmt.Printf("Kamar dengan nomor %d tidak ditemukan.\n", nomorKamarDicari)
	}
}

func cekKetersediaanKamar(no int) bool {
	for i := 0; i < numKamar; i++ {
		if kamar[i].NomorKamar == no && kamar[i].Status == penghuni[i].Nama {
			return false
		}
	}
	return true
}

func pembayaran() {
	clearScreen()
	setGreenText()
	width := 80
	fmt.Println(strings.Repeat("═", width))
	printCentered("Pembayaran Penghuni", width)
	fmt.Println(strings.Repeat("═", width))

	var nama string
	var bayar int
	var found bool = false
	fmt.Print("Nama penghuni yang akan melakukan pembayaran: ")
	fmt.Scan(&nama)

	for i := 0; i < numPenghuni; i++ {
		if penghuni[i].Nama == nama {
			fmt.Println("Sisa uang yang belum dibayarkan:", penghuni[i].Total-penghuni[i].Pembayaran)
			fmt.Print("Jumlah uang yang dibayarkan: ")
			fmt.Scan(&bayar)
			penghuni[i].Pembayaran += bayar
			found = true
			if penghuni[i].Pembayaran >= penghuni[i].Total {
				penghuni[i].Lunas = true
				fmt.Println("Pembayaran sudah lunas")
				for j := i; j < numPenghuni-1; j++ {
					penghuni[j] = penghuni[j+1]
				}
				numPenghuni--
			} else {
				fmt.Println("Sisa uang yang belum dibayarkan:", penghuni[i].Total-penghuni[i].Pembayaran)
			}
			break
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan penghuni dengan nama tersebut.")
	}
	var keluar string
    fmt.Println("\nTekan 1 lalu enter untuk kembali ke menu utama...")
    fmt.Scan(&keluar)
	if keluar=="1"{
		menu()
	}
	
}


func editDataKamar() {
	clearScreen()
	setGreenText()
	width := 80
	fmt.Println(strings.Repeat("═", width))
	printCentered("Edit Data Kamar", width)
	fmt.Println(strings.Repeat("═", width))
	var nomorKamar int

	fmt.Println("Masukkan nomor kamar yang ingin diedit:")
	fmt.Scan(&nomorKamar)

	ketemu := false
	for i := 0; i < numKamar; i++ {
		if kamar[i].NomorKamar == nomorKamar {
			ketemu = true
			var tipeBaru string
			var hargaBaru float64
			var fasilitasBaru string

			fmt.Println("Masukkan tipe kamar baru:")
			fmt.Scan(&tipeBaru)
			fmt.Println("Masukkan harga sewa per bulan baru:")
			fmt.Scan(&hargaBaru)
			fmt.Println("Masukkan fasilitas kamar baru (pisahkan dengan spasi, ketik '.' untuk selesai):")

			j := 0
			kamar[i].Fasilitas = [NMAX]string{}
			for {
				fmt.Scan(&fasilitasBaru)
				if fasilitasBaru == "." {
					break
				}
				kamar[i].Fasilitas[j] = fasilitasBaru
				j++
			}

			kamar[i].Tipe = tipeBaru
			kamar[i].Harga = hargaBaru

			fmt.Println("Data kamar berhasil diupdate.")
			break
		}
	}

	if !ketemu {
		fmt.Printf("Kamar dengan nomor %d tidak ditemukan.\n", nomorKamar)
	}
}

func hapusDataKamar() {
	clearScreen()
	setGreenText()
	width := 80
	fmt.Println(strings.Repeat("═", width))
	printCentered("Hapus Data Kamar", width)
	fmt.Println(strings.Repeat("═", width))
	resetTextColor()

	var nomorKamar int
	ketemu := false
	for i := 0; i < numKamar; i++ {
		if kamar[i].NomorKamar == nomorKamar {
			ketemu = true
			for j := i; j < numKamar-1; j++ {
				kamar[j] = kamar[j+1]
			}
			numKamar--
			fmt.Println("Data kamar berhasil dihapus.")
			break
		}
	}

	if !ketemu {
		fmt.Printf("Kamar dengan nomor %d tidak ditemukan.\n", nomorKamar)
	}
}

func editDataPenghuni() {
	clearScreen()
	setGreenText()
	var nomorKamar int
	fmt.Println("Masukkan nomor kamar penghuni yang ingin diedit:")
	fmt.Scan(&nomorKamar)

	ketemu := false
	for i := 0; i < numPenghuni; i++ {
		if penghuni[i].NomorKamar == nomorKamar {
			ketemu = true
			var alamatBaru string
			var tanggalMasukBaru string
			var jenisKamarBaru string

			fmt.Println("Masukkan alamat penghuni baru:")
			fmt.Scan(&alamatBaru)
			fmt.Println("Masukkan tanggal masuk penghuni baru:")
			fmt.Scan(&tanggalMasukBaru)
			fmt.Println("Masukkan jenis kamar penghuni baru:")
			fmt.Scan(&jenisKamarBaru)

			penghuni[i].Alamat = alamatBaru
			penghuni[i].TanggalMasuk = tanggalMasukBaru
			penghuni[i].JenisKamar = jenisKamarBaru

			fmt.Println("Data penghuni berhasil diupdate.")
			break
		}
	}

	if !ketemu {
		fmt.Printf("Tidak ada penghuni yang tinggal di kamar dengan nomor %d.\n", nomorKamar)
	}
}

func hapusDataPenghuni() {
	clearScreen()
	setGreenText()
	width := 80
	fmt.Println(strings.Repeat("═", width))
	printCentered("Hapus Data Penghuni", width)
	fmt.Println(strings.Repeat("═", width))

	var nomorKamar int
	ketemu := false
	for i := 0; i < numPenghuni; i++ {
		if penghuni[i].NomorKamar == nomorKamar {
			ketemu = true
			for j := 0; j < numKamar; j++ {
				if kamar[j].NomorKamar == nomorKamar {
					kamar[j].Status = "tersedia"
					break
				}
			}
			for j := i; j < numPenghuni-1; j++ {
				penghuni[j] = penghuni[j+1]
			}
			numPenghuni--
			fmt.Println("Data penghuni berhasil dihapus.")
			break
		}
	}

	if !ketemu {
		fmt.Printf("Tidak ada penghuni yang tinggal di kamar dengan nomor %d.\n", nomorKamar)
	}
}

func tampilkanDataPenghuni() {
	clearScreen()
	setGreenText()
	width := 80
	fmt.Println(strings.Repeat("═", width))
	printCentered("Tampilkan Data Penghuni", width)
	fmt.Println(strings.Repeat("═", width))
	fmt.Println("==========================================================================================================================================================")
	fmt.Printf("| %-20s | %-10s | %-15s | %-15s | %-50s | %-15s | %-15s |\n", "Nama", "No. Kamar", "No. Identitas", "Alamat", "No telpon", "Tanggal Masuk", "Jenis Kamar")
	fmt.Println("==========================================================================================================================================================")
	for i := 0; i < numPenghuni; i++ {
		fmt.Printf("| %-20s | %-10d | %-15d | %-15s | %-50s | %-15s | %-20s |\n", penghuni[i].Nama, penghuni[i].NomorKamar, penghuni[i].NomorIdentitas, penghuni[i].NomorTelepon, penghuni[i].Alamat, penghuni[i].TanggalMasuk, penghuni[i].JenisKamar)
	}
	fmt.Println("===========================================================================================================================================================")
	var keluar string
    fmt.Println("\nTekan 1 lalu enter untuk kembali ke menu utama...")
    fmt.Scan(&keluar)
	if keluar=="1"{
		menu()
	}
	
}

func tampilkanDataKamar() {
    // Tampilkan data kamar
    clearScreen()
    setGreenText()
    width := 80
    fmt.Println(strings.Repeat("═", width))
    printCentered("Tampilkan Data Kamar", width)
    fmt.Println(strings.Repeat("═", width))


    fmt.Println("===================================================================================================")
    fmt.Printf("| %-10s | %-15s | %-10s | %-60s | %-10s |\n", "No. Kamar", "Tipe Kamar", "Harga", "Fasilitas", "Status")
    fmt.Println("===================================================================================================")
    for i := 0; i < numKamar; i++ {
        fasilitas := ""
        for j := 0; j < len(kamar[i].Fasilitas); j++ {
            if kamar[i].Fasilitas[j] != "" {
                if j == 0 {
                    fasilitas += kamar[i].Fasilitas[j]
                } else {
                    fasilitas += ", " + kamar[i].Fasilitas[j]
                }
            }
        }
        fmt.Printf("| %-10d | %-15s | %-10.2f | %-60s | %-10s |\n", kamar[i].NomorKamar, kamar[i].Tipe, kamar[i].Harga, fasilitas, kamar[i].Status)
    }
    fmt.Println("===================================================================================================")


    var keluar string
    fmt.Println("\nTekan 1 lalu enter untuk kembali ke menu utama...")
    fmt.Scan(&keluar)
	if keluar=="1"{
		menu()
	}
	
}
