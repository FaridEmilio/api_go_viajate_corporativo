package config

import (
	"fmt"
	"os"
)

var (
	USER  = "lolo"
	AUTH  = "lolo"
	DB    = "lolo"
	PASSW = "lolo"
	HOST  = "lolo"
	PORT  = "lolo"
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}
	if fallback != "" {
		return fallback
	}
	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
