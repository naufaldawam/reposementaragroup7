package entities

import "database/sql"

type Topup struct {
	ID            int
	UserAccountId int
	WhereTopupGo  int
	TotalTopup   int
	TopupDate    sql.NullTime //nanti main klo mau get datanya mungkin bisa di parse, masih belum tau dia bisa nyimpen tanggal atau ngak mas ghalib
}
