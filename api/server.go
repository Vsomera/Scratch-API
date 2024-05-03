package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/vsomera/scratch-api/storage"
	"github.com/vsomera/scratch-api/types"
)

type Server struct {
	listenAddr string
	store      *storage.MySqlStorage
}

func NewApiServer(listenAddr string, store *storage.MySqlStorage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/fruit", s.handleGetFruitByName) // Example route
	http.HandleFunc("/addFruit", s.handleAddFruit)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetFruitByName(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		body, err := io.ReadAll(r.Body)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "could not read request body"})
			return
		}

		var request types.GetFruitByNameRequest
		err = json.Unmarshal(body, &request) // decode into the struct
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body format"})
			return
		}

		fruitName := request.Name

		// find the fruit in the database
		fruit, err := s.store.GetFruitByName(fruitName)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "could not find fruit"})
			return
		}

		WriteJSON(w, http.StatusOK, fruit)
		return
	}

	WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})

}

func (s *Server) handleAddFruit(w http.ResponseWriter, r *http.Request) {
	// TODO : use request body to specify fruit and count
	err := s.store.AddFruit("bananas", 8)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, "could not add fruit")
	}
	WriteJSON(w, http.StatusOK, "added")
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
