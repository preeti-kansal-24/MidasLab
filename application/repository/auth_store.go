package repository

import (
	"context"
	"gorm.io/gorm"
	"preeti-kansal-24/MidasLab.git/domain/repository"
	"preeti-kansal-24/MidasLab.git/schema"
)

type authStore struct {
	db *gorm.DB
}

func NewAuthStore(db *gorm.DB) repository.AuthStore {
	return &authStore{db: db}
}

func (a *authStore) SignUp(ctx context.Context, profile *schema.UserProfile) error {
	tx := a.db.WithContext(ctx).Model(&profile).Save(profile)
	if tx.RowsAffected == 0 {
		return a.db.WithContext(ctx).Create(profile).Error
	}
	return tx.Error
}

func (a *authStore) Get(ctx context.Context, searchString string, value string) (*schema.UserProfile, error) {
	rsp := schema.UserProfile{}
	err := a.db.WithContext(ctx).Where(searchString, value).Find(&rsp).Error
	return &rsp, err
}
