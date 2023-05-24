package view

import (
	"fmt"
	controller "project-uts-mvc/controller/controllerkasir"
)

func InsertDataKasir() {
	var id int
	var nama string
	fmt.Print("Masukkan Id Kasir: ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Kasir: ")
	fmt.Scan(&nama)

	err := controller.ControllerInsertKasir(id, nama)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data Kasir Telah Ditambahkan!")
	}
}

func UpdateDataKasir() {
	var id int
	var nama string
	fmt.Print("Masukkan Id Kasir Terbaru: ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Kasir Terbaru: ")
	fmt.Scan(&nama)

	err := controller.ControllerUpdateKasir(id, nama)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data Kasir Telah Ditambahkan!")
	}
}

func DeleteDataKasir() {
	var id int
	fmt.Println("Masukkan Id Kasir Yang Akan Dihapus: ")
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
		fmt.Println("Id Customer: ", dataKasir.Id)
		fmt.Println("Nama Customer: ", dataKasir.Nama)
	}
}

func ViewByIdDataKasir() {
	var id int
	fmt.Println("Masukkan Id Kasir Yang Ingin Dicari: ")
	fmt.Scan(&id)
	current := controller.ControllerViewByIdKasir(id)
	if current != nil {
		fmt.Println(current.Id)
		fmt.Println(current.Nama)
	} else {
		fmt.Println("<FAILED> Data Dengan Id: ", id, "Tidak Ditemukan!")
	}
}
