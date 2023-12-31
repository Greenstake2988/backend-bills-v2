// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"time"
)

type Bill struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Concept string `json:"concept"`
	// price must be positive
	Price int64     `json:"price"`
	Fecha time.Time `json:"fecha"`
}

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
