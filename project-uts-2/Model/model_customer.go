package Model

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

	currCustomer := list.head
	for currCustomer.Next != nil {
		if currCustomer.Next.Id == ID {
			currCustomer.Next = currCustomer.Next.Next
			return
		}
		currCustomer = currCustomer.Next
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