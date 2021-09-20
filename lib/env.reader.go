package lib

import (
	"os"
)

//returns specified environment variable
func GetEnvVariable(key string) string {
	if os.Getenv(key) == "" {
		panic(key + " env variable not found!")
	}
	return os.Getenv(key)
}
