package query

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type TimeSeriesPoint struct {
	Timestamp time.Time
	//Timestamp time.Time `gorm:"uniqueIndex"`
	Subject string
	Value   float64
}

func CreateSchema() {
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_USER", "postgres")
	_ = os.Setenv("DB_PASS", "postgres")
	_ = os.Setenv("DB_NAME", "postgres")
	_ = os.Setenv("DB_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s;", "time_series_point"))

	// Migrate the schema
	err = db.AutoMigrate(&TimeSeriesPoint{})
	if err != nil {
		panic("failed to migrate database")
	} else {
		log.Default().Printf("ok migrated")
	}

	// convert to timescale
	db.Exec(fmt.Sprintf("SELECT create_hypertable('%s', '%s');", "time_series_points", "timestamp"))
	log.Default().Printf("ok hypertable created")

	// Create
	db.Create(&TimeSeriesPoint{Timestamp: time.Now(), Subject: "ABC-CLOSE", Value: 123.0})
	log.Default().Printf("ok created")
	// Read
	var product TimeSeriesPoint
	db.First(&product, "subject = ?", "ABC-CLOSE") // find product with code D42
	log.Default().Printf("ok read")
	// Update
	db.Model(&product).Where("subject = ?", "ABC-CLOSE").Update("value", 200)
	log.Default().Printf("ok update")
	// Updates
	db.Model(&product).Where("subject = ?", "ABC-CLOSE").Updates(map[string]interface{}{"subject": "ABC-OPEN"})
	log.Default().Printf("ok updates")
	// Delete - delete product
	db.Delete(&product, map[string]interface{}{"subject": "ABC-OPEN"})
	log.Default().Printf("ok delete")
}
