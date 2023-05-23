package Model

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

	curr := list.head
	for curr.Next != nil {
		if curr.Next.Id == Id {
			curr.Next = curr.Next.Next
			return
		}
		curr = curr.Next
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