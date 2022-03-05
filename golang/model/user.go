package model

import(
	"time"
)

type User struct {
	ID       int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
