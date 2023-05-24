package entity

type Customer struct {
	Id        int
	Nama      string
	NomorMeja int
}

type LinkedListCustomer struct {
	Data Customer
	Next *LinkedListCustomer
}
