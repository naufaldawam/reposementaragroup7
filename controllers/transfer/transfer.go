package transfer

import (
	_entities "be9_group7_project1/entities"
	"database/sql"
	"fmt"
)

// history atau get all data dari transfer
func GetAllTransfer(db *sql.DB) ([]_entities.Transfer, error) {
	results, err := db.Query("SELECT ID, UserIdSent, UserIdReceiver, TotalBalanceTransferred, TrasnferDate FROM transfer")
	if err != nil {
		// fmt.Println("error", err.Error())
		return []_entities.Transfer{}, err
	}
	defer db.Close()

	var dataAll []_entities.Transfer
	for results.Next() {
		var transfer _entities.Transfer
		err := results.Scan(&transfer.ID, &transfer.UserIdReceiver, &transfer.UserIdSent, &transfer.TotalBalanceTransferred, &transfer.TransferDate)
		if err != nil {
			fmt.Println("error scan transfer", err.Error())
		}
		dataAll = append(dataAll, transfer)
	}

	return dataAll, nil
}