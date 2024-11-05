package databases

import (
	"database/sql"
)

type MySqlDatabaseRepository struct {
	Connection *sql.DB
}
