/* Package model*/
package model

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkDBG(err error) {
	if err != nil {
		log.Println(err)
	}
}

var db *sql.DB

func init() {
	newDB, err := sql.Open("sqlite3", "./db.sqlite3")
	check(err)

	_, err = newDB.Exec(`create table if not exists phones (
		id integer primary key, 
		name varchar(30),
		phone varchar(30)
		)`)

	check(err)

	db = newDB
}

func Create(name, phone string) int {
	res, err := db.Exec(`insert into phones (name, phone) values (?, ?)`, name, phone)
	check(err)

	res, err = db.Exec(`insert into phones (name, phone) values (?, ?)`, name, phone)
	check(err)
	id, err := res.LastInsertId()
	check(err)

	return int(id)
}

func Read(name string) (phone string) {
	rows, err := db.Query(`select phone from phones where name = ?`, name)
	check(err)

	if rows.Next() {
		var phone string
		res := rows.Scan(&phone)
		checkDBG(res)
	}

	return
}

func ReadAll() [][]string {
	rows, err := db.Query(`select name, phone from phones`)
	check(err)

	retSlice := make([][]string, 0)
	for rows.Next() {
		c := make([]string, 2)
		res := rows.Scan(&c[0], &c[1])
		checkDBG(res)
		retSlice = append(retSlice, c)
	}

	return retSlice
}

func Delete(name string) {
	_, err := db.Exec(`delete from phones where name =?`, name)
	check(err)
}

func Update(name string, new_name string, new_phone string) {
	_, err := db.Exec(`update phones set name =?, phone=? where name=?`, new_name, new_phone, name)
	check(err)
}

/* CloseDB() close database connection*/
func CloseDB() {
	db.Close()
}
