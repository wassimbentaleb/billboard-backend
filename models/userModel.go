package models

type User struct {
	ID        uint `gorm:"primarykey"`
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
}

type Post struct {
	ID        uint `gorm:"primarykey"`
	Email     string
	FirstName string
	LastName  string
	State     string
}
