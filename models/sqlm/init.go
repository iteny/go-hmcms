package sqlm

import (
	"database/sql"
	"go-hmcms/models/common"

	"github.com/iteny/hmgo/sqlm"
)

// var Db *sqlm.DB
var schema = `
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

var Tx *sqlm.Tx

func init() {
	db, err := sqlm.NewEngine("sqlite3", "./sql/hmcms.db")
	if err != nil {
		common.Log.Error(err)
	}
	// db.MustExec(schema)

	Tx = db.MustBegin()
	Tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	Tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	Tx.Commit()

}
