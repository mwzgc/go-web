package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "root"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "test"
)

var cacheDb *sql.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)

	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	cacheDb = db
}

func getDb() *sql.DB {
	return cacheDb
}

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"` //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
	Age  int    `db:"age"`
}

func QueryMulti() []User {
	DB := getDb()
	var users []User
	rows, err := DB.Query("select * from user limit 20")
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return users
	}

	for rows.Next() {
		user := new(User)
		err = rows.Scan(&user.ID, &user.Name, &user.Age)

		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return users
		}

		users = append(users, *user)
	}
	// fmt.Print(users)
	return users
}
