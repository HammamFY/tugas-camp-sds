package main

import (
	"fmt"
)

/*
Soal:
1. Buatlah beberapa data menggunakan MAP
2. Buatlah beberapa struct dengan methodnya dibebaskan
3. Buatlah manipulasi data di struct dan pointer
*/

// Struct
type Mahasiswa1 struct {
	Nama, Jurusan, Fakultas string
	Umur                    int
}

func main() {
	//Map
	person := map[string]string{
		"Name":      "Hammam",
		"Education": "S1",
		"Adress":    "Kebumen",
		"Age":       "21",
	}
	fmt.Println(person)
	fmt.Println("Jumlah data pada Map person:", len(person))
	fmt.Println("Nama saya adalah", person["Name"], "Umurku adalah", person["Age"])

	//Penerapan Struct
	Fajar := Mahasiswa1{
		Nama:     "Fajar",
		Jurusan:  "Peternakan",
		Fakultas: "Peternakan",
		Umur:     22,
	}

	fmt.Println(Mahasiswa1{})
	fmt.Println(Fajar)

	//Manipulasi Struct
	Rizki := Mahasiswa1{
		Nama:     "Rizki",
		Jurusan:  "Peternakan",
		Fakultas: "Peternakan",
		Umur:     21,
	}

	fmt.Println(Rizki)

	//Pointer
	var Alamat = "Purwokerto Timur"
	fmt.Println(Alamat)
	var Alamat1 *string = &Alamat

	//Manipulasi Pointer
	Alamat = "Purwokerto Barat"
	fmt.Println(Alamat)
	fmt.Print(*Alamat1)

}
