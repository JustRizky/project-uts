package View

func InsertCustomer() {
	var Id int
	var Nama string
	fmt.Print("masukkan ID Baru : ")
	fmt.Scan(&Id)
	fmt.Print("masukkan Nama Baru : ")
	fmt.Scan(&Nama)
	controller.ControllerInsertCustomer(Id, Nama)
}

func ViewCustomer() {
	allCustomer := controller.ControllerFindAllCustomer()
	for _, Customer := range allCustomer {
		fmt.Println("id Customer :", Customer.Id)
		fmt.Println("Nama Customer :", Customer.Nama)
	}
}