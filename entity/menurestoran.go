package entity

type MenuRestoran struct {
	Id      int
	Makanan string
	Minuman string
}

type LinkedListMenuRestoran struct {
	Data MenuRestoran
	Next *LinkedListMenuRestoran
}
