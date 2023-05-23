package model

import (
	"fmt"
	"project-uts-mvc/entity"
)

type CustomerModel struct {
	DaftarCustomer *entity.Customer
}

func (list *CustomerModel) InsertCS(cs *entity.Customer) {
	if list.DaftarCustomer == nil {
		list.DaftarCustomer = cs
	} else {
		tempCs := list.DaftarCustomer
		for tempCs.Next != nil {
			tempCs = tempCs.Next
		}
		tempCs.Next = cs
	}

}

func (list *CustomerModel) UpdatenamaCS(Id int, newCS string) {
	if list.DaftarCustomer == nil {
		fmt.Println("Tidak ada customer!")
	}

	curr := list.DaftarCustomer
	for curr != nil {
		if curr.Id == Id {
			curr.Nama = newCS
			return
		}
		curr = curr.Next
	}
}

func (list *CustomerModel) DeleteCS(ID int) {
	if list.DaftarCustomer == nil {
		fmt.Println("Customer tidak ditemukan!")
		return
	}

	if list.DaftarCustomer.Id == ID {
		list.DaftarCustomer = list.DaftarCustomer.Next
		return
	}

	currCustomer := list.DaftarCustomer
	for currCustomer.Next != nil {
		if currCustomer.Next.Id == ID {
			currCustomer.Next = currCustomer.Next.Next
			return
		}
		currCustomer = currCustomer.Next
	}

	fmt.Println("Customer tidak ditemukan!")
}

func (list *CustomerModel) ViewAllCustomer() {
	if list.DaftarCustomer == nil {
		fmt.Println("")
	} else {
		tempRestoran := list.DaftarCustomer
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

func (list *CustomerModel) viewByIDCS(id int) {
	curr := list.DaftarCustomer
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
