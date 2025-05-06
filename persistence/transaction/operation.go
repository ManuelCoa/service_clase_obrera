package transaction

import (
	"database/sql"

	"claseobrera/persistence/connection"
)

// Función de orden superior que retorna las funciones `read` y `write`
func DatabaseOperations() (func(string, ...interface{}) (*sql.Rows, error), func(string, ...interface{}) (sql.Result, error)) {

	openConn, _ := connection.ConnectionFactory()

	read := func(query string, args ...interface{}) (*sql.Rows, error) {
		db, err := openConn()
		if err != nil {
			return nil, err
		}

		stmt, err := db.Prepare(query)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		rows, err := stmt.Query(args...)
		if err != nil {
			return nil, err
		}

		return rows, nil
	}

	write := func(query string, args ...interface{}) (sql.Result, error) {
		db, err := openConn()
		if err != nil {
			return nil, err
		}

		stmt, err := db.Prepare(query)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		result, err := stmt.Exec(args...)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	return read, write
}
