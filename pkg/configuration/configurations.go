// Package configuration contains ...
package configuration

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

/*
Author : Nagarjuna S
Date : 30-04-2022 18:27
Project : sample-http-service
File : configurations.go
*/

type ServiceConfigurations struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
	Port     string `envconfig:"PORT" default:"50050"`
}

func GetServiceConfigurations() (serviceConf *ServiceConfigurations, err error) {
	serviceConf = &ServiceConfigurations{}
	if err = envconfig.Process("", serviceConf); err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
