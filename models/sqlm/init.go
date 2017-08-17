package sqlm

import (
	"database/sql"
	"go-hmcms/models/common"

	"github.com/jmoiron/sqlx"
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

var DB *sqlx.DB

func init() {
	var err error
	DB, err = sqlx.Connect("sqlite3", "./sql/hmcms.db")
	if err != nil {
		common.Log.Error(err)
	}
	// db.MustExec(schema)

	// tx := DB.MustBegin()
	// tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	// tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	// tx.Commit()
	// sx := DB.MustBegin()
	// sx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	// sx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	// sx.Commit()

}
