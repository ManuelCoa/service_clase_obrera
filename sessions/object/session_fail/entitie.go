package failsession

import (
	"database/sql"

	"claseobrera/sessions/entitie"
)

func scanDato(rows *sql.Rows, dato *FailedLogin) error {
	return rows.Scan(&dato.AttemptID, &dato.IPAddress, &dato.UserAgent, &dato.AttemptedUsername)
}

func searchFailedLogin() ([]FailedLogin, error) {
	rows, err := selectLoginFailed()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanDato)
}

func searchFailedLoginID(id int) ([]FailedLogin, error) {
	rows, err := selectLoginFailedID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanDato)
}

func insertionFailedLogin(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[FailedLogin](data)
	if err != nil {
		return err
	}

	_, err = insertLoginFailed(dato)
	if err != nil {
		return err
	}

	return nil
}
