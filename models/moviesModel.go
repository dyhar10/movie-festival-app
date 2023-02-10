package models

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	UUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    uint      `json:"duration"`
	Artist      string    `json:"artist"`
	Genre       string    `json:"genre"`
	URL         string    `json:"url"`
	Viewers     uint16    `gorm:"default:0" json:"viewers"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
