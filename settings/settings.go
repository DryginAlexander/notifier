package settings

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// db
	DBDialect string
	DBHost    string
	DBName    string
	DBPort    string
	DBUser    string
	DBPw      string
	// web
	Port   int
	APIKey string
)

func Init(dotenvFileName string) error {
	if dotenvFileName != "" {
		err := godotenv.Load(dotenvFileName)
		if err != nil {
			return err
		}
	}

	// data base settings
	DBDialect = getenv("DB_DIALECT", "sqlite")
	DBHost = os.Getenv("DB_HOST")
	DBName = getenv("DB_NAME", "local.db")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPw = os.Getenv("DB_PW")

	// web settings
	var err error
	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return err
	}
	APIKey = os.Getenv("API_KEY")

	return nil
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}