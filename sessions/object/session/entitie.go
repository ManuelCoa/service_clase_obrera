package session

import (
	"database/sql"

	"claseobrera/sessions/entitie"
)

func scanDato(rows *sql.Rows, dato *Session) error {
	return rows.Scan(&dato.UserID, &dato.SessionID, &dato.UserID, &dato.IPAddress, &dato.UserAgent, &dato.DeviceInfo, &dato.Expiration)
}

func searchParsedSession() ([]Session, error) {
	rows, err := selectSessionGeneral()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanDato)
}

func searchParsedSessionID(id int) ([]Session, error) {
	rows, err := selectSessionGeneralID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanDato)
}

func insertionSession(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Session](data)
	if err != nil {
		return err
	}

	_, err = insertSessionGeneral(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionSession(token string) error {
	_, err := deleteSessionGeneral(token)
	if err != nil {
		return err
	}

	return nil
}
