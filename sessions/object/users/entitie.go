package users

import (
	"database/sql"

	"claseobrera/sessions/entitie"
)

func scanDato(rows *sql.Rows, dato *Usuarios) error {
	return rows.Scan(&dato.UserID, &dato.RolesID, &dato.InstitucionID, &dato.Nombre, &dato.User, &dato.Password)
}

func searchParsedUsuarios() ([]Usuarios, error) {
	rows, err := selectUsuarios()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanDato)
}

func searchParsedUsuariosID(id int) ([]Usuarios, error) {
	rows, err := selectUsuariosID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanDato)
}

func searchParsedUsuariosUSER(user string) ([]Usuarios, error) {
	rows, err := selectUsuariosUSER(user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanDato)
}

func insertionUsuarios(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Usuarios](data)
	if err != nil {
		return err
	}

	_, err = insertUsuarios(dato)
	if err != nil {
		return err
	}

	return nil
}

func putUsuarios(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Usuarios](data)
	if err != nil {
		return err
	}

	_, err = updateUsuarios(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionUsuarios(id int) error {
	_, err := deleteUsuarios(id)
	if err != nil {
		return err
	}

	return nil
}
