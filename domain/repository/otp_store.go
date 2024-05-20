package repository

import (
	"context"
	"preeti-kansal-24/MidasLab.git/schema"
)

type OTPStore interface {
	Save(model *schema.Otps) error
	Verify(ctx context.Context, model *schema.Otps) bool
	Get(ctx context.Context, model *schema.Otps) (*schema.Otps, error)
}
