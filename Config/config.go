// config.go

package Config

import (
	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	Database struct {
		Host     string `mapstructure:"DB_HOST"`
		Port     int    `mapstructure:"DB_PORT"`
		User     string `mapstructure:"DB_USER"`
		Password string `mapstructure:"DB_PASSWORD"`
		Name     string `mapstructure:"DB_NAME"`
	} `mapstructure:"DATABASE"`

	Server struct {
		Port int `mapstructure:"SERVER_PORT"`
	} `mapstructure:"SERVER"`

	JWT struct {
		SecretKey string `mapstructure:"JWT_SECRET_KEY"`
	} `mapstructure:"JWT"`
}

var AppConfig Config

// LoadConfig loads configuration from file and environment variables
func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// You can add additional configuration sources like environment variables here

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return err
	}

	return nil
}
