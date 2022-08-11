package models

import "time"

type User struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	LastLoginAt time.Time `json:"last_login"`
	Status      int8      `json:"status"`
}

type UserLogin struct {
	Email       string    `json:"email"`
	LastLoginAt time.Time `json:"last_login"`
}
