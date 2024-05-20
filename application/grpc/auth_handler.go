package grpc

import (
	"context"
	"github.com/micro/go-micro/v2/errors"
	log "github.com/sirupsen/logrus"
	"preeti-kansal-24/MidasLab.git/domain/service"
	"preeti-kansal-24/MidasLab.git/domain/transformer"
	proto "preeti-kansal-24/MidasLab.git/proto/src/go"
	"preeti-kansal-24/MidasLab.git/utils"
)

type authHandler struct {
	proto.UnimplementedAuthServiceServer
	s service.AuthService
	t transformer.AuthTransformer
}

func NewAuthHandler(s service.AuthService, t transformer.AuthTransformer) proto.AuthServiceServer {
	return &authHandler{s: s, t: t}
}

func (a *authHandler) SignupWithPhoneNumber(ctx context.Context, req *proto.SignupWithPhoneNumberRequest) (*proto.SignupWithPhoneNumberResponse, error) {
	logging := log.WithContext(ctx).WithFields(log.Fields{"method": "SignupWithPhoneNumber"})

	phoneNumber, err := utils.ValidateAndFormatPhoneNumber(req.PhoneNumber, "India")
	if err != nil {
		logging.Warnf("Please provide valid phone number %v", err)
		return nil, err
	}
	userProfile := a.t.ReqToUserModel(ctx, req)

	userProfile.PhoneNo = phoneNumber

	err = a.s.SignUpWithPhoneNumber(ctx, userProfile)
	if err != nil {
		return nil, err
	}
	return &proto.SignupWithPhoneNumberResponse{UserId: userProfile.Id}, nil
}

func (a *authHandler) VerifyNumber(ctx context.Context, req *proto.VerifyNumberReq) (*proto.VerifyNumberResp, error) {
	logging := log.WithContext(ctx).WithFields(log.Fields{"method": "VerifyNumber"})

	phoneNumber, err := utils.ValidateAndFormatPhoneNumber(req.PhoneNumber, "India")
	if err != nil {
		logging.Warnf("Please provide valid phone number %v", err)
		return nil, err
	}

	if len(req.Otp) != 6 {
		logging.Warn("Please provide valid otp")
		return nil, errors.BadRequest("bad_request", "Please provide valid otp")
	}

	err = a.s.VerifyPhoneNumber(ctx, phoneNumber, req.GetOtp())
	if err != nil {
		return nil, err
	}
	return &proto.VerifyNumberResp{Message: "valid otp"}, nil
}

func (a *authHandler) Login(ctx context.Context, req *proto.VerifyNumberReq) (*proto.VerifyNumberResp, error) {
	return a.VerifyNumber(ctx, req)
}

func (a *authHandler) GetProfile(ctx context.Context, in *proto.GetProfileReq) (*proto.GetProfileResponse, error) {
	logging := log.WithContext(ctx).WithFields(log.Fields{"method": "GetProfile"})

	phoneNumber, err := utils.ValidateAndFormatPhoneNumber(in.PhoneNumber, "India")
	if err != nil {
		logging.Warnf("Please provide valid phone number %v", err)
		return nil, err
	}
	profile, err := a.s.GetProfile(ctx, phoneNumber)
	if err != nil {
		return nil, err
	}
	return a.t.ModelToProfileResp(ctx, profile), nil

}
