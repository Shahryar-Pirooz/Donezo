package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnectOption struct {
	User     string
	Password string
	Host     string
	Port     uint
	DBName   string
	SSLMode  string
}

func (opt DBConnectOption) setDSN() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", opt.Host, opt.User, opt.Password, opt.DBName, opt.Port, opt.SSLMode)
	return dsn
}

func NewConnection(opt DBConnectOption) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(opt.setDSN()), &gorm.Config{Logger: logger.Discard})
}
