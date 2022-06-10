package entities

import "database/sql"

type UserAccount struct{
	ID int
	UserEmail string
	UserName sql.NullString //sama ini bener gk ya tulis not nullnya
	UserPassword string
	UserPhone int //ini aku gk tau cara tulis struct buat nyimpen db klo dia unique
}