package view

import (
	"fmt"
	controller "project-uts-mvc/controller/controllercustomer"
)

func InsertDataCustomer() {
	var id int
	var nama string
	var nomormeja int
	fmt.Print("Masukkan Id Customer: ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Customer: ")
	fmt.Scanln(&nama)
	fmt.Println("Masukkan Nomor Meja Customer: ")
	fmt.Scan(&nomormeja)

	err := controller.ControllerInsertCustomer(id, nomormeja, nama)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data Customer Telah Ditambahkan!")
	}
}

func UpdateDataCustomer() {
	var id int
	var nama string
	var nomormeja int
	fmt.Print("Masukkan Id Customer Terbaru: ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Customer Terbaru: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan Nomor Meja Customer Terbaru: ")
	fmt.Scan(&nomormeja)

	err := controller.ControllerUpdateCustomer(id, nomormeja, nama)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data Customer Telah Ditambahkan!")
	}
}

func DeleteDataCustomer() {
	var id int
	fmt.Print("Masukkan Id Customer Yang Akan Dihapus: ")
	fmt.Scan(&id)
	err := controller.ControllerDeleteCustomer(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data Dengan Id: ", id, "Customer Telah Dihapus!")
	}
}

func ViewAllDataCustomer() {
	customer := controller.ControllerViewAllCustomer()
	for _, dataCustomer := range customer {
		fmt.Println("__________________________________")
		fmt.Println("Id Customer: ", dataCustomer.Id)
		fmt.Println("Nama Customer: ", dataCustomer.Nama)
		fmt.Println("Nomor Meja Customer: ", dataCustomer.NomorMeja)
		fmt.Println("__________________________________")
	}
}

func ViewByIdDataCustomer() {
	var id int
	fmt.Print("Masukkan Id Customer Yang Ingin Dicari: ")
	fmt.Scan(&id)
	current := controller.ControllerViewByIdCustomer(id)
	if current != nil {
		fmt.Println("__________________________________")
		fmt.Println("Id Customer: ", current.Id)
		fmt.Println("Nama Customer: ", current.Nama)
		fmt.Println("Nomor Meja Customer: ", current.NomorMeja)
		fmt.Println("__________________________________")
	} else {
		fmt.Println("<FAILED> Data Dengan Id: ", id, "Tidak Ditemukan!")
	}
}
