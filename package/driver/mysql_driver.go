package driver

import "database/sql"

type MySqlDatabase interface{}

type MySqlDatabaseImplementation struct{}

func (d *MySqlDatabaseImplementation) Open(driverName, dataSourceName string) (*sql.DB, error) {
	return sql.Open(driverName, dataSourceName)
}
