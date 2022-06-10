package entities

import "database/sql"

type Transfer struct {
	ID                      int
	UserIdSent              int
	UserIdReceiver          int
	TotalBalanceTransferred int
	TransferDate            sql.NullTime
}