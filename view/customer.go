package view

import (
	"fmt"
	"project-uts-mvc/entity"
)

type CustomerView struct{}

func (view *CustomerView) ViewCustomer(customer *entity.Customer) {
	if customer == nil {
		fmt.Println("Customer Tidak Ada!")
	} else {
		fmt.Printf("Id: %d\n", customer.Id)
		fmt.Printf("Nama: %s\n", customer.Nama)
	}
}
func (view *CustomerView) ViewUpdateCustomer(status bool) {
	if status {
		fmt.Println("Data Customer Diupdate!")
	} else {
		fmt.Println("Data Customer Gagal Diupdate!")
	}
}

func (view *CustomerView) ViewDeleteCustomer(status bool) {
	if status {
		fmt.Println("Data Customer Didelete!")
	} else {
		fmt.Println("Data Customer Gagal Didelete!")
	}
}
