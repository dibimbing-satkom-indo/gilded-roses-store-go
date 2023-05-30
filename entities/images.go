package entities

import "time"

type Images struct {
	ID        uint      `gorm:"primaryKey;type:bigint unsigned auto_increment" json:"id"`
	Url       string    `json:"url"`
	ItemID    uint      `gorm:"type:bigint unsigned" json:"itemId"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// untuk override nama table
func (i Images) TableName() string {
	return "images"
}
