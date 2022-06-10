package main

import (
	_config "be9_group7_project1/config"
	_entities "be9_group7_project1/entities"
	"database/sql"
	"fmt"

)


var DBConn *sql.DB //variable ini dibuat untuk menyimpan hasil fungsi yang berada di func init

func init() { //fungsi init adalah fungsi yang harus pertama kali dijalankan(kebetulan didalam fungsi ini terdapat koneksi DB) 
	DBConn = _config.ConnectionDB() 
}

func main(){

}