package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Enilsonn/CRUD_GO/api/tools"
)

func (bd *BD) MethodCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		r.Body = http.MaxBytesReader(w, r.Body, 3000)

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			tools.SendJson(
				w,
				tools.Response{
					Error: "invalid-body",
				},
				http.StatusUnprocessableEntity,
			)
		}

		id := bd.Len + 1
		bd.Len = id

		if _, ok := bd.Users[id]; !ok {
			bd.Users[id] = user
			tools.SendJson(
				w,
				tools.Response{
					Data: user,
				},
				http.StatusOK,
			)
			return
		}

		tools.SendJson(
			w,
			tools.Response{
				Error: "User already exists",
			},
			http.StatusBadRequest,
		)
	}
}
