package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load("config/.env")

var (
	ConnectionString = fmt.Sprintf("server=%s; user id=%s;password=%s; database=%s",
		os.Getenv("host"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("database"))
)

const AllowedCORSDomain = "http://localhost"
