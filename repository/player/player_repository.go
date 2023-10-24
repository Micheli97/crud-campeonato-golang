package player

import "database/sql"

type playerRepository struct {
	database *sql.DB
}
