package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func LoadEnv() *Config {

/*
* err := godotenv.Load() .env ফাইল থেকে সব environment variables লোড করে OS environment এ সেট করে দেয়।
* এরপর তুমি os.Getenv("KEY") দিয়ে সেগুলো ব্যবহার করতে পারো।
* যদি .env না পায় বা error হয়, তাহলে err এ error থাকে।
*/

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found, using system environment variables")
	}

	return &Config{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DB_URL"),
	}
}
