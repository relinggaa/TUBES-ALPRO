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
func exit(){
	var keluar string
    fmt.Println("\nTekan 1 lalu enter untuk kembali ke menu utama...")
    fmt.Scan(&keluar)
	if  keluar=="1"{
		menu()
	}
	
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
	var tampilkandataKamar, tampilkanpenghuni, tampilkanpembayaran, hapuskamar, inputkamar, inputpenghuni, editkamar, editpenghuni, pencariankamar, hapuspenghuni, caripenghuni bool

	for {
		if !tampilkandataKamar && !tampilkanpenghuni && !tampilkanpembayaran &&!inputkamar&&!inputpenghuni&&!editkamar&& !editpenghuni &&!pencariankamar &&!hapuskamar&&!hapuspenghuni&&!caripenghuni{
		clearScreen()
		setGreenText()
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
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "4.Pencarian Data Penghuni")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "5.Edit data kamar")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "6.Edit data penghuni")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "7.Hapus data kamar")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "8.Hapus data penghuni")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "9.Tampilkan data penghuni")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "10.Tampilkan data kamar")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "11.Pembayaran penghuni")
		fmt.Printf("%25s║%4s%-69s║\n", "", "", "12. Keluar")
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
			inputkamar=true
		case 2:
			pendataanPenghuni()
			inputpenghuni=true
		case 3:
			pencarianDataKamar()
			pencariankamar=true
		case 4:
			pencarianDataPenghuni()
			caripenghuni=true
		case 5:
			editDataKamar()
			editkamar=true
		case 6:
			editDataPenghuni()
			editpenghuni=true
		case 7:
			hapusDataKamar()
			hapuskamar=true
		case 8:
			hapusDataPenghuni()
			hapuspenghuni=true
		case 9:
			tampilkanDataPenghuni()
			tampilkanpenghuni=true
		case 10:
			tampilkanDataKamar()
			tampilkandataKamar=true
		case 11:
			pembayaran()
			tampilkanpembayaran=true
		case 12:
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
	fmt.Printf("%25s /$$$$$$                                 /$$           /$$$$$$$              /$$                     /$$$$$$$                               /$$                           /$$      \n", "")
	fmt.Printf("%25s|_  $$_/                                | $$          | $$__  $$            | $$                    | $$__  $$                             | $$                          |__/      \n", "")
	fmt.Printf("%25s  | $$   /$$$$$$$   /$$$$$$  /$$   /$$ /$$$$$$        | $$  \\ $$  /$$$$$$  /$$$$$$    /$$$$$$       | $$  \\ $$ /$$$$$$  /$$$$$$$   /$$$$$$ | $$$$$$$  /$$   /$$ /$$$$$$$  /$$      \n", "")
	fmt.Printf("%25s  | $$  | $$__  $$ /$$__  $$| $$  | $$|_  $$_/        | $$  | $$ |____  $$|_  $$_/   |____  $$      | $$$$$$$//$$__  $$| $$__  $$ /$$__  $$| $$__  $$| $$  | $$| $$__  $$| $$      \n", "")
	fmt.Printf("%25s  | $$  | $$  \\ $$| $$  \\ $$| $$  | $$  | $$          | $$  | $$  /$$$$$$$  | $$      /$$$$$$$      | $$____/| $$$$$$$$| $$  \\ $$| $$  \\ $$| $$  \\ $$| $$  | $$| $$  \\ $$| $$      \n", "")
	fmt.Printf("%25s  | $$  | $$  | $$| $$  | $$| $$  | $$  | $$ /$$      | $$  | $$ /$$__  $$  | $$ /$$ /$$__  $$      | $$     | $$_____/| $$  | $$| $$  | $$| $$  | $$| $$  | $$| $$  | $$| $$      \n", "")
	fmt.Printf("%25s /$$$$$$| $$  | $$| $$$$$$$/|  $$$$$$/  |  $$$$/      | $$$$$$$/|  $$$$$$$  |  $$$$/|  $$$$$$$      | $$     |  $$$$$$$| $$  | $$|  $$$$$$$| $$  | $$|  $$$$$$/| $$  | $$| $$      \n", "")
	fmt.Printf("%25s|______/|__/  |__/| $$____/  \\______/    \\___/        |_______/  \\_______/   \\___/   \\_______/      |__/      \\_______/|__/  |__/ \\____  $$|__/  |__/ \\______/ |__/  |__/|__/      \n", "")
	fmt.Printf("%25s                  | $$                                                                                                            /$$  \\ $$                                        \n", "")
	fmt.Printf("%25s                  | $$                                                                                                           |  $$$$$$/                                         \n", "")
	fmt.Printf("%25s                  |__/                                                                                                            \\______/                                         \n", "")
	



	var penghuniBaru Penghuni
	fmt.Printf("%25sMasukkan nama penghuni:", "")
	fmt.Scan(&penghuniBaru.Nama)
	fmt.Printf("%25sMasukkan nomor kamar penghuni:", "")
	fmt.Scan(&penghuniBaru.NomorKamar)
	if !cekKetersediaanKamar(penghuniBaru.NomorKamar) {
		fmt.Printf("%25sKamar tidak tersedia. Silakan pilih kamar lain.\n", "")
		return
	}
	fmt.Printf("%25sMasukkan nomor identitas penghuni:", "")
	fmt.Scan(&penghuniBaru.NomorIdentitas)
	fmt.Printf("%25sMasukkan nomor telepon penghuni:", "")
	fmt.Scan(&penghuniBaru.NomorTelepon)
	fmt.Printf("%25sMasukkan alamat penghuni:", "")
	fmt.Scan(&penghuniBaru.Alamat)
	fmt.Printf("%25sMasukkan tanggal masuk penghuni (YYYY-MM-DD):", "")
	fmt.Scan(&penghuniBaru.TanggalMasuk)
	fmt.Printf("%25sMasukkan jenis kamar penghuni:", "")
	fmt.Scan(&penghuniBaru.JenisKamar)
	fmt.Printf("%25sDurasi sewa:", "")
	fmt.Scan(&penghuniBaru.LamaSewa)
	fmt.Printf("%25sJumlah uang yang dibayarkan:", "")
	fmt.Scan(&penghuniBaru.Pembayaran)

	// Find the price of the room
	var hargakamar float64
	for i := 0; i < numKamar; i++ {
		if kamar[i].NomorKamar == penghuniBaru.NomorKamar {
			hargakamar = kamar[i].Harga
			kamar[i].Status = penghuniBaru.Nama
			break
		}
	}

	penghuniBaru.Total = int(hargakamar) * penghuniBaru.LamaSewa
	penghuniBaru.Lunas = false

	penghuni[numPenghuni] = penghuniBaru
	numPenghuni++
	fmt.Printf("%25sPenghuni berhasil didata.\n", "")
	exit()
}


