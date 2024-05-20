package transformer

import (
	"context"
	"preeti-kansal-24/MidasLab.git/domain/transformer"
	proto "preeti-kansal-24/MidasLab.git/proto/src/go"
	"preeti-kansal-24/MidasLab.git/schema"
)

type authTransformer struct {
}

func NewAuthTransformer() transformer.AuthTransformer {
	return &authTransformer{}
}

func (a authTransformer) ReqToUserModel(ctx context.Context, req *proto.SignupWithPhoneNumberRequest) *schema.UserProfile {
	return &schema.UserProfile{
		Name:    req.Name,
		Email:   req.Email,
		PhoneNo: req.PhoneNumber,
	}
}

func (a authTransformer) ModelToProfileResp(ctx context.Context, model *schema.UserProfile) *proto.GetProfileResponse {
	return &proto.GetProfileResponse{
		Id:          model.Id,
		Name:        model.Name,
		Email:       model.Email,
		PhoneNumber: model.PhoneNo,
	}
}
