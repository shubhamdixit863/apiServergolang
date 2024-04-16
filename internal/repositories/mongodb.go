package repositories

type mongodb struct {
}

func (m mongodb) CreateUser() string {
	return "Hey mongo db"
}

func NewMongodb() Repository {
	return &mongodb{}
}
