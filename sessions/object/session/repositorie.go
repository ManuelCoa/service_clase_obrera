package session

import (
	"database/sql"

	"claseobrera/sessions/repository"
)

func selectSessionGeneral() (*sql.Rows, error) {
	query := "SELECT session_token, id_user, ip_address, user_agent, device_info, expiration_time FROM gestion_user_sessions"
	return repository.FetchRows(query)
}

func selectSessionGeneralID(id int) (*sql.Rows, error) {
	query := "SELECT session_token, id_user, ip_address, user_agent, device_info, expiration_time FROM gestion_user_session WHERE session_id = ?"
	return repository.FetchRows(query, id)
}

func insertSessionGeneral(data Session) (sql.Result, error) {
	query := "INSERT INTO gestion_user_session (session_token, id_user, ip_address, user_agent, device_info, expiration_time) VALUES (?, ?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.Token, data.UserID, data.IPAddress, data.UserAgent, data.DeviceInfo, data.Expiration)
}

func deleteSessionGeneral(token string) (sql.Result, error) {
	query := "DELETE FROM gestion_user_session WHERE session_token = ?"
	return repository.ExecuteQuery(query, token)
}
