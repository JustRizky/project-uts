package Model

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

	currKasir := list.head
	for currKasir.Next != nil {
		if currKasir.Next.Id == ID {
			currKasir.Next = currKasir.Next.Next
			return
		}
		currKasir = currKasir.Next
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