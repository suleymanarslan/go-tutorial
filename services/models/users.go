package models

import
(
	"time"
)

type User struct {
	Id     string `json:"id" form:"-"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email string `json:"email" form:"email"`
	DateJoined time.Time `json:"datejoined" form:"datejoined"`
	IsActive bool `json:"active" form:"active"`
	Name string `json:"name" form:"name"`
	Surname string `json:"surname" form:"surname"`
}
