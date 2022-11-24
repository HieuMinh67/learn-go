package models

import "time"

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	CreateAt time.Time
}
