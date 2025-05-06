package historysession

type LoginHistory struct {
	LogID      int    `json:"log_id"`
	Token      string `json:"token"`
	UserID     int    `json:"user_id"`
	Success    bool   `json:"success"`
	IPAddress  string `json:"ip_address"`
	UserAgent  string `json:"user_agent"`
	DeviceInfo string `json:"device_info"`
	Action     string `json:"action"` // "login" o "logout"
}
