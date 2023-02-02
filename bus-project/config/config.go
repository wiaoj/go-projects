package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func GetPort() string {
	return fmt.Sprintf(":%d", viper.GetInt("PORT"))
}

func GetConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		viper.GetString("POSTGRES_HOST"),
		viper.GetString("POSTGRES_USER"),
		viper.GetString("POSTGRES_PASSWORD"),
		viper.GetString("POSTGRES_DATABASE_NAME"),
		viper.GetInt("POSGRES_PORT"),
		viper.GetString("SSL_MODE"))
}

func GetJwtSecretKey() string {
	return fmt.Sprintf("%s", viper.GetString("JWT_SECRET_KEY"))
}
func GetJwtExpirationMinutes() int {
	return viper.GetInt("JWT_EXPIRATION_MINUTES")
}
