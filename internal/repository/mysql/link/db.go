package link

import "linkpulse/internal/repository/mysql"

type Repository struct {
	conn mysql.MySQLDB
}

func New(conn mysql.MySQLDB) Repository {
	return Repository{
		conn: conn,
	}
}
