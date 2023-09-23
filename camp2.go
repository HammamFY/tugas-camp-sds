package main

import (
	"fmt"
)

/*

Soal:
1. Buatlah algoritma pengurutan bilangan dari yang terbesar sampai yang terkecil!
2. Buatlah slice dengan jumlah data di dalam slice yaitu ada 5 buah dan tambahkan data ke dalam slice sebanyak 3!

*/

// 1. Buatlah algoritma pengurutan bilangan dari yang terbesar sampai yang terkecil!
var angka = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Println(arr)
			fmt.Println("Element yang di tukar yaitu ", arr[j], "dengan ", arr[j+1])
			if arr[j] < arr[j+1] {
				// tukar elemen
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				fmt.Println("tidak ada penukaran")
			}
		}
		fmt.Println("jumlah Iterasi ke ", i+1)

	}
}

func main() {

	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Data Sebelum Diurutkan:", data)
	bubbleSort(data)
	fmt.Println("Data Setelah Diurutkan:", data)

	// 2. Buatlah slice dengan jumlah data di dalam slice yaitu ada 5 buah dan tambahkan data ke dalam slice sebanyak 3!
	var nomor = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("\ndata keseluruhan adalah:", nomor)
	// Slice
	// array[low:high]
	slice := nomor[0:5]
	fmt.Println("nilai dari slice adalah:", slice)
	//menambah data slice
	databaru := append(slice, 5, 6, 7)
	fmt.Println("slice yang sudah ditambahkan :", databaru)

}
