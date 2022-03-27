package biz

import (
	"context"

	v1 "github.com/jianwudi/apiserver/api/helloworld/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Apiserver is a Apiserver model.
type Apiserver struct {
	Hello string
}

// ApiserverRepo is a Greater repo.
type ApiserverRepo interface {
	Save(context.Context, *Apiserver) (*Apiserver, error)
	Update(context.Context, *Apiserver) (*Apiserver, error)
	FindByID(context.Context, int64) (*Apiserver, error)
	ListByHello(context.Context, string) ([]*Apiserver, error)
	ListAll(context.Context) ([]*Apiserver, error)
}

// ApiserverUsecase is a Apiserver usecase.
type ApiserverUsecase struct {
	repo ApiserverRepo
	log  *log.Helper
}

// NewApiserverUsecase new a Apiserver usecase.
func NewApiserverUsecase(repo ApiserverRepo, logger log.Logger) *ApiserverUsecase {
	return &ApiserverUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateApiserver creates a Apiserver, and returns the new Apiserver.
func (uc *ApiserverUsecase) CreateApiserver(ctx context.Context, g *Apiserver) (*Apiserver, error) {
	uc.log.WithContext(ctx).Infof("CreateApiserver: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
