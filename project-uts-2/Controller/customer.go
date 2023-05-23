package Controller

func ControllerFindAllCustomer() []entity.Customer {
	return model.ModelViewAllCustomer()
}

func ControllerInsertCustomer(Id int, Nama string) {
	data := entity.Customer{
		Id:    id,
		Nama:  nama,
	}
	fmt.Println("testing passing data controller ", data)
	model.ModelInsertCustomer(data)
}