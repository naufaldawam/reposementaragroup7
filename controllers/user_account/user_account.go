package userAccount

import (
	_entities "be9_group7_project1/entities"
	"database/sql"
	"fmt"
)

// history atau get all data dari userAccount
func GetAllUserAccount(db *sql.DB) ([]_entities.UserAccount, error) {
	results, err := db.Query("SELECT ID, UserName, UserEmail, UserPassword, UserPhone FROM user_account")
	if err != nil {
		// fmt.Println("error", err.Error())
		return []_entities.UserAccount{}, err
	}
	defer db.Close()

	var dataAll []_entities.UserAccount
	for results.Next() {
		var userAccount _entities.UserAccount
		err := results.Scan(&userAccount.ID, &userAccount.UserName, &userAccount.UserEmail, &userAccount.UserPassword, &userAccount.UserPhone)
		if err != nil {
			fmt.Println("error scan transfer", err.Error())
		}
		dataAll = append(dataAll, userAccount)
	}

	return dataAll, nil
}

//create user account
func CreateUserAccount(db *sql.DB, newUserAccount _entities.UserAccount) (int, error) {
	// var query = fmt.Sprintf("INSERT INTO mahasiswa (id, nama, jenis_kelamin, alamat, jurusan_id, telp, status) values (%d, '%s', '%s', '%s', %d, '%s','%s')", newMahasiswa.ID, newMahasiswa.Nama, newMahasiswa.JenisKelamin, newMahasiswa.Alamat, newMahasiswa.JurusanId, newMahasiswa.Telp, newMahasiswa.Status.String)
	var query = (`INSERT INTO user_account (ID, UserName, UserEmail, UserPassword, UserPhone) VALUES (?, ?, ?, ?, ?)`)
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	result, err := statement.Exec(newUserAccount.ID, newUserAccount.UserName, newUserAccount.UserEmail, newUserAccount.UserPassword, newUserAccount.UserPhone)

	defer db.Close()

	if err != nil {
		// fmt.Println("error ketika menambahkan account baru", err.Error())
		return 0, err
	} else {
		// fmt.Println("Akun Baru Berhasil Dibuat")
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

