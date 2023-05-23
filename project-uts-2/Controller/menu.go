package Controller

func ControllerFindAllMenu() []entity.Menu {
	return model.ModelViewAllMenu()
}

func ControllerInsertMenu(Id int, Makanan string, Minuman string) {
	data := entity.Menu{
		Id:    Id,
		Makanan: Makanan,
		Minuman: Minuman,
	}
	fmt.Println("testing passing data controller ", data)
	model.ModelInsertMenu(data)
}