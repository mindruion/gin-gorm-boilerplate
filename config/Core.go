package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

var EnvConfigs map[string]string

func LoadEnv() {
	EnvConfigs, err = godotenv.Read()

	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
}
