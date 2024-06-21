package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type ApiServer struct {
	svc Service
}

func newApiServer(svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (s *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/", s.handleGetFact)
	return http.ListenAndServe(listenAddr, nil)
}

func (s *ApiServer) handleGetFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.svc.getCatFact(context.Background())
	if err != nil {
		writeJson(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	writeJson(w, http.StatusOK, fact)
}

func writeJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
