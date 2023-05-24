package model

import (
	"project-uts-mvc/db"
	"project-uts-mvc/entity"
)

func ModelInsertCustomer(container entity.Customer) {
	newGerbong := entity.LinkedListCustomer{}
	newGerbong.Data = container
	temp := &db.DBCustomer
	if temp.Next == nil {
		temp.Next = &newGerbong
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newGerbong
	}
}

func ModelUpdateCustomer(newData entity.Customer) bool {
	current := &db.DBCustomer
	for current != nil {
		if current.Data.Id == newData.Id {
			current.Data = newData
			return true
		}
		current = current.Next
	}
	return false
}

func ModelDeleteCustomer(id int) bool {
	prev := &db.DBCustomer
	current := db.DBCustomer.Next

	for current != nil {
		if current.Data.Id == id {
			prev.Next = current.Next
			return true
		}
		prev = current
		current = current.Next
	}

	return false
}

func ModelViewByIdCustomer(id int) *entity.Customer {
	temp := &db.DBCustomer
	for temp != nil {
		if temp.Data.Id == id {
			return &temp.Data
		}
		temp = temp.Next
	}
	return nil
}

func ModelViewAllCustomer() []entity.Customer {
	temp := db.DBCustomer.Next
	members := []entity.Customer{}
	for temp != nil {
		members = append(members, temp.Data)
		temp = temp.Next
	}
	return members
}

func Search(id int) *entity.LinkedListCustomer {
	current := db.DBCustomer.Next
	for current != nil {
		if current.Data.Id == id {
			return current
		}
		current = current.Next
	}
	return nil
}
