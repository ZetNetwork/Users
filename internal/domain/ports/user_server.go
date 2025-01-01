package ports

type IUserServer interface {
	CreateUser(email string, password string) error
	DeleteUser(email string) error
	UpdateUser(email string, password string) error
	GetUser(email string)
}
