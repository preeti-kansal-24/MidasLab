package transformer

import (
	"context"
	proto "preeti-kansal-24/MidasLab.git/proto/src/go"
	"preeti-kansal-24/MidasLab.git/schema"
)

type AuthTransformer interface {
	ReqToUserModel(ctx context.Context, req *proto.SignupWithPhoneNumberRequest) *schema.UserProfile
	ModelToProfileResp(ctx context.Context, model *schema.UserProfile) *proto.GetProfileResponse
}
