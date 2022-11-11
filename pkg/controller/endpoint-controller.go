// Package controller contains ...
package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/connect2naga/logger/logging"
	"github.com/gorilla/mux"
)

/*
Author : Nagarjuna S
Date : 30-04-2022 18:18
Project : sample-http-service
File : endpoint-controller.go
*/

type EmployeeDetails struct {
	Id        string
	Name      string
	Locations string
}

type EndpointHandler struct {
	logger          logging.Logger
	EmployeeDetails map[string]EmployeeDetails
}

func NewEndpointHandler(logger logging.Logger) *EndpointHandler {
	return &EndpointHandler{logger: logger, EmployeeDetails: make(map[string]EmployeeDetails)}
}
func (e *EndpointHandler) Status(w http.ResponseWriter, r *http.Request) {
	e.logger.Infof(context.Background(), "endpoint hit......")
	w.WriteHeader(http.StatusOK)
}

func (e *EndpointHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	e.logger.Infof(context.Background(), "GetAllEmployees hit......")
	data, err := json.Marshal(e.EmployeeDetails)
	if err != nil {
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

	vars := mux.Vars(r)
	empId := vars["id"]

	empDetails, ok := e.EmployeeDetails[empId]
	if !ok {
		fmt.Printf("no data availale...")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("given EmpID %s not found", empId)))
		return
	}

	data, err := json.Marshal(empDetails)
	if err != nil {
		fmt.Printf("failed to marshl...")
		w.Write([]byte("error while marshling data"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func (e *EndpointHandler) SelectEmployee(w http.ResponseWriter, r *http.Request) {
	e.logger.Infof(context.Background(), "GetAllEmployees hit......")
	//params := mux.Vars(r)
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	fmt.Printf("--------> %s", string(bodyBytes))
	var emp EmployeeDetails

	json.Unmarshal(bodyBytes, &emp)

	json.NewEncoder(w).Encode(emp)
}

func (e *EndpointHandler) CreateEmployees(w http.ResponseWriter, r *http.Request) {
	e.logger.Infof(context.Background(), "GetAllEmployees hit......")
	var emp EmployeeDetails
	_ = json.NewDecoder(r.Body).Decode(&emp)
	fmt.Println(emp)
	//rand.Seed(time.Now().UnixNano())
	key := strconv.Itoa(rand.Intn(1000000000))
	e.EmployeeDetails[key] = emp
	json.NewEncoder(w).Encode(&emp)
	w.Write([]byte("Success"))
}

//  for _, value := range e.EmployeeDetails{
//  	if  == e.EmployeeDetails{
// 		value(emp)

// 	}

//  }

// func (e *EndpointHandler) SelectById(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	e.logger.Infof(context.Background(), "GetAllEmployees hit......")

// 	params := mux.Vars(r)

// 	for _, item := range e.EmployeeDetails {
// 		if item.Id == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return

// 		}
// 	}
//}

// func (e *EndpointHandler) PutEmployees(w http.ResponseWriter, r *http.Request) {
// 	var emp = []EmployeeDetails{}
// 	p := EmployeeDetails{}
// 	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
// 		fmt.Println(err)
// 		http.Error(w, "Error decodng response ", http.StatusBadRequest)
// 		return
// 	}
// 	emp = append(emp, p)
// 	fmt.Println(p)
// 	fmt.Println((emp))
// 	//  sample := e.EmployeeDetails
// 	//fmt.Println(sample)
// 	//  sample=map["1"]p
// 	e.EmployeeDetails = map[string]EmployeeDetails{"1": p}
// 	//  sample := map[EmployeeDetails]int{p: 1}
// 	fmt.Println("Map is ", e.EmployeeDetails)
// 	response, err := json.Marshal(&p)
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, "Error encoding response ", http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write(response)
// 	fmt.Fprintf(w, "Success")
// }
