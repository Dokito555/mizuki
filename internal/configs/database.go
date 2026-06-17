package configs

import (
	"fmt"
	"time"

	"github.com/Dokito555/mizuki/internal/entities"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_SSLMODE"),
	)

	gormLogLevel := logger.Silent
	if viper.GetString("APP_ENV") == "development" {
		gormLogLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(gormLogLevel),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(viper.GetInt("DB_MAX_IDLE"))
	sqlDB.SetMaxOpenConns(viper.GetInt("DB_MAX_OPEN"))
	sqlDB.SetConnMaxLifetime(time.Duration(viper.GetInt("DB_MAX_LIFETIME_MIN")) * time.Minute)

	if err := db.AutoMigrate(
		&entities.Upload{},
		&entities.Flow{},
		&entities.FlowPacketSample{},
	); err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	db.Exec(`CREATE INDEX IF NOT EXISTS idx_flows_src_dst ON flows (src_ip, dst_ip, src_port)`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_flows_raw_file ON flows (raw_file_id, first_seen)`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_flows_protocol ON flows (protocol, first_seen)`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_flows_threat_score ON flows (score DESC) WHERE score > 0`)

	log.Info("database connected and migrated")
	return db
}

