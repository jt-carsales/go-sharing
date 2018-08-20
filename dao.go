package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//DBconn ...
type DBconn struct {
	user string
	pass string
	host string
	port string
	name string
	conn *sql.DB
}

//ConnectDB ...
func (db *DBconn) ConnectDB() error {

	var dbURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db.user, db.pass, db.host, db.port, db.name)
	var err error
	db.conn, err = sql.Open("mysql", dbURL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//GetNumbers ...
func (db *DBconn) GetNumbers() ([]Number, error) {
	var nums []Number
	var num Number
	query := "SELECT * FROM numbers ORDER BY isocc asc LIMIT 50 "
	res, err := db.conn.Query(query)
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		return nil, err
	}
	for res.Next() {
		err = res.Scan(&num.DID, &num.ISOCC)
		if err != nil {
			fmt.Println(fmt.Sprint(err))
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, err
}

//GetNumber ...
func (db *DBconn) GetNumber(number Number) ([]Number, error) {
	var nums []Number
	var num Number
	query := "SELECT * FROM numbers WHERE did like '" + number.DID + "%'"
	res, err := db.conn.Query(query)
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		return nil, err
	}
	for res.Next() {
		err = res.Scan(&num.DID, &num.ISOCC)
		if err != nil {
			fmt.Println(fmt.Sprint(err))
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, err
}

//AddNumber ...
func (db *DBconn) AddNumber(number Number) error {
	_, err := db.conn.Query("INSERT INTO numbers values ('" + number.DID + "','" + number.ISOCC + "')")
	if err != nil {
		return err
	}
	return nil
}

//DeleteNumber ...
func (db *DBconn) DeleteNumber(number Number) error {
	_, err := db.conn.Query("DELETE FROM numbers where did='" + number.DID + "' and isocc='" + number.ISOCC + "'")
	if err != nil {
		return err
	}
	return nil
}
