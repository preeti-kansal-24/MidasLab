package repository

import (
	"context"
	"preeti-kansal-24/MidasLab.git/schema"
)

type AuthStore interface {
	SignUp(ctx context.Context, profile *schema.UserProfile) error
	Get(ctx context.Context, searchString string, value string) (*schema.UserProfile, error)
}
