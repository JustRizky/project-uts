package view

import (
	"fmt"
	"project-uts-mvc/entity"
)

type KasirView struct{}

func (view *KasirView) ViewKasir(kasir *entity.Kasir) {
	if kasir == nil {
		fmt.Println("Kasir Tidak Ada!")
	} else {
		fmt.Printf("Id: %d\n", kasir.Id)
		fmt.Printf("Nama: %s\n", kasir.Nama)
	}
}
func (view *KasirView) ViewUpdateKasir(status bool) {
	if status {
		fmt.Println("Data Kasir Diupdate!")
	} else {
		fmt.Println("Data Kasir Gagal Diupdate!")
	}
}

func (view *KasirView) ViewDeleteKasir(status bool) {
	if status {
		fmt.Println("Data Kasir Didelete!")
	} else {
		fmt.Println("Data Kasir Gagal Didelete!")
	}
}
