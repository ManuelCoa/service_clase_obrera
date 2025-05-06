package connection

import (
	"database/sql"

	"claseobrera/persistence/driver"
	"claseobrera/persistence/enviroment"
)

var db *sql.DB

/*
ConnectionFactory es una función que devuelve dos funciones: `open` y `close`, ambas diseñadas para
gestionar la conexión a una base de datos MySQL de manera controlada a través de una interfaz de orden superior.
La función `open` utiliza una fábrica de conexiones que crea una instancia de conexión a MySQL (si aún no existe)
y la reutiliza en llamadas posteriores.
La función `close` se encarga de cerrar la conexión a la base de datos si está abierta y restablecer el valor
de la variable global 'db' a nil para permitir futuras reconexiones.
El acceso a estas dos operaciones está encapsulado, permitiendo su uso exclusivamente a través de esta interfaz.
*/

func ConnectionFactory() (func() (*sql.DB, error), func() error) {

	dsn := enviroment.BuildDSN()

	connectFunc := driver.CreateConnectionFactory(dsn)

	open := func() (*sql.DB, error) {
		return connectFunc()
	}

	close := func() error {
		if db != nil {
			err := db.Close()
			if err != nil {
				return err
			}
			db = nil
		}
		return nil
	}

	return open, close
}
