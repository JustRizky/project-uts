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
	Nama string
	Nik  int
	Next *Kasir
}

type Customer struct {
	Nama string
	Np   int
	Next *Customer
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

//UPDATE

func (list *LinkedListCS) UpdatenamaCS(oldCS *Customer, newCS string) {
	if list.head == nil {
		fmt.Println("TIDAK ADA BARANG")
	}

	curr := list.head
	for curr != nil {
		if curr == oldCS {
			curr.Nama = newCS
			return
		}
		curr = curr.Next
	}
}

func (list *LinkedListKSR) UpdatenamaKSR(oldKSR *Kasir, newKSR string) {
	if list.head == nil {
		fmt.Println("TIDAK ADA BARANG")
	}

	curr := list.head
	for curr != nil {
		if curr == oldKSR {
			curr.Nama = newKSR
			return
		}
		curr = curr.Next
	}
}

func (list *LinkedListMenu) UpdatenamaMENU(oldMENU *Menu, newMAKAN string, newMINUM string) {
	if list.head == nil {
		fmt.Println("TIDAK ADA BARANG")
	}

	curr := list.head
	for curr != nil {
		if curr == oldMENU {
			curr.Makanan = newMAKAN
			curr.Minuman = newMINUM
			return
		}
		curr = curr.Next
	}
}

//DELETE

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

//VIEWALL

func (list *LinkedListCS) ViewAllCustomer() {
	if list.head == nil {
		fmt.Println("")
	} else {
		tempRestoran := list.head
		fmt.Println("")
		fmt.Println("Customer :")
		for tempRestoran != nil {

			fmt.Println("_________")
			fmt.Println("NP   : ", tempRestoran.Np)
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

			fmt.Println("_________")
			fmt.Println("NIK   : ", tempRestoran.Nik)
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

			fmt.Println("_________")
			fmt.Println("ID   : ", tempRestoran.ID)
			fmt.Println("Makanan : ", tempRestoran.Makanan)
			fmt.Println("Minuman : ", tempRestoran.Minuman)
			tempRestoran = tempRestoran.Next
		}
	}
}

//VIEW BY ID

func (list *LinkedListKSR) viewByIDKSR(id int) {
	curr := list.head
	for curr != nil {
		if curr.Nik == id {
			fmt.Println()
			fmt.Println()
			fmt.Println("VIEW BY NIK KASIR YANG DI TUJU:")
			fmt.Println("NAMA :", curr.Nama, "\n NIK :", curr.Nik)
			return
		}
		curr = curr.Next
	}
	fmt.Println("NIK DENGAN NOMOR ", id, "TIDAK ADA YAA")
}

func (list *LinkedListCS) viewByIDCS(np int) {
	curr := list.head
	for curr != nil {
		if curr.Np == np {
			fmt.Println()
			fmt.Println()
			fmt.Println("VIEW BY NP CUSTOMER YANG DI TUJU:")
			fmt.Println("NAMA :", curr.Nama, "\n NIK :", curr.Np)
			return
		}
		curr = curr.Next
	}
	fmt.Println("NIK DENGAN NOMOR ", np, "TIDAK ADA YAA")
}

func (list *LinkedListMenu) viewByIDMENU(id int) {
	curr := list.head
	for curr != nil {
		if curr.ID == id {
			fmt.Println()
			fmt.Println()
			fmt.Println("VIEW BY ID MENU YANG DI TUJU:")
			fmt.Println("ID   : ", curr.ID, "\nNAMA :", curr.Makanan, "& MINUMAN :", curr.Minuman)
			return
		}
		curr = curr.Next
	}
	fmt.Println("NIK DENGAN NOMOR ", id, "TIDAK ADA YAA")
}

func main() {
	//listCS := &LinkedListCS{}
	listKSR := &LinkedListKSR{}
	//listMenu := &LinkedListMenu{}

	var pilihan1 uint
	var pilihan2 uint
	for pilihan1 != 3 {
		menu1()
		fmt.Print("Masukkan Pilihan: ")
		fmt.Scan(&pilihan1)
		switch pilihan1 {
		//Kasir
		case 1:
			menu2()
			fmt.Print("Masukkan Pilihan (1-6): ")
			fmt.Scan(&pilihan2)
			//Insert
			if pilihan2 == 1 {
				var namaKasir string
				var nikKasir int
				var banyakKasir []int

				fmt.Print("Masukkan Nama: ")
				fmt.Scan(&namaKasir)
				fmt.Print("Masukkan Nik : ")
				fmt.Scan(&nikKasir)

				kasir1 := &Kasir{
					Nama: namaKasir,
					Nik:  nikKasir,
				}
				listKSR.InsertKSR(kasir1)

				kasir2 := &Kasir{
					Nama: namaKasir,
					Nik:  nikKasir,
				}
				listKSR.InsertKSR(kasir2)

			}
			//Update
			if pilihan2 == 2 {
				var namaKasir string
				var nikKasir int

				fmt.Print("Masukkan Nama: ")
				fmt.Scan(&namaKasir)
				fmt.Print("Masukkan Nik : ")
				fmt.Scan(&nikKasir)
				listKSR.UpdatenamaKSR1(kasir2, "RIZKI")

			}
			//Delete

			//View

			//View By Id

		//Customer
		case 2:
			menu2()
			fmt.Print("Masukkan Pilihan (1-6): ")
			fmt.Scan(&pilihan2)

		//Menu
		case 3:
		}
	}

}
