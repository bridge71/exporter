package models

type User struct {
	Email        string `gorm:"column:email;unique;not null" json:"email"`
	PasswordHash string `gorm:"column:passwordHash; not null" json:"password"`
	UserName     string `gorm:"column:userName;not null" json:"userName"`
	UserId       uint   `gorm:"column:userId;unique;primaryKey" `
	Priority     uint   `gorm:"default:1" `
}

type Message struct {
	RetMessage string
	User       User
}
