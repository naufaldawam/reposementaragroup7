package userBalance

import(
	_entities "be9_group7_project1/entities"
	"database/sql"
	"fmt"
)

// history atau get all data dari userBalance
func GetAllUserBalance(db *sql.DB) ([]_entities.UserBalance, error) {
	results, err := db.Query("SELECT UserAccountID, TotalBalance User FROM user_account")
	if err != nil {
		// fmt.Println("error", err.Error())
		return []_entities.UserBalance{}, err
	}
	defer db.Close()

	var dataAll []_entities.UserBalance
	for results.Next() {
		var userAccount _entities.UserBalance
		err := results.Scan(&userAccount.UserAccountId, &userAccount.TotalBalance)
		if err != nil {
			fmt.Println("error scan transfer", err.Error())
		}
		dataAll = append(dataAll, userAccount)
	}

	return dataAll, nil
}
