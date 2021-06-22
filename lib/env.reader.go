package lib

import (
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "go-api-automation-example"

//Loads the environment variable file (this was added because we are using testify)
func LoadEnv() {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		panic(err)
	}
}

//returns specified environment variable
func GetEnvVariable(key string) string {
	if os.Getenv(key) == "" {
		panic(key + " env variable not found!")
	}
	return os.Getenv(key)
}
