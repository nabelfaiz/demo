package main

import (
	"GO/model"
	"GO/node"
	"GO/view"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func menuUtama() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	for {
		reader.Discard(reader.Buffered())

		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Tambah Data Sepatu")
		fmt.Println("2. Tambah Data Pembeli")
		fmt.Println("3. Tampilkan Data Sepatu dan Pembeli")
		fmt.Println("4. Update Data Sepatu")
		fmt.Println("5. Update Data Pembeli")
		fmt.Println("6. Hapus Data Sepatu")
		fmt.Println("7. Hapus Data Pembeli")
		fmt.Println("8. Cari Data Sepatu")
		fmt.Println("9. Cari Data Pembeli")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		if scanner.Scan() {
			input := scanner.Text()
			input = strings.TrimSpace(input)
			choice, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Input Tidak Valid, silakan masukkan angka")
				continue
			}
			switch choice {
			case 1:
				view.Insert()

			case 3:
				view.Views()
			case 4:
				view.Update()

			case 6:
				view.Delete()

			case 8:
				view.Search()

			case 0:
				fmt.Println("Terima Kasih.")
				return
			default:
				fmt.Println("Pilihan tidak valid, silakan coba lagi.")
			}
		}
	}
}

func main() {
	// Inisialisasi data (jika diperlukan)
	model.DaftarSepatu = node.ListSepatu{}
	model.DaftarPembeli = node.ListPembeli{}

	// Tambahkan data sepatu di awal (opsional)
	model.CreateSepatu(node.Sepatu{IDSepatu: 1, NamaSepatu: "Nike Air Max", HargaSepatu: 1200000})
	model.CreateSepatu(node.Sepatu{IDSepatu: 2, NamaSepatu: "Adidas Ultraboost", HargaSepatu: 1500000})

	// Tampilkan daftar sepatu di awal (opsional)
	fmt.Println("Daftar Sepatu yang Tersedia:")
	for _, sepatu := range model.ReadSepatu() {
		fmt.Printf("ID: %d, Nama: %s, Harga: %d\n", sepatu.IDSepatu, sepatu.NamaSepatu, sepatu.HargaSepatu)
	}

	// Jalankan menu utama
	menuUtama()
}
