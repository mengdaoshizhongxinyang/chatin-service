package global

import "time"

type BASE_MODEL struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
