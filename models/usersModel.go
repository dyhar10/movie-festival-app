package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UUID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	Role      string    `gorm:"default:user" json:"role"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
