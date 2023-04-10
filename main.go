package main

import (
	"fmt"
)

type Mahasiswa struct {
	ID   int
	Nama string
	NPM  string
	Next *Mahasiswa
}

type LinkedList struct {
	Head *Mahasiswa
}

func menu1() {
	fmt.Println("----------------")
	fmt.Println("1. Mahasiswa")
	fmt.Println("2. Dosen")
	fmt.Println("3. Exit")
}

func menu2() {
	fmt.Println("----------------")
	fmt.Println("1. Insert Depan")
	fmt.Println("2. Insert Belakang")
	fmt.Println("3. Update")
	fmt.Println("4. Delete")
	fmt.Println("5. View")
	fmt.Println("6. View By Id")
	fmt.Println("7. Exit")
	fmt.Println("Masukkan Pilihan (1-7): ")
}

func (list *LinkedList) InsertDepan(mahasiswa *Mahasiswa) {
	if list.Head == nil {
		list.Head = mahasiswa
	} else {
		tempMahasiswa := list.Head
		for tempMahasiswa.Next != nil {
			tempMahasiswa = tempMahasiswa.Next
		}
		tempMahasiswa.Next = mahasiswa
	}
}

func (list *LinkedList) ViewAll() {
	if list.Head == nil {
		fmt.Println("Data Kosong!")
	} else {
		tempMahasiswa := list.Head
		for tempMahasiswa != nil {
			fmt.Println("___________________________")
			fmt.Println("ID   : ", tempMahasiswa.ID)
			fmt.Println("Nama : ", tempMahasiswa.Nama)
			fmt.Println("NPM  : ", tempMahasiswa.NPM)
			tempMahasiswa = tempMahasiswa.Next
		}
	}
}

func main() {
	list := &LinkedList{}

	mahasiswa1 := &Mahasiswa{
		ID:   0,
		Nama: "Rizky",
		NPM:  "01",
	}
	mahasiswa2 := &Mahasiswa{
		ID:   1,
		Nama: "Muhammad",
		NPM:  "02",
	}

	list.InsertDepan(mahasiswa1)
	list.InsertDepan(mahasiswa2)

	list.ViewAll()
}
