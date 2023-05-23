package model

import (
	"fmt"
	"project-uts-mvc/entity"
)

type KasirModel struct {
	DaftarKasir *entity.Kasir
}

func (list *KasirModel) InsertKSR(ksr *entity.Kasir) {
	if list.DaftarKasir == nil {
		list.DaftarKasir = ksr
	} else {
		tempksr := list.DaftarKasir
		for tempksr.Next != nil {
			tempksr = tempksr.Next
		}
		tempksr.Next = ksr
	}
}

func (list *KasirModel) UpdatenamaKSR(Id int, newKSR string) {
	if list.DaftarKasir == nil {
		fmt.Println("Tidak ada kasir!")
	}

	curr := list.DaftarKasir
	for curr != nil {
		if curr.Id == Id {
			curr.Nama = newKSR
			return
		}
		curr = curr.Next
	}
}

func (list *KasirModel) DeleteKSR(ID int) {
	if list.DaftarKasir == nil {
		fmt.Println("Kasir tidak ditemukan!")
		return
	}

	if list.DaftarKasir.Id == ID {
		list.DaftarKasir = list.DaftarKasir.Next
		return
	}

	currKasir := list.DaftarKasir
	for currKasir.Next != nil {
		if currKasir.Next.Id == ID {
			currKasir.Next = currKasir.Next.Next
			return
		}
		currKasir = currKasir.Next
	}

	fmt.Println("Kasir tidak ditemukan!")
}

func (list *KasirModel) ViewAllKasir() {
	if list.DaftarKasir == nil {
		fmt.Println("")
	} else {
		tempRestoran := list.DaftarKasir
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

func (list *KasirModel) viewByIDKSR(id int) {
	curr := list.DaftarKasir
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
