package team

import "database/sql"

type teamRepository struct {
	databaseConnection *sql.DB
}
