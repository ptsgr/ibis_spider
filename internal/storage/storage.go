package storage

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Storage interface {
	CreateRun() (int, error)
	SetURLStatus(runId int, url string, status urlStatus) error
}

type PostgresStorage struct {
	Db *gorm.DB
}

func NewGormStorage(dsn string) (Storage, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{Db: db}, nil
}

func (ps *PostgresStorage) CreateRun() (int, error) {
	r := Run{}
	err := ps.Db.Save(&r).Error
	if err != nil {
		return 0, err
	}

	return r.ID, nil
}
func (ps *PostgresStorage) SetURLStatus(runId int, url string, status urlStatus) error {
	urlEntity := Url{
		RunID: runId,
		Url:   url,
		State: status,
	}

	if err := ps.Db.Create(&urlEntity).Error; err != nil {
		return err
	}
	return nil

}
