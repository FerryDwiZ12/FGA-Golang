package datasource

import (
	"fmt"
	"log"

	"github.com/FerryDwiZ12/FGA-Golang/tree/master/assigments/config"
	"github.com/FerryDwiZ12/FGA-Golang/tree/master/assigments/internal/order"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
		return nil, err
	}

	// Migrate data
	log.Println("Migrate schema...")
	db.AutoMigrate(&order.Order{}, &order.Item{})
	log.Println("Migrate Done.")

	log.Println("Successfully Connect to Database")
	return db, nil
}
