package app

import (
	"github.com/kritmet/go-task/app/config"
	"github.com/kritmet/go-task/app/database"
	"github.com/kritmet/go-task/docs"
)

// Application application
type Application struct {
	Config       config.Configs
	ReturnResult config.ReturnResult
}

// New new application
func New() (*Application, error) {
	err := config.InitConfig()
	if err != nil {
		return nil, err
	}

	err = config.InitReturnResult("configs")
	if err != nil {
		return nil, err
	}

	app := &Application{
		Config:       config.CF,
		ReturnResult: config.RR,
	}

	docs.SwaggerInfo.Title = config.CF.Swagger.Title
	docs.SwaggerInfo.Description = config.CF.Swagger.Description

	conf := &database.Configuration{
		Host:     config.CF.Database.Host,
		Username: config.CF.Database.Username,
		Password: config.CF.Database.Password,
		Name:     config.CF.Database.Name,
	}
	session, err := database.New(conf)
	if err != nil {
		return nil, err
	}
	database.Set(session.Database)

	return app, nil
}
