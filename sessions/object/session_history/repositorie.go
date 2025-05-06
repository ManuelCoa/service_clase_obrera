package historysession

import (
	"database/sql"

	"claseobrera/sessions/repository"
)

func selectLoginHistory() (*sql.Rows, error) {
	query := "SELECT log_id, token, id_user, success, ip_address, user_agent, device_info, action FROM user_login_history"
	return repository.FetchRows(query)
}

func selectLoginHistoryID(id int) (*sql.Rows, error) {
	query := "SELECT log_id, token, id_user, success, ip_address, user_agent, device_info, action FROM user_login_history WHERE log_id = ?"
	return repository.FetchRows(query, id)
}

func insertLoginHistory(data LoginHistory) (sql.Result, error) {
	query := "INSERT INTO user_login_history (token, id_user, success, ip_address, user_agent, device_info, action) VALUES (?, ?, ?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.Token, data.UserID, data.Success, data.IPAddress, data.UserAgent, data.DeviceInfo, data.Action)
}
