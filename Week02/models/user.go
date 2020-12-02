package models

type User struct {
	Id   int `db:"id"`
	Name string `db:"name"`
}