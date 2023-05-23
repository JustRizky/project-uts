package View

func InsertKasir() {
	var Id int
	var Nama string
	fmt.Print("masukkan ID Baru : ")
	fmt.Scan(&Id)
	fmt.Print("masukkan Nama Baru : ")
	fmt.Scan(&Nama)
	controller.ControllerInsertKasir(Id, Nama)
}

func ViewKasir() {
	allKasir := controller.ControllerFindAllKasir()
	for _, Kasir := range allKasir {
		fmt.Println("id Kasir :", Kasir.Id)
		fmt.Println("Nama Kasir :", Kasir.Nama)
	}
}