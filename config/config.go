package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host             string
	User             string
	Password         string
	Database         string
	Port             string
	Sslmode          string
	GRPCADMINPORT    string
	SID              string
	TOKEN            string
	SERVICETOKEN     string
	COORDINATORAPORT string
	USERPORT         string
}

func LoadConfig() *Config {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error loading the .env file:", err)
		os.Exit(1)
	}
	var config Config

	// Use os.Getenv to retrieve environment variables
	config.Host = os.Getenv("HOST")
	config.User = os.Getenv("USER")
	config.Password = os.Getenv("PASSWORD")
	config.Database = os.Getenv("DATABASE")
	config.Port = os.Getenv("PORT")
	config.Sslmode = os.Getenv("SSLMODE")
	config.GRPCADMINPORT = os.Getenv("GRPCADMINPORT")
	config.SID = os.Getenv("SID")
	config.TOKEN = os.Getenv("TOKEN")
	config.SERVICETOKEN = os.Getenv("SERVICETOKEN")
	config.COORDINATORAPORT = os.Getenv("COORDINATORAPORT")
	config.USERPORT = os.Getenv("USERPORT")

	return &config
}
