package login

import "database/sql"

type loginRepository struct {
	databaseConnection *sql.DB
}
