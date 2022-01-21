package envaccess

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvironment will load a specific .env file
// based on the assumed "env" environment variable
func LoadEnvironment() {
	env := os.Getenv("env")

	godotenv.Load(".env." + env)
}

// GetKey returns an environment key
func GetKey(key string) string {
	return os.Getenv(key)
}
