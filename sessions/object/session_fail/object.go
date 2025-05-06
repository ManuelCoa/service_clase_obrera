package failsession

type FailedLogin struct {
	AttemptID         int    `json:"attempt_id"`
	IPAddress         string `json:"ip_address"`
	UserAgent         string `json:"user_agent"`
	AttemptedUsername string `json:"attempted_username"`
}
