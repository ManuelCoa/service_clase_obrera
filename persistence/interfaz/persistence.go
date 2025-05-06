package interfaz

import (
	"database/sql"

	"claseobrera/persistence/transaction"
)

type ReadFunc func(query string, args ...interface{}) (*sql.Rows, error)

type WriteFunc func(query string, args ...interface{}) (sql.Result, error)

type Persistence struct {
	Read  ReadFunc
	Write WriteFunc
}

func NewPersistence() Persistence {
	readFunc, writeFunc := transaction.DatabaseOperations()

	return Persistence{
		Read:  readFunc,
		Write: writeFunc,
	}
}
