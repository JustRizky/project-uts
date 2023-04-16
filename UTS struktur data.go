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
			curr=curr.Next
		}
}

func(list *LinkedListKSR) UpdatenamaKSR(oldKSR *Kasir, newKSR string) {
	if list.head == nil {
		fmt.Println("TIDAK ADA BARANG")
	}

	curr := list.head
		for curr != nil {
			if curr == oldKSR {
				curr.Nama = newKSR
				return
			}
			curr=curr.Next
		}
}

func(list *LinkedListMenu) UpdatenamaMENU(oldMENU *Menu, newMAKAN string, newMINUM string) {
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
			curr=curr.Next
		}
}


//DELETE

func (list *LinkedListKSR) DeleteKSR(ID int) {
	// jika node kosong
	if list.head == nil {
		fmt.Println("Kasir tidak ditemukan")
		return
	}

	// jika node 
	if list.head.Nik == ID {
		list.head = list.head.Next
		return
	}

	currKasir := list.head
	for currKasir.Next != nil {
		if currKasir.Next.Nik == ID {
			currKasir.Next = currKasir.Next.Next
			return
		}
		currKasir = currKasir.Next
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

	currCustomer := list.head
	for currCustomer.Next != nil {
		if currCustomer.Next.Np == ID {
			currCustomer.Next = currCustomer.Next.Next
			return
		}
		currCustomer = currCustomer.Next
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

	currMenu := list.head
	for currMenu.Next != nil {
		if currMenu.Next.ID == ID {
			currMenu.Next = currMenu.Next.Next
			return
		}
		currMenu = currMenu.Next
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

			fmt.Println("_________")
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

			fmt.Println("_________")
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

	list3.ViewAllKasir()
	list2.ViewAllCustomer()
	list4.ViewAllMenu()

	fmt.Println("\n========= SESUDAH DELETE ============")

	list3.DeleteKSR(01)
	list4.DeleteMenu(01)
	list2.DeleteCS(99)

	list3.ViewAllKasir()
	list2.ViewAllCustomer()
	list4.ViewAllMenu()

	fmt.Println("\n========= SESUDAH UPDATE ============")
	list3.UpdatenamaKSR(kasir2 , "RIZKI IMOET")
	list4.UpdatenamaMENU(menu2, "SEGO JAGUNG" , "ES DEGAN")
	list2.UpdatenamaCS(customer2 , "RANGGA GANTENG")
	list3.ViewAllKasir()
	list2.ViewAllCustomer()
	list4.ViewAllMenu()
}