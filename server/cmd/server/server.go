package main

import (
	"context"
	"fmt"
	"github.com/invictoprojects/architecture-lab-3/server/balancers"
	"github.com/invictoprojects/architecture-lab-3/server/machines"
	"net/http"
)

type HttpPortNumber int

type BalancersApiServer struct {
	Port HttpPortNumber

	BalancersHandler balancers.HttpHandlerFunc

	MachinesHandler machines.HttpHandlerFunc

	server *http.Server
}

func (s *BalancersApiServer) Start() error {
	if s.BalancersHandler == nil {
		return fmt.Errorf("balancers HTTP handler is not defined - cannot start")
	}
	if s.MachinesHandler == nil {
		return fmt.Errorf("machines HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/balancers", s.BalancersHandler)
	handler.HandleFunc("/machines", s.MachinesHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *BalancersApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
