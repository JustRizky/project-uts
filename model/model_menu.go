package model

import (
	"fmt"
	"project-uts-mvc/entity"
)

type MenuModel struct {
	DaftarMenu *entity.Menu
}

func (list *MenuModel) InsertMenu(menu *entity.Menu) {
	if list.DaftarMenu == nil {
		list.DaftarMenu = menu
	} else {
		tempMenu := list.DaftarMenu
		for tempMenu.Next != nil {
			tempMenu = tempMenu.Next
		}
		tempMenu.Next = menu
	}

}

func (list *MenuModel) UpdatenamaMENU(Id int, newMAKAN string, newMINUM string) {
	if list.DaftarMenu == nil {
		fmt.Println("Tidak ada menu!")
	}

	curr := list.DaftarMenu
	for curr != nil {
		if curr.Id == Id {
			curr.Makanan = newMAKAN
			curr.Minuman = newMINUM
			return
		}
		curr = curr.Next
	}
}

func (list *MenuModel) DeleteMenu(Id int) {
	if list.DaftarMenu == nil {
		fmt.Println("Menu tidak ditemukan!")
		return
	}

	if list.DaftarMenu.Id == Id {
		list.DaftarMenu = list.DaftarMenu.Next
		return
	}

	curr := list.DaftarMenu
	for curr.Next != nil {
		if curr.Next.Id == Id {
			curr.Next = curr.Next.Next
			return
		}
		curr = curr.Next
	}

	fmt.Println("Menu tidak ditemukan!")
}

func (list *MenuModel) ViewAllMenu() {
	if list.DaftarMenu == nil {
		fmt.Println("")
	} else {
		tempRestoran := list.DaftarMenu
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

func (list *MenuModel) viewByIDMENU(Id int) {
	curr := list.DaftarMenu
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
