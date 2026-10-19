package migrator

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
)

func (m Migrator) Down() {
	err := m.migrate.Down()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("migration down successfully ✅")
			return
		}
		log.Fatalf("migration down faild ❌, err: %v", err)
	}
	log.Println("migration down successfully ✅")
}
