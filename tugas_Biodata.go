package main

import (
	"fmt"
	"os"
)

type variabel_Teman struct {
	No        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var list_Teman = []variabel_Teman{
	{1, "Maulana", "Jl. Tanjung Gusta", "Network Engineer", "Menambahh Pengetahuan"},
	{2, "Syarfan", "Jl. Amplas", "Freelance", "Suka"},
	{3, "Santo", "Jl. Medan Baru", "Freelane", "Penasaran"},
}

func data_Teman(no int) {
	for _, i := range list_Teman {
		if i.No == no {
			fmt.Println("Nama: ", i.Nama)
			fmt.Println("Alamat: ", i.Alamat)
			fmt.Println("Pekerjaan: ", i.Pekerjaan)
			fmt.Println("Alasan memilih kelas Golang: ", i.Alasan)
			return
		}
	}
	fmt.Println("teman dengan nomor", no, "tidak ditemukan, jumlah teman cuma 3.")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("gunakan go run biodata.go (nomor 1-3)")
		os.Exit(1)
	}
	no := os.Args[1]

	var nomor int
	_, err := fmt.Sscanf(no, "%d", &nomor)
	if err != nil {
		fmt.Println("nomor harus berupa angka")
		os.Exit(1)
	}

	data_Teman(nomor)
}
