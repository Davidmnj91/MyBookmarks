package bookmarks

import (
	"github.com/google/uuid"
	"time"
)

type Bookmark struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	URL       string    `gorm:"uniqueIndex"`
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
