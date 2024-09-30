package models

// UserList связь n:n для пользователя и его списков
type UserList struct {
	Id     int
	UserId int
	ListId int
}
