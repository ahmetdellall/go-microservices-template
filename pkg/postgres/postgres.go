package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	DBName   string `yaml:"dbName"`
	SSLMode  bool   `yaml:"sslMode"`
	Password string `yaml:"password"`
}

const (
	maxConn         = 50
	maxConnIdleTime = 1 * time.Minute
	maxConnLifetime = 3 * time.Minute
)

func Init(cfg *Config) (*gorm.DB, error) {

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DBName,
		cfg.Password,
	)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetConnMaxIdleTime(maxConnIdleTime)
	sqlDB.SetConnMaxLifetime(maxConnLifetime)

	if err != nil {
		return nil, err
	}
	return db, nil
}
