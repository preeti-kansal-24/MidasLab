package schema

type Otps struct {
	TimeEntity
	UserProfile   *UserProfile `gorm:"foreignKey:UserProfileId"`
	UserProfileId uint32       `gorm:"not null;index"`
	Otp           string       `gorm:"not null"`
}
