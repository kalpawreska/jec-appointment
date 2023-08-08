package dbconnect

import (
	"fmt"
	"jec-appointment/pkg/utils"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB Main database connection
var DB *gorm.DB

// InitDatabase initialize database connection
func InitDatabase() {
	if nil == DB {
		db := dbconnect()
		if nil != db {
			DB = db
		}
	}
}

func dbconnect() *gorm.DB {
	logLevel := logger.Info

	switch utils.GetEnv("ENVIRONTMENT") {
	case "staging":
		logLevel = logger.Error
	case "production":
		logLevel = logger.Silent
	}

	config := gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		// DisableForeignKeyContraintWhenMigrating: true,
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		utils.GetEnv("DB_HOST"),
		utils.GetEnv("DB_PORT"),
		utils.GetEnv("DB_USER"),
		utils.GetEnv("DB_PASS"),
		utils.GetEnv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &config)

	if nil != err {
		panic(err)
	}

	if nil != db {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(3)
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetConnMaxLifetime(time.Second * 5)
	}

	return db
}
