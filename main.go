package main

import (
	"github.com/kritmet/go-task/api/route"
	"github.com/kritmet/go-task/app"
	"github.com/sirupsen/logrus"
)

func main() {
	application, err := app.New()
	if err != nil {
		logrus.Panic(err)
	}

	route.Setup(application)
}