func inputDataKamar() {
	clearScreen()
	setGreenText()
	fmt.Printf("%25s$$$$$$\\                                 $$\\           $$$$$$$\\             $$\\                     $$\\   $$\\                                            \n", "")
	fmt.Printf("%25s\\_$$  _|                                $$ |          $$  __$$\\            $$ |                    $$ | $$  |                                           \n", "")
	fmt.Printf("%25s  $$ |  $$$$$$$\\   $$$$$$\\  $$\\   $$\\ $$$$$$\\         $$ |  $$ | $$$$$$\\ $$$$$$\\    $$$$$$\\        $$ |$$  / $$$$$$\\  $$$$$$\\$$$$\\   $$$$$$\\   $$$$$$\\  \n", "")
	fmt.Printf("%25s  $$ |  $$  __$$\\ $$  __$$\\ $$ |  $$ |\\_$$  _|        $$ |  $$ | \\____$$\\\\_$$  _|   \\____$$\\       $$$$$  /  \\____$$\\ $$  _$$  _$$\\  \\____$$\\ $$  __$$\\ \n", "")
	fmt.Printf("%25s  $$ |  $$ |  $$ |$$ /  $$ |$$ |  $$ |  $$ |          $$ |  $$ | $$$$$$$ | $$ |     $$$$$$$ |      $$  $$<   $$$$$$$ |$$ / $$ / $$ | $$$$$$$ |$$ |  \\__|\n", "")
	fmt.Printf("%25s  $$ |  $$ |  $$ |$$ |  $$ |$$ |  $$ |  $$ |$$\\       $$ |  $$ |$$  __$$ | $$ |$$\\ $$  __$$ |      $$ |\\$$\\ $$  __$$ |$$ | $$ | $$ |$$  __$$ |$$ |      \n", "")
	fmt.Printf("%25s$$$$$$\\ $$ |  $$ |$$$$$$$  |\\$$$$$$  |  \\$$$$  |      $$$$$$$  |\\$$$$$$$ | \\$$$$  |\\$$$$$$$ |      $$ | \\$$\\$$$$$$$ |$$ | $$ | $$ |\\$$$$$$$ |$$ |      \n", "")
	fmt.Printf("%25s\\______|\\__|  \\__|$$  ____/  \\______/    \\____/       \\_______/  \\_______|  \\____/  \\_______|      \\__|  \\__|\\_______|\\__| \\__| \\__| \\_______|\\__|      \n", "")
	fmt.Printf("%25s                  $$ |                                                                                                                                  \n", "")
	fmt.Printf("%25s                  $$ |                                                                                                                                  \n", "")
	fmt.Printf("%25s                  \\__|                                                                                                                                  \n", "")


	var kamarBaru Kamar
	var inputVasilitas string
	var nf int
	var i int

	fmt.Printf("%25sMasukkan nomor kamar: ", "")
	fmt.Scan(&kamarBaru.NomorKamar)
	
	// Periksa apakah nomor kamar sudah ada
	for j := 0; j < numKamar; j++ {
		if kamar[j].NomorKamar == kamarBaru.NomorKamar {
			fmt.Println("Maaf, nomor kamar sudah ada.")
			return
		}
	}
	
	fmt.Printf("%25sMasukkan tipe kamar: ", "")
	fmt.Scan(&kamarBaru.Tipe)
	fmt.Printf("%25sMasukkan harga sewa per bulan: ", "")
	fmt.Scan(&kamarBaru.Harga)
	fmt.Printf("%25sMasukkan fasilitas kamar (pisahkan dengan spasi, ketik '.' untuk selesai): ", "")
	
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
	exit()
}

