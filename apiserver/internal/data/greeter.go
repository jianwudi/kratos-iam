package data

import (
	"context"

	"github.com/jianwudi/apiserver/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type apiserverRepo struct {
	data *Data
	log  *log.Helper
}

// NewApiserverRepo .
func NewApiserverRepo(data *Data, logger log.Logger) biz.ApiserverRepo {
	return &apiserverRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *apiserverRepo) Save(ctx context.Context, g *biz.Apiserver) (*biz.Apiserver, error) {
	return g, nil
}

func (r *apiserverRepo) Update(ctx context.Context, g *biz.Apiserver) (*biz.Apiserver, error) {
	return g, nil
}

func (r *apiserverRepo) FindByID(context.Context, int64) (*biz.Apiserver, error) {
	return nil, nil
}

func (r *apiserverRepo) ListByHello(context.Context, string) ([]*biz.Apiserver, error) {
	return nil, nil
}

func (r *apiserverRepo) ListAll(context.Context) ([]*biz.Apiserver, error) {
	return nil, nil
}
