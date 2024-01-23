package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DBConnectionString is the string used to connect with MySQL
	DBConnectionString string
	
	// PORT used to run the API
	Port int
)

// Load will initialize env variables
func Load() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}
	
	DBConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", 
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_NAME"),
	)
}