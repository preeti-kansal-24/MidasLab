package service

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"preeti-kansal-24/MidasLab.git/clients"
	"preeti-kansal-24/MidasLab.git/constants"
	"preeti-kansal-24/MidasLab.git/domain/repository"
	"preeti-kansal-24/MidasLab.git/domain/service"
	"preeti-kansal-24/MidasLab.git/schema"
)

type otpService struct {
	as repository.AuthStore
	os repository.OTPStore
}

func NewOtpService(as repository.AuthStore, os repository.OTPStore) service.OtpService {
	return &otpService{as: as, os: os}
}

func (o otpService) GenerateOtp(ctx context.Context, phoneNumber string) (*schema.Otps, error) {
	logging := log.WithFields(log.Fields{"method": "OtpService.GenerateOtp", "phoneNumber": phoneNumber})
	logging.Info("Generating OTP for user")

	profile, err := o.as.Get(ctx, constants.PhoneNumberSearchStr, phoneNumber)
	if err != nil {
		logging.WithFields(log.Fields{"error": err}).Warn("Error getting user profile")
		return nil, err
	}
	if profile == nil {
		logging.Warn("User profile not found")
		return nil, errors.New("user profile not found")
	}
	_, otp, _ := clients.GenerateOTP(phoneNumber)
	saveOtp := &schema.Otps{UserProfileId: profile.Id, Otp: otp}
	err = o.os.Save(saveOtp)
	if err != nil {
		logging.WithFields(log.Fields{"error": err}).Warn("Error saving OTP")
		return nil, err
	}
	return saveOtp, nil
}
