package models

type Employee struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	Name string `grom:"varchar(255)" json:"name"`
	Position string `gorm:"varchar(255)" json:"position"`
}