package main

import (
	"fmt"
)

type Menu struct {
	ID      int
	Makanan string
	Minuman string
	Next    *Menu
}

type Kasir struct {
	ID   int
	Nama string
	Nik  int
	Next *Kasir
}

type Customer struct {
	ID   int
	Nama string
	Np   int
	Next *Customer
}

type LinkedList struct {
	Head *Menu
}

type LinkedListKSR struct {
	head *Kasir
}

type LinkedListCS struct {
	head *Customer
}

func menu1() {
	fmt.Println("----------------")
	fmt.Println("1. Kasir")
	fmt.Println("2. Customer")
	fmt.Println("3. Exit")
}

func menu2() {
	fmt.Println("----------------")
	fmt.Println("1. Insert")
	fmt.Println("2. Update")
	fmt.Println("3. Delete")
	fmt.Println("4. View")
	fmt.Println("5. View By Id")
	fmt.Println("6. Exit")
	fmt.Println("Masukkan Pilihan (1-6): ")
}

func (list *LinkedListKSR) Insert1(ksr *Kasir) {
	if list.head == nil {
		list.head = ksr
	} else {
		tempksr := list.head
		for tempksr.Next != nil {
			tempksr = tempksr.Next
		}
	}
}

func (list *LinkedListCS) Insert2(cs *Customer) {
	if list.head == nil {
		list.head = cs
	} else {
		tempCs := list.head
		for tempCs.Next != nil {
			tempCs = tempCs.Next
		}
		tempCs.Next = cs
	}
}

func (list *LinkedListCS) ViewAllCustomer() {
	if list.head == nil {
		fmt.Println("Data Kosong!")
	} else {
		tempRestoran := list.head
		for tempRestoran != nil {
			fmt.Println("___________________________")
			fmt.Println("ID   : ", tempRestoran.Np)
			fmt.Println("Nama : ", tempRestoran.Nama)
			tempRestoran = tempRestoran.Next
		}
	}
}

func (list *LinkedListKSR) ViewAllKasir() {
	if list.head == nil {
		fmt.Println("Data Kosong!")
	} else {
		tempRestoran := list.head
		for tempRestoran != nil {
			fmt.Println("___________________________")
			fmt.Println("ID   : ", tempRestoran.Nik)
			fmt.Println("Nama : ", tempRestoran.Nama)
			tempRestoran = tempRestoran.Next
		}
	}
}

func main() {
	list2 := &LinkedListCS{}
	list3 := &LinkedListKSR{}

	customer1 := &Customer{
		Nama: "Pikri",
		Np:   456,
	}
	kasir := &Kasir{
		Nama: "aku",
		Nik:  07606,
	}

	list3.Insert1(kasir)
	list2.Insert2(customer1)
	list3.ViewAllKasir()
	list2.ViewAllCustomer()
}
