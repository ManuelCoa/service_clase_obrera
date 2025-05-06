package historysession

import (
	"database/sql"

	"claseobrera/sessions/entitie"
)

func scanDato(rows *sql.Rows, dato *LoginHistory) error {
	return rows.Scan(&dato.LogID, &dato.Token, &dato.UserID, &dato.Success, &dato.IPAddress, &dato.UserAgent, &dato.DeviceInfo, &dato.Action)
}

func searchLoginHistory() ([]LoginHistory, error) {
	rows, err := selectLoginHistory()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanDato)
}

func searchLoginHistoryID(id int) ([]LoginHistory, error) {
	rows, err := selectLoginHistoryID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanDato)
}

func insertionLoginHistory(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[LoginHistory](data)
	if err != nil {
		return err
	}

	_, err = insertLoginHistory(dato)
	if err != nil {
		return err
	}

	return nil
}
