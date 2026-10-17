package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MySqlDriver = "mysql"
)

type Config struct {
	Username              string        `koanf:"username"`
	Password              string        `koanf:"password"`
	Host                  string        `koanf:"host"`
	Port                  uint          `koanf:"port"`
	Database              string        `koanf:"database"`
	ConnectionMaxLifeTime time.Duration `koanf:"connection_max_life_time"`
	MaxOpenConnections    int           `koanf:"max_open_connections"`
	MaxIdleConnections    int           `koanf:"max_idle_connections"`
}

func (c Config) GetDSN() string {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", c.Username, c.Password, c.Host, c.Port, c.Database)
	return dsn
}

type MySQLDB struct {
	*sql.DB
}

func New(config Config) MySQLDB {
	db, err := sql.Open(MySqlDriver, config.GetDSN())
	if err != nil {
		log.Fatalf("error in connection to database, err -> %v", err)
	}

	// db settings
	db.SetConnMaxLifetime(config.ConnectionMaxLifeTime)
	db.SetMaxOpenConns(config.MaxOpenConnections)
	db.SetMaxIdleConns(config.MaxIdleConnections)

	return MySQLDB{
		DB: db,
	}
}
