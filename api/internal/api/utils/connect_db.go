package utils

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("example.db"), &gorm.Config{})

	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados")
		return nil, err
	}

	return db, nil
}
