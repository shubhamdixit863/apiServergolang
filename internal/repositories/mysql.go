package repositories

type mysql struct {
	// It will have the Db connection object
}

func (m mysql) CreateUser() string {
	return "Hello, World!"
}

func NewMysql() Repository {
	return &mysql{}
}
