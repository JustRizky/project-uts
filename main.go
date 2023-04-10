package main

import "fmt"

type Mahasiswa struct {
	id   int
	name string
	npm  string
}

type Dosen struct {
	Id   int
	Name string
	Nip  string
}

type linkedList struct {
	container  Mahasiswa
	container2 Dosen
	next       *linkedList
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

func view() {

}

func insertDepan() {

}

func insertBelakang(linkedList *linkedList, mahasiswa Mahasiswa) {

}

func delete() {

}

func update() {

}

func main() {

}
