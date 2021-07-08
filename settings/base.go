package settings

import (
	"fmt"

	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

var configPaths = []string{
	".", // For local development
}

// AppSettings contains application related settings
type AppSettings struct {
	Generic  GenericSettings
	Database DatabaseSettings
}

const (
	//SecretFilename file to get keys and configurations
	SecretFilename = "secrets"
	//DatabasePath to read database configurations
	DatabasePath = "database"
)

func readSettings(path string, subSettings interface{}, required bool) {
	// Check if the category is required or not
	if required && !viper.IsSet(path) {
		panic(fmt.Sprintf("Missing settings %s", path))
	}
	cfg := viper.Sub(path)

	if err := cfg.Unmarshal(subSettings); err != nil {
		panic(err)
	}
	// var validate *validator.Validate
	var validate = validator.New()
	if err := validate.Struct(subSettings); err != nil {
		panic(err)
	}
	defaults.SetDefaults(subSettings)
}

var globalSettings AppSettings

func init() {
	// add the paths from the configPath
	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}
	// Read the secrets file. A secret file will store all application configuration which
	// have very high security visibility, like passwords, secret keys.
	viper.SetConfigName(SecretFilename)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// Merge config and secrets
	if err := viper.MergeInConfig(); err != nil {
		panic("Failed to merge config files")
	}
}

// GetSettings retrieves the application settings
func GetSettings() *AppSettings {
	return &globalSettings
}
