package model

import (
	"project-uts-mvc/db"
	"project-uts-mvc/entity"
)

func ModelInsertMenuRestoran(container entity.MenuRestoran) {
	newGerbong := entity.LinkedListMenuRestoran{}
	newGerbong.Data = container
	temp := &db.DBMenuRestoran
	if temp.Next == nil {
		temp.Next = &newGerbong
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newGerbong
	}
}

func ModelUpdateMenuRestoran(newData entity.MenuRestoran) bool {
	current := &db.DBMenuRestoran
	for current != nil {
		if current.Data.Id == newData.Id {
			current.Data = newData
			return true
		}
		current = current.Next
	}
	return false
}

func ModelDeleteMenuRestoran(id int) bool {
	prev := &db.DBMenuRestoran
	current := db.DBMenuRestoran.Next

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

func ModelViewByIdMenuRestoran(id int) *entity.MenuRestoran {
	temp := &db.DBMenuRestoran
	for temp != nil {
		if temp.Data.Id == id {
			return &temp.Data
		}
		temp = temp.Next
	}
	return nil
}

func ModelViewAllMenuRestoran() []entity.MenuRestoran {
	temp := db.DBMenuRestoran.Next
	members := []entity.MenuRestoran{}
	for temp != nil {
		members = append(members, temp.Data)
		temp = temp.Next
	}
	return members
}

func Search(id int) *entity.LinkedListMenuRestoran {
	current := db.DBMenuRestoran.Next
	for current != nil {
		if current.Data.Id == id {
			return current
		}
		current = current.Next
	}
	return nil
}
