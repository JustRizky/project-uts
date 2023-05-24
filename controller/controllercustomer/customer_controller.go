package controller

import (
	"errors"
	"project-uts-mvc/entity"
	model "project-uts-mvc/model/modelcustomer"
)

func ControllerInsertCustomer(id, nomormeja int, nama string) error {
	find := model.Search(id)
	if find == nil {
		data := entity.Customer{
			Id:        id,
			Nama:      nama,
			NomorMeja: nomormeja,
		}
		model.ModelInsertCustomer(data)
		return nil
	}

	return errors.New("<FAILED> Id Sudah Ada!")
}

func ControllerUpdateCustomer(id, nomormeja int, nama string) error {
	container := entity.Customer{
		Id:        id,
		Nama:      nama,
		NomorMeja: nomormeja,
	}

	if model.ModelUpdateCustomer(container) {
		return nil
	}

	return errors.New("<FAILED> Id Tidak Ada!")
}

func ControllerDeleteCustomer(id int) error {
	err := model.ModelDeleteCustomer(id)
	if err {
		return nil
	}

	return errors.New("<FAILED> Id Tidak Ada!")
}

func ControllerViewAllCustomer() []entity.Customer {
	return model.ModelViewAllCustomer()
}

func ControllerViewByIdCustomer(id int) *entity.Customer {
	current := model.ModelViewByIdCustomer(id)
	if current == nil {
		return nil
	}
	return current
}
