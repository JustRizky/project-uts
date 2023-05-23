package entity

type Customer struct {
	Id   int
	Nama string
	Next *Customer
}

type LinkedListCustomer struct {
	head *Customer
}
