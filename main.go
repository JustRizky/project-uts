package main

import (
	"fmt"
)

func menuUtama() int {
	var input int
	fmt.Println("----------------")
	fmt.Println("1. Kasir")
	fmt.Println("2. Customer")
	fmt.Println("3. Menu")
	fmt.Println("4. Exit")
	fmt.Print("Masukkan Pilihan: (1-4): ")
	fmt.Scan(&input)
	return input
}

func menuKedua() int {
	var input int
	fmt.Println("----------------")
	fmt.Println("1. Insert")
	fmt.Println("2. Update")
	fmt.Println("3. Delete")
	fmt.Println("4. View")
	fmt.Println("5. View By Id")
	fmt.Println("6. Back")
	fmt.Print("Masukkan Pilihan (1-6): ")
	fmt.Scan(&input)
	return input
}

type Menu struct {
	Id      int
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

type LinkedListMenu struct {
	head *Menu
}

type LinkedListKasir struct {
	head *Kasir
}

type LinkedListCustomer struct {
	head *Customer
}

// MENU MAKANAN DAN MINUMAN
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

func (list *LinkedListMenu) UpdatenamaMENU(oldMENU *Menu, newMAKAN string, newMINUM string) {
	if list.head == nil {
		fmt.Println("Tidak ada menu!")
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

func (list *LinkedListMenu) DeleteMenu(Id int) {
	if list.head == nil {
		fmt.Println("Menu tidak ditemukan!")
		return
	}

	if list.head.Id == Id {
		list.head = list.head.Next
		return
	}

	prevMenu := list.head
	for prevMenu.Next != nil {
		if prevMenu.Next.Id == Id {
			prevMenu.Next = prevMenu.Next.Next
			return
		}
		prevMenu = prevMenu.Next
	}

	fmt.Println("Menu tidak ditemukan!")
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
			fmt.Println("ID   : ", tempRestoran.Id)
			fmt.Println("Makanan : ", tempRestoran.Makanan)
			fmt.Println("Minuman : ", tempRestoran.Minuman)
			tempRestoran = tempRestoran.Next
		}
	}
}

func (list *LinkedListMenu) viewByIDMENU(Id int) {
	curr := list.head
	for curr != nil {
		if curr.Id == Id {
			fmt.Println()
			fmt.Println()
			fmt.Println("VIEW BY ID MENU YANG DITUJU:")
			fmt.Println("ID   : ", curr.Id, "\nNAMA :", curr.Makanan, "& MINUMAN :", curr.Minuman)
			return
		}
		curr = curr.Next
	}
	fmt.Println("NIK DENGAN NOMOR ", Id, "TIDAK ADA!")
}

// Kasir
func (list *LinkedListKasir) InsertKSR(ksr *Kasir) {
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

func (list *LinkedListKasir) UpdatenamaKSR(oldKSR *Kasir, newKSR string) {
	if list.head == nil {
		fmt.Println("Tidak ada kasir!")
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

func (list *LinkedListKasir) DeleteKSR(ID int) {
	if list.head == nil {
		fmt.Println("Kasir tidak ditemukan!")
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

	fmt.Println("Kasir tidak ditemukan!")
}

func (list *LinkedListKasir) ViewAllKasir() {
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

func (list *LinkedListKasir) viewByIDKSR(id int) {
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
	fmt.Println("NIK DENGAN NOMOR ", id, "TIDAK ADA!")
}

// Customer
func (list *LinkedListCustomer) InsertCS(cs *Customer) {
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

func (list *LinkedListCustomer) UpdatenamaCS(oldCS *Customer, newCS string) {
	if list.head == nil {
		fmt.Println("Tidak ada customer!")
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

func (list *LinkedListCustomer) DeleteCS(ID int) {
	if list.head == nil {
		fmt.Println("Customer tidak ditemukan!")
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

	fmt.Println("Customer tidak ditemukan!")
}

func (list *LinkedListCustomer) ViewAllCustomer() {
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

func (list *LinkedListCustomer) viewByIDCS(np int) {
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
	fmt.Println("NIK DENGAN NOMOR ", np, "TIDAK ADA!")
}

func main() {
	listKasir := &LinkedListKasir{}
	ListCustomer := &LinkedListCustomer{}
	listMenu := &LinkedListMenu{}

menuUtama:
	var pilihanMenuUtama int = menuUtama()
	for pilihanMenuUtama != 4 {
		switch pilihanMenuUtama {
		case 1: //Kasir
			var pilihanMenuKedua int = menuKedua()
			for pilihanMenuKedua != 7 {
				switch pilihanMenuKedua {
				case 1: //Insert
					break
				case 2: //Update
					break
				case 3: //Delete
					break
				case 4: // View
					listKasir.ViewAllKasir()
					break
				case 5: // View By Id
					break
				case 6: // Back
					goto menuUtama
				}
			}

			break
		case 2: //Customer
			var pilihanMenuKedua int = menuKedua()
			for pilihanMenuKedua != 7 {
				switch pilihanMenuKedua {
				case 1: //Insert
					break
				case 2: //Update
					break
				case 3: //Delete
					break
				case 4: // View
					ListCustomer.ViewAllCustomer()
					break
				case 5: // View By Id
					break
				case 6: // Back
					goto menuUtama
				}
			}
			break

		case 3: //Menu Makanan Dan Minuman
			var pilihanMenuKedua int = menuKedua()
			for pilihanMenuKedua != 7 {
				switch pilihanMenuKedua {
				case 1: //Insert
					break
				case 2: //Update
					break
				case 3: //Delete
					break
				case 4: // View
					listMenu.ViewAllMenu()
					break
				case 5: // View By Id
					break
				case 6: // Back
					goto menuUtama
				}
			}
			break

		default:
			fmt.Println("Pilihan tidak valid!")
			break
		}

	menuUlangUtama:
		var ulangMenuUtama string
		fmt.Print("Apakah Anda ingin mengulang (Y/t): ")
		fmt.Scan(&ulangMenuUtama)
		if ulangMenuUtama == "y" || ulangMenuUtama == "Y" {
			pilihanMenuUtama = menuUtama()
		} else if ulangMenuUtama == "t" || ulangMenuUtama == "T" {
			break
		} else {
			goto menuUlangUtama
		}
	}

	fmt.Println("Akhir Program!")

}
