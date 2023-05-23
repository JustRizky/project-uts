package view

import (
	"fmt"
	"project-uts-mvc/entity"
)

type MenuRestoanView struct{}

func (view *MenuRestoanView) ViewMenuRestoran(menurestoran *entity.Menu) {
	if menurestoran == nil {
		fmt.Println("Kasir Tidak Ada!")
	} else {
		fmt.Printf("Id: %d\n", menurestoran.Id)
		fmt.Printf("Nama: %s\n", menurestoran.Makanan)
		fmt.Printf("Nama: %s\n", menurestoran.Minuman)
	}
}
func (view *MenuRestoanView) ViewUpdateKasir(status bool) {
	if status {
		fmt.Println("Data Kasir Diupdate!")
	} else {
		fmt.Println("Data Kasir Gagal Diupdate!")
	}
}

func (view *MenuRestoanView) ViewDeleteKasir(status bool) {
	if status {
		fmt.Println("Data Kasir Didelete!")
	} else {
		fmt.Println("Data Kasir Gagal Didelete!")
	}
}
