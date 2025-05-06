package enviroment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error cargando el archivo .env")
	}
}

func GetDBUser() string {
	return os.Getenv("DB_USER")
}

func GetDBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func GetDBHost() string {
	return os.Getenv("DB_HOST")
}

func GetDBName() string {
	return os.Getenv("DB_NAME")
}

func BuildDSN() string {
	user := GetDBUser()
	password := GetDBPassword()
	host := GetDBHost()
	dbName := GetDBName()

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbName)
}
