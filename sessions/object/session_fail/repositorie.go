package failsession

import (
	"database/sql"

	"claseobrera/sessions/repository"
)

func selectLoginFailed() (*sql.Rows, error) {
	query := "SELECT attemped_id, ip_address, user_agent, attempted_username FROM failed_login_attempts"
	return repository.FetchRows(query)
}

func selectLoginFailedID(id int) (*sql.Rows, error) {
	query := "SELECT attemped_id, ip_address, user_agent, attempted_username FROM failed_login_attempts WHERE attempts_id = ?"
	return repository.FetchRows(query, id)
}

func insertLoginFailed(data FailedLogin) (sql.Result, error) {
	query := "INSERT INTO failed_login_attempts (ip_address, user_agent, attempted_username) VALUES (?, ?, ?)"
	return repository.ExecuteQuery(query, data.IPAddress, data.UserAgent, data.AttemptedUsername)
}
