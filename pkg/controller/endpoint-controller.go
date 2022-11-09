// Package controller contains ...
package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/connect2naga/logger/logging"
	"github.com/gorilla/mux"
)

/*
Author : Nagarjuna S
Date : 30-04-2022 18:18
Project : sample-http-service
File : endpoint-controller.go
*/

type EmployeeDetails struct{
	Id string
	Name string
	Locations string
}

type EndpointHandler struct {
	logger logging.Logger
	EmployeeDetails map[string]EmployeeDetails
}

func NewEndpointHandler(logger logging.Logger) *EndpointHandler {
	return &EndpointHandler{logger: logger,EmployeeDetails : make(map[string]EmployeeDetails) }
}
func (e *EndpointHandler) Status(w http.ResponseWriter, r *http.Request) {
	e.logger.Infof(context.Background(), "endpoint hit......")
	w.WriteHeader(http.StatusOK)
}


func (e *EndpointHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	e.logger.Infof(context.Background(), "GetAllEmployees hit......")
	data, err := json.Marshal(e.EmployeeDetails)
	if err != nil{
		fmt.Printf("failed to marshl...")
		w.Write([]byte("error while fetching data"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func (e *EndpointHandler) GetAllEmployeeById(w http.ResponseWriter, r *http.Request) {
	e.logger.Infof(context.Background(), "GetAllEmployeeById hit......")

	vars  := mux.Vars(r)
	empId := vars["id"]


	empDetails, ok :=  e.EmployeeDetails[empId]
	if !ok {
		fmt.Printf("no data availale...")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("given EmpID %s not found",empId )))
		return
	}

	data, err := json.Marshal(empDetails)
	if err != nil{
		fmt.Printf("failed to marshl...")
		w.Write([]byte("error while marshling data"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

