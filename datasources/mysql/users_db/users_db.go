package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_usersdb_username = "mysql_usersdb_username"
	mysql_usersdb_password = "mysql_usersdb_password"
	mysql_usersdb_host     = "mysql_usersdb_host"
	mysql_usersdb_name     = "mysql_usersdb_name"
)

var (
	ClientDB *sql.DB
	username = os.Getenv(mysql_usersdb_username)
	password = os.Getenv(mysql_usersdb_password)
	host     = os.Getenv(mysql_usersdb_host)
	dbname   = os.Getenv(mysql_usersdb_name)
)

type Row struct {
	TableName string
}

//func init () {
func OpenDBConn() {
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		dbname,
	)
	var err error
	ClientDB, err = sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	if ClientDB == nil {
		panic("User DB is null!")
	}
	if err := ClientDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("Successfully connected to MySQL database!")

	/*
		// Following lines of code is to prove DB connectivity, Not needed in real application
		if err := ClientDB.Ping(); err != nil {
			panic(err)
		}
		log.Println("Pinged DB again!")

		resultset, err := ClientDB.Query("show tables;")
		if err != nil {
			panic(err)
		}
		defer resultset.Close()
		if resultset != nil {

			var results []Row
			var result Row

			for resultset.Next() {
				errr := resultset.Scan(&result.TableName)
				if errr != nil {
					log.Println(errr)
				}
				results = append(results, result)
			}
			log.Printf("Got results from DB %v", results)
			log.Println("")
		}
		// Above lines of code is to prove DB connectivity, Not needed in real application
	*/
}
