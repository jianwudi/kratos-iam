package service

import (
	"context"

	v1 "github.com/jianwudi/apiserver/api/helloworld/v1"
	"github.com/jianwudi/apiserver/internal/biz"
)

// ApiserverService is a apiserver service.
type ApiserverService struct {
	v1.UnimplementedApiserverServer

	uc *biz.ApiserverUsecase
}

// NewApiserverService new a apiserver service.
func NewApiserverService(uc *biz.ApiserverUsecase) *ApiserverService {
	return &ApiserverService{uc: uc}
}

// SayHello implements helloworld.ApiserverServer.
func (s *ApiserverService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateApiserver(ctx, &biz.Apiserver{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
