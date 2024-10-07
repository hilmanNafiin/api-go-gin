package entity

import "time"

type Users struct {
	Id int
	Name string
	Email string
	Gender string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}
// type User struct {
// 	Id int `json:"id"`
// 	Name string `json:"name"`
// 	Email string `json:"email"`
// 	Gender string `json:"gender"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }