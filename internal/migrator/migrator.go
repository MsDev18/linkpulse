package migrator

import (
	"database/sql"
	"fmt"
	mysqlrepo "linkpulse/internal/repository/mysql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct {
	migrate *migrate.Migrate
}

func New(config mysqlrepo.Config) Migrator {
	db, err := sql.Open(
		mysqlrepo.MySqlDriver,
		fmt.Sprintf("%s&%s",config.GetDSN(), "multiStatements=true"),
	)
	if err != nil {
		log.Fatalf("can't connect to database, err: %v", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{
		MigrationsTable: "migrations",
	})
	if err != nil {
		log.Fatalf("error in migration, err: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		mysqlrepo.MySqlDriver,
		driver,
	)

	if err != nil {
		log.Fatalf("error in create migrator instance")
	}

	return Migrator{
		migrate: m,
	}
}
