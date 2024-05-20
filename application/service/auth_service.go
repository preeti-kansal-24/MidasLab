package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"preeti-kansal-24/MidasLab.git/constants"
	"preeti-kansal-24/MidasLab.git/domain/repository"
	"preeti-kansal-24/MidasLab.git/domain/service"
	"preeti-kansal-24/MidasLab.git/schema"
	"preeti-kansal-24/MidasLab.git/utils"
)

type authService struct {
	as repository.AuthStore
	os repository.OTPStore
}

func NewAuthService(authStore repository.AuthStore, os repository.OTPStore) service.AuthService {
	return &authService{as: authStore, os: os}
}

func (a *authService) SignUpWithPhoneNumber(ctx context.Context, profile *schema.UserProfile) error {
	logging := log.WithContext(ctx).WithFields(log.Fields{"method": "SignUpWithPhoneNumber", "phone-number": profile.PhoneNo})
	logging.Info("Saving user profile...")
	err := a.as.SignUp(ctx, profile)
	if err != nil {
		logging.WithFields(log.Fields{"error": err}).Warn("error saving user profile")
		return err
	}
	//generate a kafka topic for generating otp and saving to db
	marshal, err := json.Marshal(profile)
	if err != nil {

		return err
	}

	utils.PublishMessage(profile.PhoneNo, marshal, constants.GenerateOtpTopic)
	return nil
}

func (a *authService) VerifyPhoneNumber(ctx context.Context, phone string, otp string) error {
	profile, err := a.GetProfile(ctx, phone)
	if err != nil {
		return err
	}
	fmt.Println("profile is", profile)
	//verify otp
	verify := a.os.Verify(ctx, &schema.Otps{
		UserProfileId: profile.Id,
		Otp:           otp,
	})
	fmt.Println("verify is", verify)

	if verify {
		return nil
	} else {
		return errors.New("invalid otp")
	}

}

func (a *authService) GetProfile(ctx context.Context, phone string) (*schema.UserProfile, error) {
	logging := log.WithContext(ctx).WithFields(log.Fields{"method": "GetProfile", "phone-number": phone})
	logging.Info("Getting user profile...")
	searchStr := "phone_no = ?"
	profile, err := a.as.Get(ctx, searchStr, phone)
	if err != nil {
		logging.WithFields(log.Fields{"error": err}).Warn("error getting user profile")
		return profile, err
	}
	if profile == nil {
		logging.Warn("user profile not found")
		return nil, errors.New("user profile not found")
	}
	return profile, nil
}
