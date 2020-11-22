package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type IApplication interface {
	Build() IApplication
	Run()

	BuildGrpc() IApplication
	RunGrpc() error
}

type Application struct {
	logger logrus.Logger
	engine *gin.Engine
	db     *gorm.DB

	grpcserver      *grpc.Server
}
