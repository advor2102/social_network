package configs

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/advor2102/socialnetwork/internal/models"
)

var AppSettings models.Config

func ReadSettings() error {
	configFile, err := os.Open(`C:\Users\AlexanderDvornikov\golang_course\SocialNetwork\internal\configs\configs.json`)
	if err != nil {
		return fmt.Errorf("error while opening config file: %w", err)
	}
	defer configFile.Close()

	if err = json.NewDecoder(configFile).Decode(&AppSettings); err != nil {
		return fmt.Errorf("error while parsing config file: %w", err)
	}

	return nil
}
