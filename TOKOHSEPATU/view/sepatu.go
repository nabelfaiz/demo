package view

import (
	"GO/model"
	"GO/node"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk menambahkan data Sepatu baru
func Insert() {
	var id, harga int
	var nama string

	reader := bufio.NewReader(os.Stdin)

	// Input data Sepatu
	fmt.Print("Masukkan ID Sepatu: ")
	idStr, _ := reader.ReadString('\n')
	id, _ = strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Masukkan Nama Sepatu: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan Harga Sepatu: ")
	hargaStr, _ := reader.ReadString('\n')
	harga, _ = strconv.Atoi(strings.TrimSpace(hargaStr))

	// Membuat data Sepatu baru
	sepatu := node.Sepatu{
		IDSepatu:    id,
		NamaSepatu:  nama,
		HargaSepatu: harga,
	}

	// Memanggil model untuk menambahkan data Sepatu
	model.CreateSepatu(sepatu)
	fmt.Println("== Data Sepatu berhasil ditambahkan ==")
	fmt.Println()
}

// Fungsi untuk menampilkan semua data Sepatu dan Pembeli
func Views() {
	fmt.Println("== Daftar Sepatu ==")
	for i, sepatu := range model.ReadSepatu() {
		fmt.Println("--- Sepatu ke -", i+1, " ---")
		fmt.Println("ID Sepatu\t: ", sepatu.IDSepatu)
		fmt.Println("Nama Sepatu\t: ", sepatu.NamaSepatu)
		fmt.Println("Harga Sepatu\t: ", sepatu.HargaSepatu)
		fmt.Println()
	}

	fmt.Println("== Daftar Pembeli ==")
	for i, pembeli := range model.ReadPembeli() {
		namaSepatu := model.GetNamaSepatu(pembeli.SepatuID) // Ambil nama sepatu dari model
		fmt.Println("--- Pembeli ke -", i+1, " ---")
		fmt.Println("ID Pembeli\t: ", pembeli.ID)
		fmt.Println("Nama Pembeli\t: ", pembeli.Nama)
		fmt.Println("Sepatu yang dibeli\t: ", namaSepatu)
		fmt.Println()
	}
}

// Fungsi untuk memperbarui data Sepatu
func Update() {
	reader := bufio.NewReader(os.Stdin)
	var id, harga int
	var nama string

	fmt.Print("Masukkan ID Sepatu yang akan diupdate: ")
	idStr, _ := reader.ReadString('\n')
	id, _ = strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Masukkan Nama Sepatu: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan Harga Sepatu: ")
	hargaStr, _ := reader.ReadString('\n')
	harga, _ = strconv.Atoi(strings.TrimSpace(hargaStr))

	// Membuat data Sepatu yang baru
	sepatu := node.Sepatu{
		IDSepatu:    id,
		NamaSepatu:  nama,
		HargaSepatu: harga,
	}

	// Memanggil model untuk mengupdate data Sepatu
	cek := model.UpdateSepatu(sepatu, id)
	if cek {
		fmt.Println("== Data Sepatu berhasil diupdate ==")
	} else {
		fmt.Println("Sepatu gagal diupdate")
	}
	fmt.Println()
}

// Fungsi untuk menghapus data Sepatu
func Delete() {
	reader := bufio.NewReader(os.Stdin)
	var id int

	fmt.Print("Masukkan ID Sepatu yang akan dihapus: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Input ID tidak valid!")
		return
	}

	// Memanggil model untuk menghapus data Sepatu
	cek := model.DeleteSepatu(id)
	if cek {
		fmt.Println("== Data Sepatu berhasil dihapus ==")
	} else {
		fmt.Println("Sepatu gagal dihapus")
	}
	fmt.Println()
}

// Fungsi untuk mencari data Sepatu
func Search() {
	reader := bufio.NewReader(os.Stdin)
	var id int

	fmt.Print("Masukkan ID Sepatu yang akan dicari: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Input ID tidak valid!")
		return
	}

	sepatu, found := model.GetSepatuByID(id)
	if found {
		fmt.Println("== Data Sepatu ditemukan ==")
		fmt.Printf("ID: %d, Nama: %s, Harga: %d\n", sepatu.IDSepatu, sepatu.NamaSepatu, sepatu.HargaSepatu)
	} else {
		fmt.Println("Sepatu tidak ditemukan")
	}
	fmt.Println()
}
