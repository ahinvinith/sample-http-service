// Package sample-service contains ...
package main

import (
	"context"
	"github.com/connect2naga/logger/logging"
	"gitlab.com/tariandev_intelops/sample-http-service/pkg/server"
)

/*
Author : Nagarjuna S
Date : 30-04-2022 18:04
Project : sample-http-service
File : main.go
*/

var logger = logging.NewLogger()

func main() {
	s, err := server.NewServer(logger)
	if err != nil {
		logger.Fatalf(context.Background(), "unable to create the server instance, error: %v", err)
	}
	s.Start()
}
