package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Enilsonn/CRUD_GO/api/tools"
	"github.com/go-chi/chi/v5"
)

func (bd *BD) MethodUpdateByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			tools.SendJson(
				w,
				tools.Response{
					Error: "Invalid id parameter",
				},
				http.StatusBadRequest,
			)
			return
		}

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			tools.SendJson(
				w,
				tools.Response{
					Error: "Invalid body",
				},
				http.StatusUnprocessableEntity,
			)
			return
		}

		if _, ok := bd.Users[id]; !ok {
			tools.SendJson(
				w,
				tools.Response{
					Error: "User id does not exist",
				},
				http.StatusNotFound,
			)
		}

		bd.Users[id] = user

		tools.SendJson(
			w,
			tools.Response{
				Data: user,
			},
			http.StatusOK,
		)
	}
}
