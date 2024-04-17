package entities

import (
	"time"
)

// entity or model

type User struct {
	ID        uint      `gorm:"primaryKey"` // Standard field for the primary key
	Name      string    // A regular string field
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}
