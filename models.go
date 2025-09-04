package main

type UserProfile struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Username string `gorm:"unique"`
	Password string `json:"-"`
}
