package users

import (
	"database/sql"

	"claseobrera/sessions/hash"
	"claseobrera/sessions/repository"
)

func selectUsuarios() (*sql.Rows, error) {
	query := "SELECT id_user, id_roles, id_institucion, nombres_apellidos, nombre_user, password_user FROM gestion_user"
	return repository.FetchRows(query)
}

func selectUsuariosID(id int) (*sql.Rows, error) {
	query := "SELECT id_user, id_roles, id_institucion, nombres_apellidos, nombre_user, password_user FROM gestion_user WHERE id_user = ?"
	return repository.FetchRows(query, id)
}

func selectUsuariosUSER(user string) (*sql.Rows, error) {
	query := "SELECT id_user, id_roles, id_institucion, nombres_apellidos, nombre_user, password_user FROM gestion_user WHERE nombre_user = ?"
	return repository.FetchRows(query, user)
}

func insertUsuarios(data Usuarios) (sql.Result, error) {
	hashedPassword, err := hash.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}
	query := "INSERT INTO gestion_user (id_roles, id_institucion, nombres_apellidos, nombre_user, password_user) VALUES (?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.RolesID, data.InstitucionID, data.Nombre, data.User, hashedPassword)
}

func updateUsuarios(data Usuarios) (sql.Result, error) {
	hashedPassword, err := hash.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}
	query := "UPDATE gestion_user SET id_roles = ?, id_institucion = ?, nombres_apellidos = ?, nombre_user = ?, password_user = ? WHERE id_user = ?"
	return repository.ExecuteQuery(query, data.RolesID, data.InstitucionID, data.Nombre, data.User, hashedPassword, data.UserID)
}

func deleteUsuarios(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_user WHERE id_user = ?"
	return repository.ExecuteQuery(query, id)
}
