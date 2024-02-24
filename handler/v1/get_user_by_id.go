package v1

import (
	"encoding/json"
	"net/http"

	"github.com/mercy34mercy/go-http-server/model/user"
	"github.com/mercy34mercy/go-http-server/usecase"
)

func GetUserHandler(usecase usecase.GetUserByID) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("user_id")
		if userID == "" {
			http.Error(w, "user_id is required", http.StatusBadRequest)
			return
		}
		user, err := usecase.Execute(user.UserID(userID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		userJSON, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(userJSON)
	}
}
