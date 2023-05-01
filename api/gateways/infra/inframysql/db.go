package inframysql

import (
	"context"
	"database/sql"
	"fmt"
	"os"

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
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname) // 修正!!

	Client, err = sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
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
