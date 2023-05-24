package entity

type Kasir struct {
	Id   int
	Nama string
	Umur int
}

type LinkedListKasir struct {
	Data Kasir
	Next *LinkedListKasir
}
