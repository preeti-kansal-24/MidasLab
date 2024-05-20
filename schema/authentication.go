package schema

import "time"

type UserProfile struct {
	TimeEntity
	Id      uint32 `gorm:"primaryKey"`
	Name    string `gorm:"size:255"`
	PhoneNo string `gorm:"size:255;unique"`
	Email   string `gorm:"size:255"`
}

type TimeEntity struct {
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
