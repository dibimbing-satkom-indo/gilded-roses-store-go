package entities

import "time"

type Users struct {
	ID        uint      `gorm:"primaryKey;type:bigint unsigned auto_increment" json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// untuk override nama table
func (i Users) TableName() string {
	return "users"
}
