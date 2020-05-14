package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
)

type Database struct {
	Db     *sql.DB
	Params *DbParams
}

type DbParams struct {
	user    string
	pass    string
	address string
	dbName  string
}

// type Service interface {
// 	getOne() (string, err)
// 	getAll() (string, err)
// 	update() (string, err)
// 	insert() (string, err)
// }

// New Authstring and DbOpen using Params struct
func GenerateAuthString(params DbParams) string {
	const temp = "{{.root}}:{{.}}"
	return fmt.Sprintf("%s:%s@%s/%s", params.user, params.pass, params.address, params.dbName)
}

func DbOpen(dbtype string, authString string) *sql.DB {
	db, err := sql.Open(dbtype, authString)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func GenerateDb(user string, pass string, address string, dbName string) Database {
	params := DbParams{user, pass, address, dbName}
	authString := GenerateAuthString(params)

	db := DbOpen("mysql", authString)

	return Database{db, &params}
}

func (db Database) Insert(dbName string, data interface{}) string {
	// Gets the fields of Struct obj and uses it to initialize names of database rows

	// Get FieldNames
	datastr := reflect.ValueOf(data).Type()
	dataFields := make([]string, datastr.NumField())
	dataValues := make([]interface{}, datastr.NumField())
	for i := 0; i < datastr.NumField(); i++ {
		field := string(datastr.Field(i).Name)
		dataValues[i] = reflect.ValueOf(data).Field(i).Interface()
		dataFields[i] = field
	}

	// Generate correct number of question marks in Query
	t_arr := make([]string, datastr.NumField())
	for i, _ := range t_arr {
		t_arr[i] = "?"
	}

	// Create Database Insert Query String
	sqlInsert := `INSERT INTO %s (%s) VALUES (%s)`
	structFields := strings.Join(dataFields, ",")
	sqlInsert = fmt.Sprintf(sqlInsert, dbName, structFields, strings.Join(t_arr, ","))

	// Generate Query
	stmt, err := db.Db.Prepare(sqlInsert)
	if err != nil {
		fmt.Println(err)
	}
	_, err = stmt.Exec(dataValues...)
	if err != nil {
		fmt.Println(err)
	}

	return "It did thing"
}
