package db

import (
	"fmt"
	"log"
	"time"

	"github.com/Temisaputra/warOnk/infrastructure/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitPostgreSQL(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", cfg.DBHost, cfg.DBUsername, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // Disable statement caching
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf(err.Error())
		panic(err)
	} else {
		log.Printf("Successfully connected to database server")
	}

	rdb, err := db.DB()
	if err != nil {
		log.Fatalf(err.Error())
		panic(err)
	}

	rdb.SetMaxIdleConns(cfg.DBMaxIdleConns)
	rdb.SetMaxOpenConns(cfg.DBMaxOpenConns)
	rdb.SetConnMaxLifetime(time.Duration(int(time.Minute) * cfg.DBMaxLifetime))
	rdb.SetConnMaxIdleTime(time.Duration(int(time.Minute) * cfg.DBMaxIdleTime))

	return db
}

func DbClose(db *gorm.DB) {
	rdb, err := db.DB()
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	_ = rdb.Close()
}
