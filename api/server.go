package api

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)

}

type Server struct {
	listenAddr string
}

func NewApiServer(listenAddr string) *Server {
	// create a new server
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {

	http.HandleFunc("/", s.GetAllFruits)

	return http.ListenAndServe(s.listenAddr, nil)
}

// methods
func (s *Server) GetAllFruits(w http.ResponseWriter, r *http.Request) {

	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 34
	fruits["Bananas"] = 15
	fruits["Strawberry"] = 18

	WriteJSON(w, http.StatusOK, fruits)
}
