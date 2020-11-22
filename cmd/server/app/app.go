package app

import (
	"fmt"
	"github.com/bozd4g/fb.grpc/pkg/postgresql"
	"github.com/sirupsen/logrus"
	"os"
)

func NewApplication() IApplication {
	return &Application{}
}

func (application *Application) Build() IApplication {
	application.logger = *logrus.New()
	application.AddPostgreSql(postgresql.Opts{
		Host:     "localhost",
		User:     "postgres",
		Password: "123456",
		Database: "fbgrpc",
		Port:     5432,
	})

	application.AddRouter()
	application.AddControllers().InitMiddlewares().AddSwagger()

	return application
}

func (application *Application) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := application.engine.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}