package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"regexp"
	"strings"
	"sync"
	"time"
)

// RateLimiter es una estructura para manejar el rate limiting.
type RateLimiter struct {
	mu       sync.Mutex
	clients  map[string]int
	limit    int
	interval time.Duration
}

// NewRateLimiter crea un nuevo RateLimiter.
func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		clients:  make(map[string]int),
		limit:    limit,
		interval: interval,
	}
}

// Allow verifica si un cliente puede hacer una solicitud.
func (rl *RateLimiter) Allow(clientIP string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if count, exists := rl.clients[clientIP]; exists {
		if count >= rl.limit {
			return false
		}
		rl.clients[clientIP]++
	} else {
		rl.clients[clientIP] = 1
		go func() {
			time.Sleep(rl.interval)
			rl.mu.Lock()
			defer rl.mu.Unlock()
			delete(rl.clients, clientIP)
		}()
	}
	return true
}

// GenerateNonce genera un nonce único para proteger contra replay attacks.
func GenerateNonce(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data + time.Now().String()))
	return hex.EncodeToString(hash.Sum(nil))
}

// ValidateNonce valida un nonce para proteger contra replay attacks.
func ValidateNonce(nonce string, storedNonces map[string]bool) bool {
	if storedNonces[nonce] {
		return false
	}
	storedNonces[nonce] = true
	return true
}

// ContainsSQLInjection verifica si los datos contienen posibles inyecciones SQL.
func ContainsSQLInjection(data map[string]interface{}) bool {
	sqlKeywords := []string{"SELECT", "INSERT", "UPDATE", "DELETE", "DROP", "ALTER", "EXEC", "--", ";", "/*", "*/"}
	for _, value := range data {
		strValue, ok := value.(string)
		if !ok {
			continue
		}
		for _, keyword := range sqlKeywords {
			if strings.Contains(strings.ToUpper(strValue), keyword) {
				return true
			}
		}
	}
	return false
}

// ContainsNoSQLInjection verifica si los datos contienen posibles inyecciones NoSQL.
func ContainsNoSQLInjection(data map[string]interface{}) bool {
	nosqlKeywords := []string{"$where", "$regex", "$ne", "$lt", "$lte", "$gt", "$gte", "$in", "$nin", "$or", "$and", "$not", "$exists"}
	for key := range data {
		for _, keyword := range nosqlKeywords {
			if strings.Contains(key, keyword) {
				return true
			}
		}
	}
	return false
}

// ContainsCommandInjection verifica si los datos contienen posibles inyecciones de comandos.
func ContainsCommandInjection(data map[string]interface{}) bool {
	commandKeywords := []string{"&&", "||", ";", "`", "|", ">", "<", "&", "exec", "system", "command"}
	for _, value := range data {
		strValue, ok := value.(string)
		if !ok {
			continue
		}
		for _, keyword := range commandKeywords {
			if strings.Contains(strValue, keyword) {
				return true
			}
		}
	}
	return false
}

// ContainsXMLInjection verifica si los datos contienen posibles inyecciones XML.
func ContainsXMLInjection(data map[string]interface{}) bool {
	for _, value := range data {
		strValue, ok := value.(string)
		if !ok {
			continue
		}
		if strings.Contains(strValue, "<!DOCTYPE") || strings.Contains(strValue, "<![CDATA[") {
			return true
		}
	}
	return false
}

// ContainsJSONInjection verifica si los datos contienen posibles inyecciones JSON.
func ContainsJSONInjection(data map[string]interface{}) bool {
	for _, value := range data {
		strValue, ok := value.(string)
		if !ok {
			continue
		}
		// Ajusta la lógica para permitir ciertos patrones que sabes que son seguros
		if strings.Contains(strValue, "\"") || strings.Contains(strValue, "'") {
			var js json.RawMessage
			if err := json.Unmarshal([]byte(strValue), &js); err != nil {
				return true
			}
		}
	}
	return false
}

// ContainsBinaryData verifica si los datos contienen binarios inconsistentes.
func ContainsBinaryData(data map[string]interface{}) bool {
	for _, value := range data {
		strValue, ok := value.(string)
		if !ok {
			continue
		}
		if matched, _ := regexp.MatchString(`[^\x00-\x7F]`, strValue); matched {
			return true
		}
	}
	return false
}

// ContainsDamagedData verifica si los datos están dañados o son potencialmente peligrosos.
func ContainsDamagedData(data map[string]interface{}) bool {
	for _, value := range data {
		strValue, ok := value.(string)
		if !ok {
			continue
		}
		if strings.Contains(strValue, "<script>") || strings.Contains(strValue, "</script>") {
			return true
		}
	}
	return false
}
