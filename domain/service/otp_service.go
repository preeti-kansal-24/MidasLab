package service

import (
	"context"
	"preeti-kansal-24/MidasLab.git/schema"
)

type OtpService interface {
	GenerateOtp(ctx context.Context, phoneNumber string) (*schema.Otps, error)
}
