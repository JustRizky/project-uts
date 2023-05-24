package model

import (
	"project-uts-mvc/db"
	"project-uts-mvc/entity"
)

func ModelInsertKasir(container entity.Kasir) {
	newGerbong := entity.LinkedListKasir{}
	newGerbong.Data = container
	temp := &db.DBKasir
	if temp.Next == nil {
		temp.Next = &newGerbong
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newGerbong
	}
}

func ModelUpdateKasir(newData entity.Kasir) bool {
	current := &db.DBKasir
	for current != nil {
		if current.Data.Id == newData.Id {
			current.Data = newData
			return true
		}
		current = current.Next
	}
	return false
}

func ModelDeleteKasir(id int) bool {
	prev := &db.DBKasir
	current := db.DBKasir.Next

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

func ModelViewByIdKasir(id int) *entity.Kasir {
	temp := &db.DBKasir
	for temp != nil {
		if temp.Data.Id == id {
			return &temp.Data
		}
		temp = temp.Next
	}
	return nil
}

func ModelViewAllKasir() []entity.Kasir {
	temp := db.DBKasir.Next
	members := []entity.Kasir{}
	for temp != nil {
		members = append(members, temp.Data)
		temp = temp.Next
	}
	return members
}

func Search(id int) *entity.LinkedListKasir {
	current := db.DBKasir.Next
	for current != nil {
		if current.Data.Id == id {
			return current
		}
		current = current.Next
	}
	return nil
}
