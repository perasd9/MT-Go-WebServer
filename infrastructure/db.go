package database

import "gorm.io/gorm"

type Db interface {
	GetDb() *gorm.DB
}
