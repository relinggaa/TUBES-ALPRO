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

func header(judul string){
	if judul =="penghuni"{
		fmt.Printf("%25s.___                      __    __________                      .__                 .__ \n", "")
		fmt.Printf("%25s|   | ____ ______  __ ___/  |_  \\______   \\ ____   ____    ____ |  |__  __ __  ____ |__|\n", "")
		fmt.Printf("%25s|   |/    \\____ \\|  |  \\   __\\  |     ___// __ \\ /    \\  / ___\\|  |  \\|  |  \\/    \\|  |\n", "")
		fmt.Printf("%25s|   |   |  \\  |_> >  |  /|  |    |    |   \\  ___/|   |  \\/ /_/  >   Y  \\  |  /   |  \\  |\n", "")
		fmt.Printf("%25s|___|___|  /   __/|____/ |__|    |____|    \\___  >___|  /\\___  /|___|  /____/|___|  /__|\n", "")
		fmt.Printf("%25s         \\/|__|                                \\/     \\//_____/      \\/           \\/    \n", "")
	}else if judul=="kamar"{
		fmt.Printf("%25s.___                      __     ____  __.                            \n", "")
		fmt.Printf("%25s|   | ____ ______  __ ___/  |_  |    |/ _|____    _____ _____ _______ \n", "")
		fmt.Printf("%25s|   |/    \\____ \\|  |  \\   __\\ |      < \\__  \\  /     \\__  \\_  __ \\ \n", "")
		fmt.Printf("%25s|   |   |  \\  |_> >  |  /|  |   |    |  \\ / __ \\|  Y Y  \\/ __ \\|  | \\/\n", "")
		fmt.Printf("%25s|___|___|  /   __/|____/ |__|   |____|__ (____  /__|_|  (____  /__|   \n", "")
		fmt.Printf("%25s         \\/|__|                         \\/    \\/      \\/     \\/     \n", "")
		
	}else if judul=="carikamar"{
		fmt.Printf("%25s__________                                  .__                 ____  __.                            \n", "")
		fmt.Printf("%25s\\______   \\ ____   ____   ____ _____ _______|__|____    ____   |    |/ _|____    _____ _____ _______ \n", "")
		fmt.Printf("%25s |     ___// __ \\ /    \\_/ ___\\__  \\_  __ \\  \\__  \\  /    \\  |      < \\__  \\  /     \\__  \\_  __ \\\n", "")
		fmt.Printf("%25s |    |   \\  ___/|   |  \\  \\___ / __ \\|  | \\/  |/ __ \\|   |  \\ |    |  \\ / __ \\|  Y Y  \\/ __ \\|  | \\\n", "")
		fmt.Printf("%25s |____|    \\___  >___|  /\\___  >____  /__|  |__(____  /___|  / |____|__ (____  /__|_|  (____  /__|   \n", "")
		fmt.Printf("%25s               \\/     \\/     \\/     \\/              \\/     \\/          \\/    \\/      \\/     \\/       \n", "")

	}else if judul=="editkamar"{
		fmt.Printf("%25s___________    .___.__  __     ____  __.                            \n", "")
		fmt.Printf("%25s\\_   _____/  __| _/|__|/  |_  |    |/ _|____    _____ _____ _______ \n", "")
		fmt.Printf("%25s |    __)_  / __ | |  \\   __\\ |      < \\__  \\  /     \\\\__  \\\\_  __ \\\n", "")
		fmt.Printf("%25s |        \\/ /_/ | |  ||  |   |    |  \\ / __ \\|  Y Y  \\/ __ \\|  | \\/\n", "")
		fmt.Printf("%25s/_______  /\\____ | |__||__|   |____|__ (____  /__|_|  (____  /__|   \n", "")
		fmt.Printf("%25s        \\/      \\/                    \\/    \\/      \\/     \\/     \n", "")

	}else if judul=="caripenghuni"{
		fmt.Printf("%25s_________              .__  __________                      .__                 .__ \n", "")
		fmt.Printf("%25s\\_   ___ \\_____ _______|__| \\______   \\ ____   ____    ____ |  |__  __ __  ____ |__|\n", "")
		fmt.Printf("%25s/    \\  \\/\\__  \\\\_  __ \\  |  |     ___// __ \\ /    \\  / ___\\|  |  \\|  |  \\/    \\|  |\n", "")
		fmt.Printf("%25s\\     \\____/ __ \\|  | \\/  |  |    |   \\  ___/|   |  \\/ /_/  >   Y  \\  |  /   |  \\  |\n", "")
		fmt.Printf("%25s \\______  (____  /__|  |__|  |____|    \\___  >___|  /\\___  /|___|  /____/|___|  /__|\n", "")
		fmt.Printf("%25s        \\/     \\/                          \\/     \\//_____/      \\/           \\/    \n", "")

	}else if judul == "hapuskamar"{
		fmt.Printf("%25s  ___ ___                               ____  __.\n", "")
		fmt.Printf("%25s /   |   \\_____  ______  __ __  ______ |    |/ _|____    _____ _____ _______ \n", "")
		fmt.Printf("%25s/    ~    \\__  \\ \\____ \\|  |  \\/  ___/ |      < \\__  \\  /     \\\\__  \\\\ _  __ \\\n", "")
		fmt.Printf("%25s\\    Y    // __ \\|  |_> >  |  /\\___ \\  |    |  \\ / __ \\|  Y Y  \\\\/ __ \\|  | \\\n", "")
		fmt.Printf("%25s \\___|_  /(____  /   __/|____//____  > |____|__ (____  /__|_|  (____  /__|   \n", "")
		fmt.Printf("%25s       \\/      \\/|__|              \\/          \\/    \\/      \\/     \\/     \n", "")


	}else if judul=="hapuspenghuni"{
		fmt.Printf("%25s ___ ___                              __________                      .__                 .__ \n", "")
		fmt.Printf("%25s/   |   \\_____  ______  __ __  ______ \\______   \\ ____   ____    ____ |  |__  __ __  ____ |__|\n", "")
		fmt.Printf("%25s/    ~    \\__  \\ \\____ \\|  |  \\/  ___/  |     ___// __ \\ /    \\  / ___\\|  |  \\|  |  \\/    \\|  |\n", "")
		fmt.Printf("%25s\\    Y    // __ \\|  |_> >  |  /\\___ \\   |    |   \\  ___/|   |  \\/ /_/  >   Y  \\  |  /   |  \\  |\n", "")
		fmt.Printf("%25s \\___|_  /(____  /   __/|____//____  >  |____|    \\___  >___|  /\\___  /|___|  /____/|___|  /__|\n", "")
		fmt.Printf("%25s       \\/      \\/|__|              \\/                 \\/     \\//_____/      \\/           \\/    \n", "")


	}else if judul =="tampilpenghuni"{
		fmt.Printf("%-25s___________                     .__.__   __                   __________                      .__                 .__ \n", "")
		fmt.Printf("%-25s\\__    ___/____    _____ ______ |__|  | |  | ______    ____   \\______   \\ ____   ____    ____ |  |__  __ __  ____ |__|\n", "")
		fmt.Printf("%-25s  |    |  \\__  \\  /     \\\\____ \\|  |  | |  |/ |__  \\  /    \\   |     ___// __ \\ /    \\  / ___\\|  |  \\|  |  \\/    \\|  |\n", "")
		fmt.Printf("%-25s  |    |   / __ \\|  Y Y  \\  |_> >  |  |_|    < / __ \\|   |  \\  |    |   \\  ___/|   |  \\/ /_/  >   Y  \\  |  /   |  \\  |\n", "")
		fmt.Printf("%-25s  |____|  (____  /__|_|  /   __/|__|____/__|_ (____  /___|  /  |____|    \\___  >___|  /\\___  /|___|  /____/|___|  /__|\n", "")
		fmt.Printf("%-25s               \\/      \\/|__|                \\/    \\/     \\/                 \\/     \\//_____/      \\/           \\/    \n", "")


	}else if judul=="tampilkamar"{
		fmt.Printf("%-25s___________                     .__.__   __                    ____  __.                            \n", "")
		fmt.Printf("%-25s\\__    ___/____    _____ ______ |__|  | |  | ______    ____   |    |/ _|____    _____ _____ _______ \n", "")
		fmt.Printf("%-25s  |    |  \\__  \\  /     \\\\____ \\|  |  | |  |/ |__  \\  /    \\  |      < \\__  \\  /     \\\\__  \\_  __ \\ \n", "")
		fmt.Printf("%-25s  |    |   / __ \\|  Y Y  \\  |_> >  |  |_|    < / __ \\|   |  \\ |    |  \\ / __ \\|  Y Y  \\/ __ \\|  | \\/\n", "")
		fmt.Printf("%-25s  |____|  (____  /__|_|  /   __/|__|____/__|_ (____  /___|  / |____|__ (____  /__|_|  (____  /__|   \n", "")
		fmt.Printf("%-25s               \\/      \\/|__|                \\/    \\/     \\/          \\/    \\/      \\/     \\/       \n", "")

	}else if judul == "editpenghuni"{
		fmt.Printf("%-25s___________    .___.__  __    __________                      .__                 .__ \n", "")
		fmt.Printf("%-25s\\_   _____/  __| _/|__|/  |_  \\______   \\ ____   ____    ____ |  |__  __ __  ____ |__|\n", "")
		fmt.Printf("%-25s |    __)_  / __ | |  \\   __\\  |     ___// __ \\ /    \\  / ___\\|  |  \\|  |  \\/    \\|  |\n", "")
		fmt.Printf("%-25s |        \\/ /_/ | |  ||  |    |    |   \\  ___/|   |  \\/ /_/  >   Y  \\  |  /   |  \\  |\n", "")
		fmt.Printf("%-25s/_______  /\\____ | |__||__|    |____|    \\___  >___|  /\\___  /|___|  /____/|___|  /__|\n", "")
		fmt.Printf("%-25s        \\/      \\/                           \\/     \\//_____/      \\/           \\/    \n", "")


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
			fmt.Println(stop())
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}


func inputDataKamar() {
	clearScreen()
	setGreenText()
	header("kamar")
	


	var kamarBaru Kamar
	var inputVasilitas string
	var nf int
	var i int

	fmt.Printf("%25sMasukkan nomor kamar: ", "")
	fmt.Scan(&kamarBaru.NomorKamar)
	
	// Periksa apakah nomor kamar sudah ada
	for j := 0; j < numKamar; j++ {
		if kamar[j].NomorKamar == kamarBaru.NomorKamar {
			fmt.Println("Maaf, nomor kamar sudah ada.,Tekan 1 untuk memasukan kembali")
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
	clearScreen()
	setGreenText()
	fmt.Println("Data kamar berhasil dimasukkan.")
	exit()
}


func pendataanPenghuni() {
	clearScreen()
    setGreenText()
    header("penghuni")

    var penghuniBaru Penghuni
	var namaDepanpenghuni string
	var namaBelakangpenghuni string
	
	var namaPenghuni string
	

	namaDepanpenghuni=" "
	namaBelakangpenghuni=" "
	fmt.Printf("%25sMasukkan nama depan penghuni :", "")
    fmt.Scan(&namaDepanpenghuni)
	fmt.Printf("%25sMasukkan nama belakang penghuni :", "")
    fmt.Scan(&namaBelakangpenghuni)
	namaPenghuni=namaDepanpenghuni+" "+namaBelakangpenghuni
   	penghuniBaru.Nama=namaPenghuni

    fmt.Printf("%25sMasukkan nomor kamar penghuni:", "")
    fmt.Scan(&penghuniBaru.NomorKamar)
    if !cekKetersediaanKamar(penghuniBaru.NomorKamar) {
        fmt.Printf("%25sKamar tidak tersedia. Silakan pilih kamar lain.Tekan 2 untuk memasukan kembali\n", "")
        return
    }

    fmt.Printf("%25sMasukkan nomor identitas penghuni:", "")
    fmt.Scan(&penghuniBaru.NomorIdentitas)
    fmt.Printf("%25sMasukkan nomor telepon penghuni:", "")
    fmt.Scan(&penghuniBaru.NomorTelepon)
    fmt.Printf("%25sMasukkan alamat penghuni:", "")
    fmt.Scan(&penghuniBaru.Alamat)
    fmt.Printf("%25sMasukkan tanggal masuk penghuni (DD-MM-YY):", "")
    fmt.Scan(&penghuniBaru.TanggalMasuk)
    // fmt.Printf("%25sMasukkan jenis kamar penghuni:", "")
    // fmt.Scan(&penghuniBaru.JenisKamar)
    fmt.Printf("%25sDurasi sewa perbulan:", "")
    fmt.Scan(&penghuniBaru.LamaSewa)
    fmt.Printf("%25sJumlah uang yang dibayarkan:", "")
    fmt.Scan(&penghuniBaru.Pembayaran)

    // Temukan harga kamar
	var hargaKamar float64
	kamarFound := false
	for i := 0; i < numKamar && !kamarFound; i++ {
		if kamar[i].NomorKamar == penghuniBaru.NomorKamar {
			hargaKamar = kamar[i].Harga
			kamar[i].Status = penghuniBaru.Nama
			penghuniBaru.JenisKamar = kamar[i].Tipe
			kamarFound = true
		}
	}
	

    penghuniBaru.Total = int(hargaKamar) * penghuniBaru.LamaSewa
    penghuniBaru.Lunas = false


    penghuni[numPenghuni] = penghuniBaru
    numPenghuni++
    clearScreen()
    setGreenText()
    fmt.Print("Penghuni berhasil didata.\n", "")
    exit()
}


func pencarianDataKamar() {
    clearScreen()
    setGreenText()
    header("carikamar")

    var nomorKamarDicari int
    ketemu := false
    fmt.Println()
    fmt.Println("Masukkan nomor kamar yang ingin dicari:")
    fmt.Scan(&nomorKamarDicari)

    for i := 0; i < numKamar && !ketemu; i++ {
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
        }
    }

    if !ketemu {
        clearScreen()
        setGreenText()
        fmt.Printf("Kamar dengan nomor %d tidak ditemukan.\n", nomorKamarDicari)
    }
    exit()
}

func pencarianDataPenghuni() {    
    clearScreen()
    setGreenText()
    header("caripenghuni")

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
        }
    }

    if !ketemu {
        clearScreen()
        setGreenText()
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
    header("pembayaran")

    var nomorKamar int
    var bayar int
    var found bool = false
    var index int
    fmt.Print("Masukkan nomor kamar yang akan melakukan pembayaran: ")
    fmt.Scan(&nomorKamar)

    for i := 0; i < numPenghuni && !found; i++ {
        if penghuni[i].NomorKamar == nomorKamar {
            index = i
            found = true
        }
    }

    if found {
        if penghuni[index].Pembayaran >= penghuni[index].Total {
            penghuni[index].Lunas = true
            clearScreen()
            fmt.Println("Pembayaran sudah lunas.")
        } else {
            sisaPembayaran := penghuni[index].Total - penghuni[index].Pembayaran
            fmt.Println("Sisa uang yang belum dibayarkan:", sisaPembayaran)
            fmt.Print("Jumlah uang yang dibayarkan: ")
            fmt.Scan(&bayar)
            penghuni[index].Pembayaran += bayar

            if penghuni[index].Pembayaran >= penghuni[index].Total {
                penghuni[index].Lunas = true
                clearScreen()
                fmt.Println("Pembayaran sudah lunas.")
            } else {
                fmt.Println("Pembayaran berhasil. Sisa pembayaran:", penghuni[index].Total - penghuni[index].Pembayaran)
            }
        }
    } else {
        fmt.Println("Tidak ditemukan kamar dengan nomor tersebut.")
    }

    exit()
}




func editDataKamar() {
	clearScreen()
	setGreenText()
	header("editkamar")

	var nomorKamar int
	fmt.Printf("%25sMasukkan nomor kamar yang ingin diedit:", "")
	fmt.Scan(&nomorKamar)
	ketemu := false
	for i := 0; i < numKamar; i++ {
		if kamar[i].NomorKamar == nomorKamar {
			ketemu = true
			var tipeBaru string
			var hargaBaru float64
			var fasilitasBaru string
			var fasilitasSelesai bool

			fmt.Printf("%25sMasukkan tipe kamar baru:", "")
			fmt.Scan(&tipeBaru)
			fmt.Printf("%25sMasukkan harga sewa per bulan baru:", "")
			fmt.Scan(&hargaBaru)
			fmt.Printf("%25sMasukkan fasilitas kamar baru (pisahkan dengan spasi, ketik '.' untuk selesai):", "")

			j := 0
			kamar[i].Fasilitas = [NMAX]string{}
			for !fasilitasSelesai {
				fmt.Scan(&fasilitasBaru)
				if fasilitasBaru == "." {
					fasilitasSelesai = true
				} else {
					kamar[i].Fasilitas[j] = fasilitasBaru
					j++
				}
			}

			kamar[i].Tipe = tipeBaru
			kamar[i].Harga = hargaBaru

			// Perbarui tipe kamar penghuni yang memiliki nomor kamar ini
			for j := 0; j < numPenghuni; j++ {
				if penghuni[j].NomorKamar == nomorKamar {
					penghuni[j].JenisKamar = tipeBaru
				}
			}

			clearScreen()
			setGreenText()

			fmt.Print("Data kamar berhasil diupdate.", "")
		}
	}

	if !ketemu {
		clearScreen()
		setGreenText()
		fmt.Printf("Kamar dengan nomor %d tidak ditemukan.\n", nomorKamar)
	}
	exit()
}


func hapusDataKamar() {
    clearScreen()
    setGreenText()
    header("hapuskamar")

    var nomorKamar int
    fmt.Printf("%25s\n", "Masukkan nomor kamar yang ingin dihapus:")
    fmt.Scan(&nomorKamar)

    ketemu := false
    for i := 0; i < numKamar; i++ {
        if kamar[i].NomorKamar == nomorKamar {
            ketemu = true
            for j := i; j < numKamar-1; j++ {
                kamar[j] = kamar[j+1]
            }
            clearScreen()
            setGreenText()
            numKamar--
            fmt.Println("Data kamar berhasil dihapus.")
        }
    }

    if !ketemu {
        clearScreen()
        setGreenText()
        fmt.Printf("Kamar dengan nomor %d tidak ditemukan.\n", nomorKamar)
    }
    exit()
}

	


func editDataPenghuni() {
	clearScreen()
	setGreenText()
	header("editpenghuni")

	var nomorKamar int
	fmt.Print("Masukkan nomor kamar penghuni yang ingin diedit:")
	fmt.Scan(&nomorKamar)

	ketemu := false
	for i := 0; i < numPenghuni; i++ {
		if penghuni[i].NomorKamar == nomorKamar {
			ketemu = true
			var namaBaru string
			var alamatBaru string
			var tanggalMasukBaru string
			var namaDepanpenghuni string
			var namaBelakangpenghuni string
			
			var namaPenghuni string
			namaDepanpenghuni=" "
			namaBelakangpenghuni=" "
			fmt.Printf("%25sMasukkan nama depan penghuni :", "")
			fmt.Scan(&namaDepanpenghuni)
			fmt.Printf("%25sMasukkan nama belakang penghuni :", "")
			fmt.Scan(&namaBelakangpenghuni)
			namaPenghuni=namaDepanpenghuni+" "+namaBelakangpenghuni
			namaBaru=namaPenghuni
			fmt.Printf("%25sMasukkan alamat penghuni baru:","")
			fmt.Scan(&alamatBaru)
			fmt.Printf("%25sMasukkan tanggal masuk penghuni baru DDD-MMM-YY:","")
			fmt.Scan(&tanggalMasukBaru)
			penghuni[i].Nama = namaBaru
			penghuni[i].Alamat = alamatBaru
			penghuni[i].TanggalMasuk = tanggalMasukBaru
			kamar[i].Status=namaBaru
		
			clearScreen()
			setGreenText()
			fmt.Print("Data penghuni berhasil diupdate.")
			exit()
		}
	}
	clearScreen()
	setGreenText()
	if !ketemu {
		fmt.Printf("Tidak ada penghuni yang tinggal di kamar dengan nomor %d.\n", nomorKamar)
	}
	exit()
}

func hapusDataPenghuni() {
    clearScreen()
    setGreenText()
    header("hapuspenghuni")

    var nomorKamar int
    ketemu := false
    fmt.Printf("%25s\n", "Masukkan nomor kamar yang ingin dihapus:")
    fmt.Scan(&nomorKamar)

    for i := 0; i < numPenghuni; i++ {
        if penghuni[i].NomorKamar == nomorKamar {
            ketemu = true
            for j := 0; j < numKamar; j++ {
                if kamar[j].NomorKamar == nomorKamar {
                    kamar[j].Status = "tersedia"
                }
            }
            for j := i; j < numPenghuni-1; j++ {
                penghuni[j] = penghuni[j+1]
            }
            clearScreen()
            setGreenText()
            numPenghuni--
            fmt.Printf("%25s\n", "Data penghuni berhasil dihapus.")
            exit()
        }
    }

    if !ketemu {
        clearScreen()
        setGreenText()
        fmt.Printf("Kamar dengan nomor %d tidak ditemukan.\n", nomorKamar)
    }
    exit()
}

func tampilkanDataPenghuni() {
    clearScreen()
    setGreenText()
    header("tampilpenghuni")
    
    if numPenghuni == 0 {
        fmt.Println("==========================================================================================================================================================")
        fmt.Println("|                                                     Data penghuni masih kosong                                                                         |")
        fmt.Println("==========================================================================================================================================================")
    } else {
        // selection sort 
        for i := 0; i < numPenghuni-1; i++ {
            minIdx := i
            for j := i + 1; j < numPenghuni; j++ {
                if penghuni[j].NomorKamar < penghuni[minIdx].NomorKamar {
                    minIdx = j
                }
            }
            penghuni[i], penghuni[minIdx] = penghuni[minIdx], penghuni[i]
        }

        fmt.Println("==========================================================================================================================================================")
        fmt.Printf("| %-20s | %-10s | %-15s | %-15s | %-50s | %-15s | %-15s |\n", "Nama", "No. Kamar", "No. Identitas", "No.Telepon", "Alamat", "Tanggal Masuk", "Jenis Kamar")
        fmt.Println("==========================================================================================================================================================")
        for i := 0; i < numPenghuni; i++ {
            fmt.Printf("| %-20s | %-10d | %-15d | %-15s | %-50s | %-15s | %-15s |\n", penghuni[i].Nama, penghuni[i].NomorKamar, penghuni[i].NomorIdentitas, penghuni[i].NomorTelepon, penghuni[i].Alamat, penghuni[i].TanggalMasuk, penghuni[i].JenisKamar)
        }
        fmt.Println("==========================================================================================================================================================")
    }
    exit()
}




func tampilkanDataKamar() {
    clearScreen()
    setGreenText()
    header("tampilkamar")

    if numKamar == 0 {
        fmt.Println("===================================================================================================")
        fmt.Println("|                                       Data kamar masih kosong                                      |")
        fmt.Println("===================================================================================================")
    } else {
        //selection sort
        for i := 0; i < numKamar-1; i++ {
            minIdx := i
            for j := i + 1; j < numKamar; j++ {
                if kamar[j].NomorKamar < kamar[minIdx].NomorKamar {
                    minIdx = j
                }
            }
            kamar[i], kamar[minIdx] = kamar[minIdx], kamar[i]
        }

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
    }
    exit()
}

func stop() string {
	clearScreen()
	setGreenText()
	message := "Terima kasih. Sampai jumpa!"
	fmt.Println(message)
	numKamar = 0
	numPenghuni = 0
	os.Exit(0)
	return message 
}
