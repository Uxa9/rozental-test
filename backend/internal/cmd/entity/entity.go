package entity

var User = struct {
	UserTable string
	name      string
}{}

func Init() {
	User.UserTable = "user"
	User.name = "name"
}
