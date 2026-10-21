package artikel

import (
	"time"
)

type Artikel struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title" gorm:"type:varchar(200)"`
	Content   string    `json:"content" gorm:"type:text"`
	Category  string    `json:"category" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
	Status    string    `json:"status" gorm:"type:varchar(100)"`
}
