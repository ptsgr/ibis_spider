package storage

import "time"

type Run struct {
	ID        int       `gorm:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;"`
}
