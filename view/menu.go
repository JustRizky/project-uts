package view

import "fmt"

func menuUtama() int {
	var input int
	fmt.Println("----------------")
	fmt.Println("1. Kasir")
	fmt.Println("2. Customer")
	fmt.Println("3. Menu")
	fmt.Println("4. Exit")
	fmt.Print("Masukkan Pilihan: (1-4): ")
	fmt.Scan(&input)
	return input
}

func menuKedua() int {
	var input int
	fmt.Println("----------------")
	fmt.Println("1. Insert")
	fmt.Println("2. Update")
	fmt.Println("3. Delete")
	fmt.Println("4. View")
	fmt.Println("5. View By Id")
	fmt.Println("6. Back")
	fmt.Print("Masukkan Pilihan (1-6): ")
	fmt.Scan(&input)
	return input
}