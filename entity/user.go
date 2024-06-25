package entity

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar;not null" json:"name"`
	Email     string    `gorm:"type:varchar;uniqueIndex;not null" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
