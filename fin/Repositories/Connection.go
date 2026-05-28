package repositories

import (
	"encoding/json"
	"io"
	"log"
	"os"

	m "Financiering/Models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Connection_String string `json:"connection_string"`
}

func readConfig(path string) Config {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var config Config
	json.Unmarshal(byteValue, &config)
	return config
}

var database *gorm.DB

func Database_Get() *gorm.DB {
	if database != nil {
		return database
	}

	config := readConfig("config.json")

	db, err := gorm.Open(postgres.Open(config.Connection_String), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	database = db
	err = AutoMigrateDatabaseTables()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func AutoMigrateDatabaseTables() error {
	err := database.AutoMigrate(
		&m.Budget{},
		&m.FinanceDossier{},
		&m.Invoice{})
	return err
}
