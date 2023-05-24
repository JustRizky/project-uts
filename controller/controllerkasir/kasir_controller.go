package controller

import (
	"errors"
	"project-uts-mvc/entity"
	model "project-uts-mvc/model/modelkasir"
)

func ControllerInsertKasir(id int, nama string) error {
	find := model.Search(id)
	if find == nil {
		data := entity.Kasir{
			Id:   id,
			Nama: nama,
		}
		model.ModelInsertKasir(data)
		return nil
	}

	return errors.New("<FAILED> Id Sudah Ada!")
}

func ControllerUpdateKasir(id int, nama string) error {
	container := entity.Kasir{
		Id:   id,
		Nama: nama,
	}

	if model.ModelUpdateKasir(container) {
		return nil
	}

	return errors.New("<FAILED> Id Tidak Ada!")
}

func ControllerDeleteKasir(id int) error {
	err := model.ModelDeleteKasir(id)
	if err {
		return nil
	}

	return errors.New("<FAILED> Id Tidak Ada!")
}

func ControllerViewAllKasir() []entity.Kasir {
	return model.ModelViewAllKasir()
}

func ControllerViewByIdKasir(id int) *entity.Kasir {
	current := model.ModelViewByIdKasir(id)
	if current == nil {
		return nil
	}
	return current
}
