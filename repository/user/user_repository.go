package user

import "database/sql"

type userRespository struct {
	databaseConnection *sql.DB
}
