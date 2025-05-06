package driver

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/*
	CreateConnectionFactory es una función de orden superior que toma una cadena DSN (Data Source Name)
	como parámetro y retorna una función que cuando se ejecuta, establece y devuelve una conexión a MySQL.
	Esta función gestiona la apertura de la conexión utilizando sql.Open, y asegura que solo se cree una única
	instancia de la base de datos (singleton), almacenándola en la variable global 'db'.
	También configura un pool de conexiones con límites en el número máximo de conexiones abiertas,
	conexiones en espera y la duración máxima de vida de una conexión.
	Finalmente, verifica que la conexión esté activa mediante db.Ping antes de retornar la instancia de conexión.
*/

func CreateConnectionFactory(dsn string) func() (*sql.DB, error) {
	return func() (*sql.DB, error) {
		if db == nil {
			var err error
			db, err = sql.Open("mysql", dsn)
			if err != nil {
				return nil, err
			}

			db.SetMaxOpenConns(20)
			db.SetMaxIdleConns(10)
			db.SetConnMaxLifetime(30 * time.Minute)

			if err := db.Ping(); err != nil {
				return nil, err
			}
		}
		return db, nil
	}
}
