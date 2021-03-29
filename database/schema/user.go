package schema

type User struct {
	Base
	Name  string
	Email string `gorm:"unique"`
}

func (User) TableName() string {
	return "users"
}
