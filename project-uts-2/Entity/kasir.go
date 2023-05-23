package Entity

type Kasir struct {
	Id   int
	Nama string
	Next *Kasir
}

type LinkedListKasir struct {
	head *Kasir
}