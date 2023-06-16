package storage

type urlStatus string

const (
	StatusOK  urlStatus = "OK"
	StatusNOK urlStatus = "NOT OK"
)

type Url struct {
	ID    uint64    `gorm:"id"`
	RunID int       `gorm:"column:run_id;"`
	Url   string    `gorm:"column:url;type:varchar;size:128;"`
	State urlStatus `gorm:"column:state"`
}
