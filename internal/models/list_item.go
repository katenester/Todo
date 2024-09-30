package models

// ListItem связь n:n для списка и дел
type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
