package migrator

import "log"

func (m Migrator) Close() {
	sourceErr, databaseErr := m.migrate.Close()
	if sourceErr != nil {
		log.Fatalf("close migrator err: %v", sourceErr)
	}
	if databaseErr != nil {
		log.Fatalf("close migrator err: %v", databaseErr)
	}
}
