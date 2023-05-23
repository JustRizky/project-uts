package Controller

func ControllerFindAllKasir() []entity.Kasir {
	return model.ModelViewAllKasir()
}

func ControllerInsertKasir(Id int, Nama string) {
	data := entity.Kasir{
		Id:    Id,
		Nama:  Nama,
	}
	fmt.Println("testing passing data controller ", data)
	model.ModelInsertKasir(data)
}