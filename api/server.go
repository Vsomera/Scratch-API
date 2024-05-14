package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/vsomera/scratch-api/storage"
	"github.com/vsomera/scratch-api/types"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewApiServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/getFruits", s.handleGetAllFruits) // GET
	http.HandleFunc("/addFruit", s.handleAddFruit)      // POST
	http.HandleFunc("/editFruit/", s.handleEditFruit)   // PUT ( /editFruit/{id} )
	return http.ListenAndServe(s.listenAddr, nil)
}

// HANDLERS //

func (s *Server) handleGetAllFruits(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		// get all fruits in the database
		fruits, err := s.store.GetAllFruits()
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		WriteJSON(w, http.StatusOK, fruits)
		return
	}

	WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})

}

func (s *Server) handleAddFruit(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		// decoding body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "could not read request body"})
			return
		}

		// check if request matches expected request type
		var request types.AddFruitRequest
		err = json.Unmarshal(body, &request) // decode into the struct
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body format"})
			return
		}

		// adding fruit to database
		err = s.store.AddFruit(request.Name, request.Count)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, "could not add fruit")
			return
		}
		WriteJSON(w, http.StatusOK, map[string]string{"message": "fruit added to database"})
		return
	}

	WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
}

func (s *Server) handleEditFruit(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "PUT":

		// extract id value from url
		urlParts := strings.Split(r.URL.Path, "/")
		fruitId := urlParts[len(urlParts)-1]
		if fruitId == "" {
			errMsg := fmt.Sprintf("invalid URL missing /%s/{id}", urlParts[1])
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": errMsg})
			return
		}

		// decoding body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "could not read request body"})
			return
		}

		// check if request matches expected request type
		var request types.EditFruitRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		// edit the fruit in the database
		err = s.store.EditFruit(fruitId, request.Count)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		WriteJSON(w, http.StatusOK, map[string]string{"message": "fruit count updated"})
		return
	}

	WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
}

// HELPER //

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
