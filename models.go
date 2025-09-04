package main

import "time"

type UserProfile struct {
	ID           uint      `gorm:"primaryKey"     json:"id"`
	Username     string    `gorm:"unique"         json:"username"`
	Password     string    `                      json:"password"`
	RegisteredAt time.Time `gorm:"autoCreateTime" json:"registed_at"`
}

// A user can make many posts
// A post is made by one user
type Post struct {
	ID            uint      `gorm:"primaryKey"     json:"id"`
	Content       string    `gorm:"not null"       json:"content"`
	Likes         uint      `                      json:"likes"`
	Dislikes      uint      `                      json:"dislikes"`
	PostedAt      time.Time `gorm:"autoCreateTime" json:"posted_at"`
	UserProfileID uint      `gorm:"not null;index" json:"user_id"`
}

// A post has many comments
// A comment is linked to one post
type Comment struct {
	ID            uint      `gorm:"primaryKey"     json:"id"`
	Content       string    `gorm:"not null"       json:"content"`
	CommentedAt   time.Time `gorm:"autoCreateTime" json:"commented_at"`
	UserProfileID uint      `gorm:"not null;index" json:"user_id"`
	PostID        uint      `gorm:"not null;index" json:"post_id"`
}
