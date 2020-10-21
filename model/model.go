package model

import "database/sql"

var db *sql.DB

func SetUpDatabase(d *sql.DB)  {
	db = d
}

type User struct {
	Id int
	Email string
	Name string
	ProfileImage string
	Description string
	Password string
}

type Post struct {
	Id int
	Content string
	Image string
	UserId int
}

type Like struct {
	Id int
	User int
	Post int
}

type Comment struct {
	Id int
	Content string
	User int
	Post int
}