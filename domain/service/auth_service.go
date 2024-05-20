package service

import (
	"context"
	"preeti-kansal-24/MidasLab.git/schema"
)

type AuthService interface {
	SignUpWithPhoneNumber(ctx context.Context, profile *schema.UserProfile) error
	VerifyPhoneNumber(ctx context.Context, phone string, otp string) error
	GetProfile(ctx context.Context, phone string) (*schema.UserProfile, error)
}
