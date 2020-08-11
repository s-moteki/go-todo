package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const targetEnvName = "GO_ENV"

// EnvLoad env load
func EnvLoad() {
	if "" == os.Getenv(targetEnvName) {
		_ = os.Setenv(targetEnvName, "local")
	}
	err := godotenv.Load(fmt.Sprintf("../envfiles/%s.env", os.Getenv(targetEnvName)))

	if err != nil {
		log.Fatalf("Error loading env target env is %s", os.Getenv(targetEnvName))
	}
}
