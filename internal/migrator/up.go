package migrator

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
)

func (m Migrator) Up() {
	err := m.migrate.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("migration up successfully ✅")
			return 
		}
		log.Fatalf("migration up faild ❌ err: %v", err)
	}
	log.Println("migration up successfully ✅")
}
