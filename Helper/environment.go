package env

import "os"

func Getenv(key string, defualt string) string {
	 if value := os.Getenv(key); value != "" {
		return value
	 }

	 return defualt
}