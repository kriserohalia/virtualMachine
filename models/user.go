package models

type User struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	FullName string `grom:"varchar(255)" json:"full_name"`
	Username string `gorm:"varchar(255)" json:"username"`
	Password string `gorm:"varchar(255)" json:"passowrd"`
}