package entity

type Menu struct {
	Id      int
	Makanan string
	Minuman string
	Next    *Menu
}

type LinkedListMenu struct {
	head *Menu
}
