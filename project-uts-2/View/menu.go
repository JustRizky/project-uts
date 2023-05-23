package View

func InsertMenu() {
	var Id int
	var Makanan string
	var Minuman string
	fmt.Print("masukkan ID Baru : ")
	fmt.Scan(&Id)
	fmt.Print("masukkan Makanan Baru : ")
	fmt.Scan(&Makanan)
	fmt.Print("masukkan Minuman Baru : ")
	fmt.Scan(&Minuman)
	controller.ControllerInsertMenu(Id, Makanan, Minuman)
}

func ViewMenu() {
	allMenu := controller.ControllerFindAllMenu()
	for _, Menu := range allMenu {
		fmt.Println("id Menu :", Menu.Id)
		fmt.Println("Makanan Menu :", Menu.Makanan)
		fmt.Println("Minuman Menu :", Menu.Minuman)
	}
}