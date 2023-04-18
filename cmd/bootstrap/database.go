package bootstrap

import (
	"database/sql"
	"fmt"
	"jagch/boletia/freecurrency/config"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func newDatabase() *sql.DB {
	db, err := sql.Open("pgx", fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		"postgres",
		config.Setting.Database.User,
		config.Setting.Database.Password,
		config.Setting.Database.Host,
		config.Setting.Database.Port,
		config.Setting.Database.Name))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Established a successful connection with the DB!")

	return db
}
