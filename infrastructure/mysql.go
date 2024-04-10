package database

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlDb struct {
	Db *gorm.DB
}

func (m *MysqlDb) GetDb() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err.Error())
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err.Error())
	}

	m.Db = db

	return db
}
