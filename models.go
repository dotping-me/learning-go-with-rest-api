package main

type UserProfile struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"unique"     json:"email"`
	Username string `gorm:"unique"     json:"username"`
	Password string `                  json:"password"`
}
