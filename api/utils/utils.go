package utils

import (
	"fmt"
	"main/database/database_models"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ParseDBConfig(path string) (database_models.DatabaseConfig, error) {

	var config database_models.DatabaseConfig

	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	printStruct("DatabaseConfig", config)

	return config, nil

}

func GetDatabaseConnection(config database_models.DatabaseConfig) (gorm.DB, error) {
	var connection *gorm.DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.User, config.Password, config.Name, config.Port)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return *connection, err
	}
	return *connection, nil

}

func printStruct(name string, data interface{}) {
	fmt.Printf("%s: %+v\n", name, data)
}
