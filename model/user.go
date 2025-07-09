package model

type User struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Firstname string `gorm:"not null" json:"firstname"`
	Lastname  string `gorm:"not null" json:"lastname"`
	Username  string `gorm:"unique;not null" json:"username"`
	Age       int    `gorm:"not null" json:"age"`
	Phone     int    `gorm:"unique;not null" json:"phone"`
	Password  string `gorm:"not null" json:"password"`
}
