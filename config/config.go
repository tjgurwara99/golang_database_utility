package config

import "os"

// GetConfs returns configurations strings for database connection
func GetConfs() (string, string, string, string, string) {
	switch statement := os.Getenv("environemtn"); statement {
	case "production":
		return "mysql", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), "127.0.0.1", os.Getenv("DATABASE_NAME")
	default:
		return "mysql", "crowntech", "This_is_a_test_database", "127.0.0.1", "crowntech"
	}
}
