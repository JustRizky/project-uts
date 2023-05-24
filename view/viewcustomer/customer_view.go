package view

import (
	"fmt"
	controller "project-uts-mvc/controller/controllercustomer"
)

func InsertDataCustomer() {
	var id int
	var nama string
	fmt.Print("Masukkan Id Customer: ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Customer: ")
	fmt.Scan(&nama)

	err := controller.ControllerInsertCustomer(id, nama)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data Customer Telah Ditambahkan!")
	}
}

func UpdateDataCustomer() {
	var id int
	var nama string
	fmt.Print("Masukkan Id Customer Terbaru: ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Customer Terbaru: ")
	fmt.Scan(&nama)

	err := controller.ControllerUpdateCustomer(id, nama)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data Customer Telah Ditambahkan!")
	}
}

func DeleteDataCustomer() {
	var id int
	fmt.Println("Masukkan Id Customer Yang Akan Dihapus: ")
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
		fmt.Println("Id Customer: ", dataCustomer.Id)
		fmt.Println("Nama Customer: ", dataCustomer.Nama)
	}
}

func ViewByIdDataCustomer() {
	var id int
	fmt.Println("Masukkan Id Customer Yang Ingin Dicari: ")
	fmt.Scan(&id)
	current := controller.ControllerViewByIdCustomer(id)
	if current != nil {
		fmt.Println(current.Id)
		fmt.Println(current.Nama)
	} else {
		fmt.Println("<FAILED> Data Dengan Id: ", id, "Tidak Ditemukan!")
	}
}
