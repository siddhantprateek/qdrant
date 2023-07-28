package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnviron(var_name string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Unable to load .env file")
	}
	variable := os.Getenv(var_name)
	return variable
}
