package main

import (
	"fmt"
	"os"
	"strings"
)

type Biodata struct {
	no_absen  string
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	var args = os.Args
	if len(args) > 1 {
		var students = []Biodata{
			{no_absen: "1", nama: "Agus", alamat: "Jakarta", pekerjaan: "Web Programmer", alasan: "Belajar ilmu baru"},
			{no_absen: "2", nama: "Ikbal", alamat: "Bandung", pekerjaan: "IT Support", alasan: "Tuntutan perusahaan"},
			{no_absen: "3", nama: "Tony", alamat: "Tasikmalaya", pekerjaan: "Fullstack Developer", alasan: "Mengikuti perkembangan zaman"},
		}

		isExist := false
		var foundStudent *Biodata
		for _, student := range students {
			if student.no_absen == args[1] {
				isExist = true
				foundStudent = &student
				break
			}
		}
		if isExist {
			showData(foundStudent)
		} else {
			fmt.Println("Data Tidak Ada...")
		}
	} else {
		fmt.Printf("Format Salah... \nExample : go run biodata.go 1")
	}
}

func showData(student *Biodata) {
	fmt.Println(strings.Repeat("=", 30))
	fmt.Printf(" No. Absen : %s\n Nama : %s\n Alamat : %s\n Pekerjaan : %s\n Alasan memilih kelas golang : %s\n", student.no_absen, student.nama, student.alamat, student.pekerjaan, student.alasan)
	fmt.Println(strings.Repeat("=", 30))
}