func pencarianDataKamar() {
	clearScreen()
	setGreenText()
	fmt.Printf("%25s\n", "$$$$$$$\\                                                   $$\\                           $$$$$$$\\             $$\\                     $$\\   $$\\                                            ")
	fmt.Printf("%25s\n", "$$  __$$\\                                                  \\__|                          $$  __$$\\            $$ |                    $$ | $$  |                                           ")
	fmt.Printf("%25s\n", "$$ |  $$ | $$$$$$\\  $$$$$$$\\   $$$$$$$\\ $$$$$$\\   $$$$$$\\  $$\\  $$$$$$\\  $$$$$$$\\        $$ |  $$ | $$$$$$\\ $$$$$$\\    $$$$$$\\        $$ |$$  / $$$$$$\\  $$$$$$\\$$$$\\   $$$$$$\\   $$$$$$\\  ")
	fmt.Printf("%25s\n", "$$$$$$$  |$$  __$$\\ $$  __$$\\ $$  _____|\\____$$\\ $$  __$$\\ $$ | \\____$$\\ $$  __$$\\       $$ |  $$ | \\____$$\\\\_$$  _|   \\____$$\\       $$$$$  /  \\____$$\\ $$  _$$  _$$\\  \\____$$\\ $$  __$$\\ ")
	fmt.Printf("%25s\n", "$$  ____/ $$$$$$$$ |$$ |  $$ |$$ /      $$$$$$$ |$$ |  \\__|$$ | $$$$$$$ |$$ |  $$ |      $$ |  $$ | $$$$$$$ | $$ |     $$$$$$$ |      $$  $$<   $$$$$$$ |$$ / $$ / $$ | $$$$$$$ |$$ |  \\__|")
	fmt.Printf("%25s\n", "$$ |      $$   ____|$$ |  $$ |$$ |     $$  __$$ |$$ |      $$ |$$  __$$ |$$ |  $$ |      $$ |  $$ |$$  __$$ | $$ |$$\\ $$  __$$ |      $$ |\\$$\\ $$  __$$ |$$ | $$ | $$ |$$  __$$ |$$ |      ")
	fmt.Printf("%25s\n", "$$ |      \\$$$$$$$\\ $$ |  $$ |\\$$$$$$$\\\\$$$$$$$ |$$ |      $$ |\\$$$$$$$ |$$ |  $$ |      $$$$$$$  |\\$$$$$$$ | \\$$$$  |\\$$$$$$$ |      $$ | \\$$\\\\$$$$$$$ |$$ | $$ | $$ |\\$$$$$$$ |$$ |      ")
	fmt.Printf("%25s\n", "\\__|       \\_______|\\__|  \\__| \\_______|\\_______|\\__|      \\__| \\_______|\\__|  \\__|      \\_______/  \\_______|  \\____/  \\_______|      \\__|  \\__|\\_______|\\__| \\__| \\__| \\_______|\\__|      ")


		
	var nomorKamarDicari int
	ketemu := false
	fmt.Println()
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
	exit()
}
func pencarianDataPenghuni() {	
	var nomorKamarDicari int
	ketemu := false
	fmt.Println()
	fmt.Println("Masukkan nomor kamar yang ingin dicari:")
	fmt.Scan(&nomorKamarDicari)

	for i := 0; i < numPenghuni; i++ {
		if penghuni[i].NomorKamar == nomorKamarDicari {
			ketemu = true
			fmt.Printf("Penghuni dengan nomor kamar %d ditemukan.\n", nomorKamarDicari)
			fmt.Printf("Nama: %s\n", penghuni[i].Nama)
			fmt.Printf("Nomor Identitas: %d\n", penghuni[i].NomorIdentitas)
			fmt.Printf("Nomor Telepon: %s\n", penghuni[i].NomorTelepon)
			fmt.Printf("Alamat: %s\n", penghuni[i].Alamat)
			fmt.Printf("Tanggal Masuk: %s\n", penghuni[i].TanggalMasuk)
			fmt.Printf("Jenis Kamar: %s\n", penghuni[i].JenisKamar)
			fmt.Printf("Lama Sewa: %d\n", penghuni[i].LamaSewa)
			fmt.Printf("Pembayaran: %d\n", penghuni[i].Pembayaran)
			fmt.Printf("Total: %d\n", penghuni[i].Total)
			fmt.Printf("Lunas: %t\n", penghuni[i].Lunas)
			break
		}
	}
	if !ketemu {
		fmt.Printf("Penghuni dengan nomor kamar %d tidak ditemukan.\n", nomorKamarDicari)
	}
	exit()
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
	fmt.Printf("%25s\n", "$$$$$$$\\                                                   $$\\                           $$$$$$$\\             $$\\                     $$\\   $$\\                                            ")
	fmt.Printf("%25s\n", "$$  __$$\\                                                  \\__|                          $$  __$$\\            $$ |                    $$ | $$  |                                           ")
	fmt.Printf("%25s\n", "$$ |  $$ | $$$$$$\\  $$$$$$$\\   $$$$$$$\\ $$$$$$\\   $$$$$$\\  $$\\  $$$$$$\\  $$$$$$$\\        $$ |  $$ | $$$$$$\\ $$$$$$\\    $$$$$$\\        $$ |$$  / $$$$$$\\  $$$$$$\\$$$$\\   $$$$$$\\   $$$$$$\\  ")
	fmt.Printf("%25s\n", "$$$$$$$$  |$$  __$$\\ $$  __$$\\ $$  _____|\\____$$\\ $$  __$$\\ $$ | \\____$$\\ $$  __$$\\       $$ |  $$ | \\____$$\\\\_$$  _|   \\____$$\\       $$$$$  /  \\____$$\\ $$  _$$  _$$\\  \\____$$\\ $$  __$$\\ ")
	fmt.Printf("%25s\n", "$$  ____/ $$$$$$$$ |$$ |  $$ |$$ /      $$$$$$$ |$$ |  \\__|$$ | $$$$$$$ |$$ |  $$ |      $$ |  $$ | $$$$$$$ | $$ |     $$$$$$$ |      $$  $$<   $$$$$$$ |$$ / $$ / $$ | $$$$$$$ |$$ |  \\__|")
	fmt.Printf("%25s\n", "$$ |      $$   ____|$$ |  $$ |$$ |     $$  __$$ |$$ |      $$ |$$  __$$ |$$ |  $$ |      $$ |  $$ |$$  __$$ | $$ |$$\\ $$  __$$ |      $$ |\\$$\\ $$  __$$ |$$ | $$ | $$ |$$  __$$ |$$ |      ")
	fmt.Printf("%25s\n", "$$ |      \\$$$$$$$\\ $$ |  $$ |\\$$$$$$$\\\\$$$$$$$ |$$ |      $$ |\\$$$$$$$ |$$ |  $$ |      $$$$$$$  |\\$$$$$$$ | \\$$$$  |\\$$$$$$$ |      $$ | \\$$\\$$$$$$$ |$$ | $$ | $$ |\\$$$$$$$ |$$ |      ")
	fmt.Printf("%25s\n", "\\__|       \\_______|\\__|  \\__| \\_______|\\_______|\\__|      \\__| \\_______|\\__|  \\__|      \\_______/  \\_______|  \\____/  \\_______|      \\__|  \\__|\\_______|\\__| \\__| \\__| \\_______|\\__|      ")

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
	exit()
	
}


