package topup

import (
	_entities "be9_group7_project1/entities"
	"database/sql"
	"fmt"
)

// history atau get all data dari topup
func GetAllTopup(db *sql.DB) ([]_entities.Topup, error) {
	results, err := db.Query("SELECT id, user_account_id, where_top_go, total_topup, topup_date FROM topup")
	if err != nil {
		// fmt.Println("error", err.Error())
		return []_entities.Topup{}, err
	}
	defer db.Close()

	var dataAll []_entities.Topup
	for results.Next() {
		var topup _entities.Topup
		err := results.Scan(&topup.ID, &topup.UserAccountId, &topup.WhereTopupGo, &topup.TotalTopup, &topup.TopupDate)
		if err != nil {
			fmt.Println("error scan topup", err.Error())
		}
		dataAll = append(dataAll, topup)
	}

	return dataAll, nil
}

