package main

import "fmt"

type biodata struct {
	name      string
	alamat    string
	pekerjaan string
	alasan    string
}

func AbsenSiswa() {
	var absen int
	fmt.Print("Masukkan Absen Anda : ")
	fmt.Scanln(&absen)
	switch {
	case absen == 1:
		var d1 = biodata{" Name : Putri\n", "Address : Jogja\n", "Profession : Mahasiswa\n", "Reason : Karena Ingin belajar bahasa baru\n"}
		fmt.Println(d1.name, d1.alamat, d1.pekerjaan, d1.alasan)

	case absen == 2:
		var d1 = biodata{" Name : Adetya\n", "Address : Kalimantan\n", "Profession : Pengajar\n", "Reason : Karena Ingin menambah ilmu baru\n"}
		fmt.Println(d1.name, d1.alamat, d1.pekerjaan, d1.alasan)

	case absen == 3:
		var d1 = biodata{" Name : Azhari\n", "Address : Jakarta\n", "Profession : Programmer\n", "Reason : Karena Ingin mendalami bahasa pemprograman\n"}
		fmt.Println(d1.name, d1.alamat, d1.pekerjaan, d1.alasan)
	default:
		fmt.Println("No Absen tidak ada")
	}
}

func main() {
	AbsenSiswa()
}
