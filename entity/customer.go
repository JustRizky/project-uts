package entity

type Customer struct {
	Id   int
	Nama string
}

type LinkedListCustomer struct {
	Data Customer
	Next *LinkedListCustomer
}
