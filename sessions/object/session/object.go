package session

import "time"

type Session struct {
	SessionID  int       `json:"session_id"`
	Token      string    `json:"token"`
	UserID     int       `json:"user_id"`
	IPAddress  string    `json:"ip_address"`
	UserAgent  string    `json:"user_agent"`
	DeviceInfo string    `json:"device_info"`
	Expiration time.Time `json:"expiration"`
}
