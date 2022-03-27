// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/jianwudi/apiserver/internal/biz"
	"github.com/jianwudi/apiserver/internal/conf"
	"github.com/jianwudi/apiserver/internal/data"
	"github.com/jianwudi/apiserver/internal/server"
	"github.com/jianwudi/apiserver/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	apiserverRepo := data.NewApiserverRepo(dataData, logger)
	apiserverUsecase := biz.NewApiserverUsecase(apiserverRepo, logger)
	apiserverService := service.NewApiserverService(apiserverUsecase)
	httpServer := server.NewHTTPServer(confServer, apiserverService, logger)
	grpcServer := server.NewGRPCServer(confServer, apiserverService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