func editDataKamar() {
	clearScreen()
	setGreenText()
	fmt.Printf("%25s\n", " /$$$$$$$$       /$$ /$$   /$$           /$$$$$$$              /$$                     /$$   /$$                                             ")
	fmt.Printf("%25s\n", "| $$_____/      | $$|__/  | $$          | $$__  $$            | $$                    | $$  /$$/                                             ")
	fmt.Printf("%25s\n", "| $$        /$$$$$$$ /$$ /$$$$$$        | $$  \\ $$  /$$$$$$  /$$$$$$    /$$$$$$       | $$ /$$/   /$$$$$$  /$$$$$$/$$$$   /$$$$$$   /$$$$$$ ")
	fmt.Printf("%25s\n", "| $$$$$    /$$__  $$| $$|_  $$_/        | $$  | $$ |____  $$|_  $$_/   |____  $$      | $$$$$/   |____  $$| $$_  $$_  $$ |____  $$ /$$__  $$")
	fmt.Printf("%25s\n", "| $$__/   | $$  | $$| $$  | $$          | $$  | $$  /$$$$$$$  | $$      /$$$$$$$      | $$  $$    /$$$$$$$| $$ \\ $$ \\ $$  /$$$$$$$| $$  \\__/")
	fmt.Printf("%25s\n", "| $$      | $$  | $$| $$  | $$ /$$      | $$  | $$ /$$__  $$  | $$ /$$ /$$__  $$      | $$\\  $$  /$$__  $$| $$ | $$ | $$ /$$__  $$| $$      ")
	fmt.Printf("%25s\n", "| $$$$$$$$|  $$$$$$$| $$  |  $$$$/      | $$$$$$$/|  $$$$$$$  |  $$$$/|  $$$$$$$      | $$ \\  $$|  $$$$$$$| $$ | $$ | $$|  $$$$$$$| $$      ")
	fmt.Printf("%25s\n", "|________/ \\_______/|__/   \\___/        |_______/  \\_______/   \\___/   \\_______/      |__/  \\__/ \\_______/|__/ |__/ |__/ \\_______/|__/      ")


		var nomorKamar int
		fmt.Printf("%25s\n", "Masukkan nomor kamar yang ingin diedit:")
		fmt.Scan(&nomorKamar)
		
		ketemu := false
		for i := 0; i < numKamar; i++ {
			if kamar[i].NomorKamar == nomorKamar {
				ketemu = true
				var tipeBaru string
				var hargaBaru float64
				var fasilitasBaru string
		
				fmt.Printf("%25s\n", "Masukkan tipe kamar baru:")
				fmt.Scan(&tipeBaru)
				fmt.Printf("%25s\n", "Masukkan harga sewa per bulan baru:")
				fmt.Scan(&hargaBaru)
				fmt.Printf("%25s\n", "Masukkan fasilitas kamar baru (pisahkan dengan spasi, ketik '.' untuk selesai):")
		
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
			}
		}

	if !ketemu {
		fmt.Printf("Kamar dengan nomor %d tidak ditemukan.\n", nomorKamar)
	}
	exit()
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
	fmt.Println("Masukkan nomor kamar yang ingin dihapus:")
	fmt.Scan(&nomorKamar)

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
	exit()
}
	


