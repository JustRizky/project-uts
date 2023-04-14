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

type LinkedListMenu struct {
	head *Menu
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

func (list *LinkedListKSR) InsertKSR(ksr *Kasir) {
	if list.head == nil {
		list.head = ksr
	} else {
		tempksr := list.head
		for tempksr.Next != nil {
			tempksr = tempksr.Next
		}
		tempksr.Next = ksr
	}
}

func (list *LinkedListCS) InsertCS(cs *Customer) {
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

func (list *LinkedListMenu) InsertMenu(menu *Menu) {
	if list.head == nil {
		list.head = menu
	} else {
		tempMenu := list.head
		for tempMenu.Next != nil {
			tempMenu = tempMenu.Next
		}
		tempMenu.Next = menu
	}

}

func (list *LinkedListKSR) DeleteKSR(ID int) {
	if list.head == nil {
		fmt.Println("Kasir tidak ditemukan")
		return
	}

	if list.head.Nik == ID {
		list.head = list.head.Next
		return
	}

	prevKasir := list.head
	for prevKasir.Next != nil {
		if prevKasir.Next.Nik == ID {
			prevKasir.Next = prevKasir.Next.Next
			return
		}
		prevKasir = prevKasir.Next
	}

	fmt.Println("Kasir tidak ditemukan")
}

func (list *LinkedListCS) DeleteCS(ID int) {
	if list.head == nil {
		fmt.Println("Customer tidak ditemukan")
		return
	}

	if list.head.Np == ID {
		list.head = list.head.Next
		return
	}

	prevCustomer := list.head
	for prevCustomer.Next != nil {
		if prevCustomer.Next.Np == ID {
			prevCustomer.Next = prevCustomer.Next.Next
			return
		}
		prevCustomer = prevCustomer.Next
	}

	fmt.Println("Customer tidak ditemukan")
}

func (list *LinkedListMenu) DeleteMenu(ID int) {
	if list.head == nil {
		fmt.Println("Menu tidak ditemukan")
		return
	}

	if list.head.ID == ID {
		list.head = list.head.Next
		return
	}

	prevMenu := list.head
	for prevMenu.Next != nil {
		if prevMenu.Next.ID == ID {
			prevMenu.Next = prevMenu.Next.Next
			return
		}
		prevMenu = prevMenu.Next
	}

	fmt.Println("Menu tidak ditemukan")
}



func (list *LinkedListCS) ViewAllCustomer() {
	if list.head == nil {
		fmt.Println("")
	} else {
		tempRestoran := list.head
		fmt.Println("")
		fmt.Println("Customer :")
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
		fmt.Println("")
	} else {
		tempRestoran := list.head
		fmt.Println("")
		fmt.Println("Kasir :")
		for tempRestoran != nil {

			fmt.Println("___________________________")
			fmt.Println("ID   : ", tempRestoran.Nik)
			fmt.Println("Nama : ", tempRestoran.Nama)
			tempRestoran = tempRestoran.Next
		}
	}
}

func (list *LinkedListMenu) ViewAllMenu() {
	if list.head == nil {
		fmt.Println("")
	} else {
		tempRestoran := list.head
		fmt.Println("")
		fmt.Println("Makanan :")
		for tempRestoran != nil {

			fmt.Println("___________________________")
			fmt.Println("ID   : ", tempRestoran.ID)
			fmt.Println("Makanan : ", tempRestoran.Makanan)
			fmt.Println("Minuman : ", tempRestoran.Minuman)
			tempRestoran = tempRestoran.Next
		}
	}
}

func main() {
	list2 := &LinkedListCS{}
	list3 := &LinkedListKSR{}
	list4 := &LinkedListMenu{}

	kasir1 := &Kasir{
		Nama: "rizky",
		Nik:  01,
	}
	kasir2 := &Kasir{
		Nama: "fikri",
		Nik:  02,
	}

	customer1 := &Customer{
		Nama: "rangga",
		Np:   99,
	}
	customer2 := &Customer{
		Nama: "Pandu",
		Np:   100,
	}

	kasir3 := &Kasir{
		Nama: "Panji",
		Nik:  03,
	}

	menu1 := &Menu{
		ID:      01,
		Makanan: "Nasi Ayam",
		Minuman: "Jus",
	}

	menu2 := &Menu{
        ID:      02,
        Makanan: "Tahu Tek",
        Minuman: "Es Kepal",
	    }

	list2.InsertCS(customer1)
	list2.InsertCS(customer2)	
	list3.InsertKSR(kasir1)
	list3.InsertKSR(kasir2)
	list3.InsertKSR(kasir3)
	list4.InsertMenu(menu1)
	list4.InsertMenu(menu2)


	list3.DeleteKSR(01)
	list4.DeleteMenu(01)
	list2.DeleteCS(99)

	list3.ViewAllKasir()
	list2.ViewAllCustomer()
	list4.ViewAllMenu()
}