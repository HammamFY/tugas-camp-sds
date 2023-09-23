package main

import (
	"fmt"
)

/*
Soal:
1. Buatlah fungsi untuk :
a. menghitung luas persegi
b. menghitung luas segitiga
c. menghitung luas lingkaran
d. menghitung energi potensial, kinetik
e. menghitung frekuensi atau getaran
f. konversi untuk semua satuan suhu
*/

// a. menghitung luas persegi
func LuasPersegi() {
	var p float32 = 8
	var l float32 = 3

	LPersegi := p * l

	fmt.Println("Luas persegi =", LPersegi)
}

// b. menghitung luas segitiga
func LuasSegitiga() {
	var a float32 = 8
	var t float32 = 3

	LSegitiga := 0.5 * a * t

	fmt.Println("Luas segitiga =", LSegitiga)
}

// c. menghitung luas lingkaran
func LuasLingkaran() {
	var r float32 = 10

	LLingkaran := 3.14 * r * r

	fmt.Println("Luas lingkaran =", LLingkaran)
}

// d. menghitung energi potensial, kinetik
func EnergiPotensial() {
	var m float32 = 5
	var g float32 = 9.8
	var h float32 = 5

	EP := m * g * h

	fmt.Println("Nilai energi potensial =", EP)
}
func EnergiKinetik() {
	var m float32 = 10
	var v float32 = 2

	EK := 0.5 * m * v * v

	fmt.Println("Nilai energi kinetik =", EK)
}

// e. menghitung frekuensi atau getaran
func Frekuensi() {
	var T float32 = 10
	f := 1 / T

	fmt.Println("Nilai frekuensi =", f)
}

// f. konversi untuk semua satuan suhu
// diketahui celcius
func Celcius() {
	var celcius float64 = -273.15
	reamur := (4 / 5) * celcius
	fahrenheit := ((9 / 5) * celcius) + 32
	kelvin := celcius + 273.15
	fmt.Println("Hasil konversi suhu dari celcius ke reamur =", reamur)
	fmt.Println("Hasil konversi suhu dari celcius ke fahrenheit =", fahrenheit)
	fmt.Println("Hasil konversi suhu dari celcius ke kelvin =", kelvin)
}

// diketahui reamur
func Reamur() {
	var reamur float64 = -0
	celcius := (5 / 4) * reamur
	fahrenheit := ((9 / 4) * reamur) + 32
	kelvin := ((5 / 4) * reamur) + 273.15
	fmt.Println("Hasil konversi suhu dari reamur ke celcius =", celcius)
	fmt.Println("Hasil konversi suhu dari reamur ke kelvin =", fahrenheit)
	fmt.Println("Hasil konversi suhu dari reamur ke kelvin =", kelvin)
}

// diketahui fahrenheit
func Fahrenheit() {
	var fahrenheit float64 = -241.14999
	celcius := (5 / 9) * (fahrenheit - 32)
	reamur := (4 / 9) * (fahrenheit - 32)
	kelvin := ((fahrenheit - 32) * (5 / 9)) + 273.15
	fmt.Println("Hasil konversi suhu dari fahrenheit ke celcius =", celcius)
	fmt.Println("Hasil konversi suhu dari fahrenheit ke reamur =", reamur)
	fmt.Println("Hasil konversi suhu dari fahrenheit ke kelvin =", kelvin)
}

// diketahui kelvin
func Kelvin() {
	var kelvin float64 = 0
	celcius := kelvin - 273.15
	reamur := (4 / 5) * (kelvin - 273.15)
	fahrenheit := ((kelvin - 273.15) * (9 / 5)) + 32
	fmt.Println("Hasil konversi suhu dari kelvin ke celcius =", celcius)
	fmt.Println("Hasil konversi suhu dari kelvin ke reamur =", reamur)
	fmt.Println("Hasil konversi suhu dari kelvin ke fahrenheit =", fahrenheit)
}

func main() {
	// a. menghitung luas persegi
	LuasPersegi()
	// b. menghitung luas segitiga
	LuasSegitiga()
	// c. menghitung luas lingkaran
	LuasLingkaran()
	// d. menghitung energi potensial, kinetik
	// energi potensial
	EnergiPotensial()
	// energi kinetik
	EnergiKinetik()
	// e. menghitung frekuensi atau getaran
	Frekuensi()
	// f. konversi untuk semua satuan suhu
	// diketahui celcius
	Celcius()
	// diketahui reamur
	Reamur()
	// diketahui fahrenheit
	Fahrenheit()
	// diketahui kelvin
	Kelvin()
}
