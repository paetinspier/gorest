package models

type User struct {
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Dob       string `json:"dob" db:"dob"`
	Uid       string `json:"uid" db:"uid"`
	Id        int    `json:"id" db:"id"`
	Active    bool   `json:"active" db:"active"`
	Verified  bool   `json:"verified" db:"verified"`
}
