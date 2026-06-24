package user

import "time"

type ID string
type Name string
type Email string
type Password string
type CreatedAt time.Time
type UpdatedAt time.Time

type User struct {
	Id        ID        `json:"id"`
	Name      Name      `json:"name"`
	Email     Email     `json:"email"`
	Password  Password  `json:"-"`
	CreatedAt CreatedAt `json:"created_at"`
	UpdatedAt UpdatedAt `json:"updated_at"`
}
