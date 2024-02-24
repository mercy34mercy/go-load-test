package v1

import (
	"encoding/json"
	"net/http"

	"github.com/mercy34mercy/go-http-server/log"
	"github.com/mercy34mercy/go-http-server/model/user"
	"github.com/mercy34mercy/go-http-server/usecase"
)

type CreateUserHandlerInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CreateUserHandler(usecase usecase.SaveUser) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateUserHandlerInput
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			internalCreateUserHandlerError(err, "")
			http.Error(w, "failed to parse request body", http.StatusBadRequest)
			return
		}

		if req.ID == "" {
			http.Error(w, "id is required", http.StatusBadRequest)
			return
		}

		if req.Name == "" {
			http.Error(w, "name is required", http.StatusBadRequest)
			return
		}

		if req.Age <= 0 {
			http.Error(w, "age is invalid", http.StatusBadRequest)
			return
		}

		user := user.User{
			ID:   user.NewUserID(req.ID),
			Name: req.Name,
			Age:  req.Age,
		}
		err := usecase.Execute(user)
		if err != nil {
			internalCreateUserHandlerError(err, req.ID)
			http.Error(w, "failed to save user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func internalCreateUserHandlerError(err error, id string) {
	log.Errorf(err, "CreateUserHandler() error: %s, queue_id: %s", err.Error(), id)
}
