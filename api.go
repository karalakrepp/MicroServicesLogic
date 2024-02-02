package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type ApiServer struct {
	listenAddr string
	svc        Service
}

func NewApiServer(listenAddr string, svc Service) *ApiServer {
	return &ApiServer{
		svc:        svc,
		listenAddr: listenAddr,
	}
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

func makeApiFunc(fn ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := fn(w, r); err != nil {
			log.Fatal(err)
		}
	}
}

func (s *ApiServer) HandleGetFact(w http.ResponseWriter, r *http.Request) error {

	fact, err := s.svc.GetCatFact(context.Background())

	if err != nil {
		return writeJson(w, http.StatusInternalServerError, err)
	}

	return writeJson(w, http.StatusOK, fact)

}

func writeJson(w http.ResponseWriter, status int, v any) error {

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
func (s *ApiServer) Run() error {

	http.HandleFunc("/catfact", makeApiFunc(s.HandleGetFact))

	return http.ListenAndServe(s.listenAddr, nil)

}
