package entities

import "database/sql"

type UserBalance struct {
	UserAccountId int
	TotalBalance sql.NullInt64
}