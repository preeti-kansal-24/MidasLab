package grpc

import (
	"context"
	log "github.com/sirupsen/logrus"
	"preeti-kansal-24/MidasLab.git/domain/service"
	proto "preeti-kansal-24/MidasLab.git/proto/src/go"
	"preeti-kansal-24/MidasLab.git/utils"
)

type otpHandler struct {
	proto.UnimplementedOtpServiceServer
	os service.OtpService
}

func NewOtpHandler(os service.OtpService) proto.OtpServiceServer {
	return &otpHandler{
		os: os,
	}
}

func (oh *otpHandler) GenerateOtp(ctx context.Context, req *proto.GenerateOtpReq) (*proto.GenerateOtpResp, error) {
	logger := log.WithFields(log.Fields{"method": "GenerateOtp", "phone_number": req.PhoneNumber})
	phoneNumber, err := utils.ValidateAndFormatPhoneNumber(req.PhoneNumber, "India")
	if err != nil {
		logger.Warnf("Please provide valid phone number %v", err)
		return nil, err
	}

	otp, err := oh.os.GenerateOtp(ctx, phoneNumber)
	if err != nil {
		return nil, err
	}
	return &proto.GenerateOtpResp{
		PhoneNumber: phoneNumber,
		Otp:         otp.Otp,
	}, nil

}
