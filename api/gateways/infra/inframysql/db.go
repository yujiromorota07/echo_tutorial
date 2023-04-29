package inframysql

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

type Dao interface {
	Prepare(query string) (*sql.Stmt, error)
}

var TxKey = struct{}{}

func init() {
	var err error
	Client, err = sql.Open("mysql", "echo:echo@tcp(127.0.0.1:3306)/echodev")
	if err != nil {
		panic(err)
	}

	log.Println("database successfully echodev!")
}

func GetDao(ctx context.Context) Dao {
	tx, ok := GetTx(ctx)
	if !ok {
		return Client
	}

	return tx
}

func GetTx(ctx context.Context) (*sql.Tx, bool) {
	tx, ok := ctx.Value(&TxKey).(*sql.Tx)
	return tx, ok
}
