package orm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func NewDB(config *dbConfig) (*gorm.DB, error) {
	// Default gorm logger, but with extended slow threshold from 200ms to 500ms
	logger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             500 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})

	// dbParams := "charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(postgres.Open(config.url), &gorm.Config{
		Logger: logger,
	})
}
