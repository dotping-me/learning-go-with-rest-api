package main

type UserProfile struct {
	ID       uint `gorm:"primaryKey"`
	Email    string
	Username string
}
