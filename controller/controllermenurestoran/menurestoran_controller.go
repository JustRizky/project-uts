package controller

import (
	"errors"
	"project-uts-mvc/entity"
	model "project-uts-mvc/model/modelmenurestoran"
)

func ControllerInsertMenuRestoran(id int, makanan, minuman string) error {
	find := model.Search(id)
	if find == nil {
		data := entity.MenuRestoran{
			Id:      id,
			Makanan: makanan,
			Minuman: minuman,
		}
		model.ModelInsertMenuRestoran(data)
		return nil
	}

	return errors.New("<FAILED> Id Sudah Ada!")
}

func ControllerUpdateMenuRestoran(id int, makanan, minuman string) error {
	container := entity.MenuRestoran{
		Id:      id,
		Makanan: makanan,
		Minuman: minuman,
	}

	if model.ModelUpdateMenuRestoran(container) {
		return nil
	}

	return errors.New("<FAILED> Id Tidak Ada!")
}

func ControllerDeleteMenuRestoran(id int) error {
	err := model.ModelDeleteMenuRestoran(id)
	if err {
		return nil
	}

	return errors.New("<FAILED> Id Tidak Ada!")
}

func ControllerViewAllMenuRestoran() []entity.MenuRestoran {
	return model.ModelViewAllMenuRestoran()
}

func ControllerViewByIdMenuRestoran(id int) *entity.MenuRestoran {
	current := model.ModelViewByIdMenuRestoran(id)
	if current == nil {
		return nil
	}
	return current
}
