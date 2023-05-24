package view

import (
	"fmt"
	controller "project-uts-mvc/controller/controllermenurestoran"
)

func InsertDataMenuRestoran() {
	var id int
	var makanan, minuman string
	fmt.Print("Masukkan Id Menu: ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Makanan: ")
	fmt.Scanln(&makanan)
	fmt.Print("Masukkan Nama Minuman: ")
	fmt.Scanln(&minuman)

	err := controller.ControllerInsertMenuRestoran(id, makanan, minuman)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data Menu Telah Ditambahkan!")
	}
}

func UpdateDataMenuRestoran() {
	var id int
	var makanan, minuman string
	fmt.Print("Masukkan Id Menu Terbaru: ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Makanan Terbaru: ")
	fmt.Scanln(&makanan)
	fmt.Print("Masukkan Nama Minuman Terbaru: ")
	fmt.Scanln(&minuman)

	err := controller.ControllerUpdateMenuRestoran(id, makanan, minuman)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data Menu Telah Ditambahkan!")
	}
}

func DeleteDataMenuRestoran() {
	var id int
	fmt.Print("Masukkan Id Menu Yang Akan Dihapus: ")
	fmt.Scan(&id)
	err := controller.ControllerDeleteMenuRestoran(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data Menu Dengan Id: ", id, "Telah Dihapus!")
	}
}

func ViewAllDataMenuRestoran() {
	menuRestoran := controller.ControllerViewAllMenuRestoran()
	for _, dataMenuRestoran := range menuRestoran {
		fmt.Println("__________________________________")
		fmt.Println("Id Menu Restoran: ", dataMenuRestoran.Id)
		fmt.Println("Makanan Customer: ", dataMenuRestoran.Makanan)
		fmt.Println("Minuman Customer: ", dataMenuRestoran.Minuman)
		fmt.Println("__________________________________")
	}
}

func ViewByIdDataMenuRestoran() {
	var id int
	fmt.Print("Masukkan Id Menu Yang Ingin Dicari: ")
	fmt.Scan(&id)
	current := controller.ControllerViewByIdMenuRestoran(id)
	if current != nil {
		fmt.Println("__________________________________")
		fmt.Println("Id Menu Restoran: ", current.Id)
		fmt.Println("Makanan Customer: ", current.Makanan)
		fmt.Println("Minuman Customer: ", current.Minuman)
		fmt.Println("__________________________________")
	} else {
		fmt.Println("<FAILED> Data Dengan Id: ", id, "Tidak Ditemukan!")
	}
}
