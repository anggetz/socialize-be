package cores

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func ConnectToDb(ctx context.Context) *pg.DB {

	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "root",
		Database: "smart",
	})
	// defer db.Close()

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Connected to PostgreSQL")

	return db
}
