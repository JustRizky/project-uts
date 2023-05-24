package view

import (
	"fmt"
	controller "project-uts-mvc/controller/controllerkasir"
)

func InsertDataKasir() {
	var id int
	var nama string
	var umur int
	fmt.Print("Masukkan Id Kasir: ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Kasir: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan Umur Kasir: ")
	fmt.Scan(&umur)

	err := controller.ControllerInsertKasir(id, umur, nama)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data Kasir Telah Ditambahkan!")
	}
}

func UpdateDataKasir() {
	var id int
	var nama string
	var umur int
	fmt.Print("Masukkan Id Kasir Terbaru: ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Kasir Terbaru: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan Umur Kasir Terbaru: ")
	fmt.Scan(&umur)

	err := controller.ControllerUpdateKasir(id, umur, nama)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data Kasir Telah Ditambahkan!")
	}
}

func DeleteDataKasir() {
	var id int
	fmt.Print("Masukkan Id Kasir Yang Akan Dihapus: ")
	fmt.Scan(&id)
	err := controller.ControllerDeleteKasir(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data Dengan Id: ", id, "Kasir Telah Dihapus!")
	}
}

func ViewAllDataKasir() {
	kasir := controller.ControllerViewAllKasir()
	for _, dataKasir := range kasir {
		fmt.Println("__________________________________")
		fmt.Println("Id Kasir: 	 ", dataKasir.Id)
		fmt.Println("Nama Kasir: ", dataKasir.Nama)
		fmt.Println("__________________________________")
	}
}

func ViewByIdDataKasir() {
	var id int
	fmt.Print("Masukkan Id Kasir Yang Ingin Dicari: ")
	fmt.Scan(&id)
	current := controller.ControllerViewByIdKasir(id)
	if current != nil {
		fmt.Println("__________________________________")
		fmt.Println(current.Id)
		fmt.Println(current.Nama)
		fmt.Println("__________________________________")
	} else {
		fmt.Println("<FAILED> Data Dengan Id: ", id, "Tidak Ditemukan!")
	}
}