func editDataPenghuni() {
	clearScreen()
	setGreenText()
	fmt.Printf("%25s\n", " /$$$$$$$$       /$$ /$$   /$$           /$$$$$$$              /$$                     /$$$$$$$                               /$$                           /$$")
	fmt.Printf("%25s\n", "| $$_____/      | $$|__/  | $$          | $$__  $$            | $$                    | $$__  $$                             | $$                          |__/")
	fmt.Printf("%25s\n", "| $$        /$$$$$$$ /$$ /$$$$$$        | $$  \\ $$  /$$$$$$  /$$$$$$    /$$$$$$       | $$  \\ $$ /$$$$$$  /$$$$$$$   /$$$$$$ | $$$$$$$  /$$   /$$ /$$$$$$$  /$$")
	fmt.Printf("%25s\n", "| $$$$$    /$$__  $$| $$|_  $$_/        | $$  | $$ |____  $$|_  $$_/   |____  $$      | $$$$$$$//$$__  $$| $$__  $$ /$$__  $$| $$__  $$| $$  | $$| $$__  $$| $$")
	fmt.Printf("%25s\n", "| $$__/   | $$  | $$| $$  | $$          | $$  | $$  /$$$$$$$  | $$      /$$$$$$$      | $$____/| $$$$$$$$| $$  \\ $$| $$  \\ $$| $$  \\ $$| $$  | $$| $$  \\ $$| $$")
	fmt.Printf("%25s\n", "| $$      | $$  | $$| $$  | $$ /$$      | $$  | $$ /$$__  $$  | $$ /$$ /$$__  $$      | $$     | $$_____/| $$  | $$| $$  | $$| $$  | $$| $$  | $$| $$  | $$| $$")
	fmt.Printf("%25s\n", "| $$$$$$$$|  $$$$$$$| $$  |  $$$$/      | $$$$$$$/|  $$$$$$$  |  $$$$/|  $$$$$$$      | $$     |  $$$$$$$| $$  | $$|  $$$$$$$| $$  | $$|  $$$$$$/| $$  | $$| $$")
	fmt.Printf("%25s\n", "|________/ \\_______/|__/   \\___/        |_______/  \\_______/   \\___/   \\_______/      |__/      \\_______/|__/  |__/ \\____  $$|__/  |__/ \\______/ |__/  |__/|__/")
	fmt.Printf("%25s\n", "                                                                                                                    /$$  \\ $$                                  ")
	fmt.Printf("%25s\n", "                                                                                                                   |  $$$$$$/                             buat ini juga")

	var nomorKamar int
	fmt.Println("Masukkan nomor kamar penghuni yang ingin diedit:")
	fmt.Scan(&nomorKamar)

	ketemu := false
	for i := 0; i < numPenghuni; i++ {
		if penghuni[i].NomorKamar == nomorKamar {
			ketemu = true
			var namaBaru string
			var alamatBaru string
			var tanggalMasukBaru string
			var jenisKamarBaru string
			fmt.Println("Masukkan nama penghuni baru:")
			fmt.Scan(&namaBaru)
			fmt.Println("Masukkan alamat penghuni baru:")
			fmt.Scan(&alamatBaru)
			fmt.Println("Masukkan tanggal masuk penghuni baru:")
			fmt.Scan(&tanggalMasukBaru)
			fmt.Println("Masukkan jenis kamar penghuni baru:")
			fmt.Scan(&jenisKamarBaru)
			penghuni[i].Nama = namaBaru
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
	exit()
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
	exit()
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
	exit()
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
	exit()

    
	
	
}
