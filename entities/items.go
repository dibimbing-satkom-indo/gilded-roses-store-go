package entities

import "time"

type Items struct {
	ID        uint      `gorm:"primaryKey;type:bigint unsigned auto_increment" json:"id"`
	Name      string    `json:"name"`
	Images    []Images  `gorm:"foreignKey:ItemID" json:"images"`
	SellIn    int       `gorm:"column:sell_in" json:"sellIn"`
	Quality   int       `json:"quality"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// untuk override nama table
func (i Items) TableName() string {
	return "items"
}
