package app

import (
	GrpcApp "SSO/internal/app/grpc"
	"SSO/internal/config"
	"SSO/internal/service/apps"
	"SSO/internal/service/auth"
	"SSO/internal/storage"
	"log/slog"
)

type App struct {
	GRPCApp *GrpcApp.App
}

func New(l *slog.Logger, cnf *config.Config) *App {
	s, err := storage.New(&cnf.DBConfig)
	if err != nil {
		panic(err)
	}

	authService := auth.New(l, s.UserStorage, s.AppStorage, cnf.TokenTTL)
	appsService := apps.New(l, s.AppStorage)

	grpcApp := GrpcApp.New(l, authService, appsService, &cnf.BindConfig)
	return &App{
		GRPCApp: grpcApp,
	}
}
