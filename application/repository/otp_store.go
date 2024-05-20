package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"preeti-kansal-24/MidasLab.git/domain/repository"
	"preeti-kansal-24/MidasLab.git/schema"
)

type otpStore struct {
	db *gorm.DB
}

func NewOtpStore(db *gorm.DB) repository.OTPStore {
	return &otpStore{db: db}
}

func (o *otpStore) Save(otp *schema.Otps) error {
	existing := schema.Otps{}
	o.db.Model(otp).Where("user_profile_id = ?", otp.UserProfileId).
		Find(&existing)
	if existing.UserProfileId != 0 {
		tx := o.db.Model(otp).Save(otp)
		fmt.Println("Rows affected ", tx.RowsAffected)
	} else {
		fmt.Println("Creating new entry")
		return o.db.Model(otp).Create(otp).Error
	}

	return nil
}

func (o *otpStore) Verify(ctx context.Context, model *schema.Otps) bool {
	if val, _ := o.Get(ctx, model); val != nil && val.Otp != "" {
		return true
	}
	return false
}

func (o *otpStore) Get(ctx context.Context, otp *schema.Otps) (existing *schema.Otps, err error) {
	o.db.
		WithContext(ctx).
		Where(otp).
		Take(&existing)
	return
}
