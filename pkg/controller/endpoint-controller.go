// Package controller contains ...
package controller

import (
	"context"
	"github.com/connect2naga/logger/logging"
	"net/http"
)

/*
Author : Nagarjuna S
Date : 30-04-2022 18:18
Project : sample-http-service
File : endpoint-controller.go
*/

type EndpointHandler struct {
	logger logging.Logger
}

func NewEndpointHandler(logger logging.Logger) *EndpointHandler {
	return &EndpointHandler{logger: logger}
}
func (e *EndpointHandler) Status(w http.ResponseWriter, r *http.Request) {
	e.logger.Infof(context.Background(), "endpoint hit......")
	w.WriteHeader(http.StatusOK)
}
