package main

import (
	"fmt"
	"os"
	"os/exec"
)

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

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
	Id   int
	Nama string
	Next *Kasir
}

type Customer struct {
	Id   int
	Nama string
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

func (list *LinkedListMenu) UpdatenamaMENU(Id int, newMAKAN string, newMINUM string) {
	if list.head == nil {
		fmt.Println("Tidak ada menu!")
	}

	curr := list.head
	for curr != nil {
		if curr.Id == Id {
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
	fmt.Println("ID DENGAN NOMOR ", Id, "TIDAK ADA!")
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

func (list *LinkedListKasir) UpdatenamaKSR(Id int, newKSR string) {
	if list.head == nil {
		fmt.Println("Tidak ada kasir!")
	}

	curr := list.head
	for curr != nil {
		if curr.Id == Id {
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

	if list.head.Id == ID {
		list.head = list.head.Next
		return
	}

	prevKasir := list.head
	for prevKasir.Next != nil {
		if prevKasir.Next.Id == ID {
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
			fmt.Println("ID   : ", tempRestoran.Id)
			fmt.Println("Nama : ", tempRestoran.Nama)
			tempRestoran = tempRestoran.Next
		}
	}
}

func (list *LinkedListKasir) viewByIDKSR(id int) {
	curr := list.head
	for curr != nil {
		if curr.Id == id {
			fmt.Println()
			fmt.Println()
			fmt.Println("VIEW BY NIK KASIR YANG DI TUJU:")
			fmt.Println("NAMA :", curr.Nama, "\n ID :", curr.Id)
			return
		}
		curr = curr.Next
	}
	fmt.Println("ID DENGAN NOMOR ", id, "TIDAK ADA!")
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

func (list *LinkedListCustomer) UpdatenamaCS(Id int, newCS string) {
	if list.head == nil {
		fmt.Println("Tidak ada customer!")
	}

	curr := list.head
	for curr != nil {
		if curr.Id == Id {
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

	if list.head.Id == ID {
		list.head = list.head.Next
		return
	}

	prevCustomer := list.head
	for prevCustomer.Next != nil {
		if prevCustomer.Next.Id == ID {
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
			fmt.Println("ID   : ", tempRestoran.Id)
			fmt.Println("Nama : ", tempRestoran.Nama)
			tempRestoran = tempRestoran.Next
		}
	}
}

func (list *LinkedListCustomer) viewByIDCS(id int) {
	curr := list.head
	for curr != nil {
		if curr.Id == id {
			fmt.Println()
			fmt.Println()
			fmt.Println("VIEW BY NP CUSTOMER YANG DI TUJU:")
			fmt.Println("NAMA :", curr.Nama, "\n ID :", curr.Id)
			return
		}
		curr = curr.Next
	}
	fmt.Println("ID DENGAN NOMOR ", id, "TIDAK ADA!")
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
			clear()
		menuSecondary1:
			var pilihanMenuKedua int = menuKedua()
			var namaKasir string
			var namaBaruKasir string
			var idKasir int
			var jumlahKasir int

			if pilihanMenuKedua == 1 {
				fmt.Print("Masukkan Jumlah Kasir: ")
				fmt.Scan(&jumlahKasir)
				for i := 1; i <= jumlahKasir; i++ {
					fmt.Print("Masukkan Nama Kasir: ")
					fmt.Scan(&namaKasir)
					fmt.Print("Masukkan ID Kasir: ")
					fmt.Scan(&idKasir)

					kasir := &Kasir{
						Id:   idKasir,
						Nama: namaKasir,
					}
					listKasir.InsertKSR(kasir)
				}
				goto menuSecondary1
			} else if pilihanMenuKedua == 2 {
				fmt.Print("Masukkan ID yang akan diupdate: ")
				fmt.Scan(&idKasir)
				fmt.Print("Masukkan Nama yang baru: ")
				fmt.Scan(&namaBaruKasir)
				listKasir.UpdatenamaKSR(idKasir, namaBaruKasir)
				goto menuSecondary1
			} else if pilihanMenuKedua == 3 {
				fmt.Print("Masukkan ID yang akan Dihapus: ")
				fmt.Scan(&idKasir)
				listKasir.DeleteKSR(idKasir)
				goto menuSecondary1
			} else if pilihanMenuKedua == 4 {
				listKasir.ViewAllKasir()
				goto menuSecondary1
			} else if pilihanMenuKedua == 5 {
				fmt.Print("Masukkan ID yang Dicari: ")
				fmt.Scan(&idKasir)
				listKasir.viewByIDKSR(idKasir)
				goto menuSecondary1
			} else if pilihanMenuKedua == 6 {
				goto menuUtama
			}
			break
		case 2: //Customer
			clear()
		menuSecondary2:
			var pilihanMenuKedua int = menuKedua()
			var namaCustomer string
			var namaBaruCustomer string
			var idCustomer int
			var jumlahCustomer int

			if pilihanMenuKedua == 1 {
				fmt.Print("Masukkan Jumlah Customer: ")
				fmt.Scan(&jumlahCustomer)
				for i := 1; i <= jumlahCustomer; i++ {
					fmt.Print("Masukkan Nama Customer: ")
					fmt.Scan(&namaCustomer)
					fmt.Print("Masukkan ID Customer: ")
					fmt.Scan(&idCustomer)

					customer := &Customer{
						Id:   idCustomer,
						Nama: namaCustomer,
					}
					ListCustomer.InsertCS(customer)
				}
				goto menuSecondary2
			} else if pilihanMenuKedua == 2 {
				fmt.Print("Masukkan ID yang akan diupdate: ")
				fmt.Scan(&idCustomer)
				fmt.Print("Masukkan Nama yang baru: ")
				fmt.Scan(&namaBaruCustomer)
				ListCustomer.UpdatenamaCS(idCustomer, namaBaruCustomer)
				goto menuSecondary2
			} else if pilihanMenuKedua == 3 {
				fmt.Print("Masukkan ID yang akan Dihapus: ")
				fmt.Scan(&idCustomer)
				ListCustomer.DeleteCS(idCustomer)
				goto menuSecondary2
			} else if pilihanMenuKedua == 4 {
				ListCustomer.ViewAllCustomer()
				goto menuSecondary2
			} else if pilihanMenuKedua == 5 {
				fmt.Print("Masukkan ID yang Dicari: ")
				fmt.Scan(&idCustomer)
				ListCustomer.viewByIDCS(idCustomer)
				goto menuSecondary2
			} else if pilihanMenuKedua == 6 {
				goto menuUtama
			}
			break
		case 3: //Menu Makanan Dan Minuman
			clear()
		menuSecondary3:
			var pilihanMenuKedua int = menuKedua()
			var namaMakanan string
			var namaMinuman string
			var namaBaruMakanan string
			var namaBaruMinuman string
			var idMenu int
			var jumlahMenu int

			if pilihanMenuKedua == 1 {
				fmt.Print("Masukkan Jumlah Menu: ")
				fmt.Scan(&jumlahMenu)
				for i := 1; i <= jumlahMenu; i++ {
					fmt.Print("Masukkan Nama Makanan: ")
					fmt.Scan(&namaMakanan)
					fmt.Print("Masukkan Nama Minuman: ")
					fmt.Scan(&namaMinuman)
					fmt.Print("Masukkan ID Customer: ")
					fmt.Scan(&idMenu)

					menu := &Menu{
						Id:      idMenu,
						Makanan: namaMakanan,
						Minuman: namaMinuman,
					}
					listMenu.InsertMenu(menu)
				}
				goto menuSecondary3
			} else if pilihanMenuKedua == 2 {
				fmt.Print("Masukkan ID yang akan diupdate: ")
				fmt.Scan(&idMenu)
				fmt.Print("Masukkan Nama Makanan yang baru: ")
				fmt.Scan(&namaBaruMakanan)
				fmt.Print("Masukkan Nama Minuman yang baru: ")
				fmt.Scan(&namaBaruMinuman)
				listMenu.UpdatenamaMENU(idMenu, namaBaruMakanan, namaBaruMinuman)
				goto menuSecondary3
			} else if pilihanMenuKedua == 3 {
				fmt.Print("Masukkan ID yang akan Dihapus: ")
				fmt.Scan(&idMenu)
				listMenu.DeleteMenu(idMenu)
				goto menuSecondary3
			} else if pilihanMenuKedua == 4 {
				listMenu.ViewAllMenu()
				goto menuSecondary3
			} else if pilihanMenuKedua == 5 {
				fmt.Print("Masukkan ID yang Dicari: ")
				fmt.Scan(&idMenu)
				listMenu.viewByIDMENU(idMenu)
				goto menuSecondary3
			} else if pilihanMenuKedua == 6 {
				goto menuUtama
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
