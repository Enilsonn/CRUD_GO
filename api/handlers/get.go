package handlers

import (
	"net/http"
	"strconv"

	"github.com/Enilsonn/CRUD_GO/api/tools"
	"github.com/go-chi/chi/v5"
)

func (bd *BD) MethodReadByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			tools.SendJson(
				w,
				tools.Response{
					Error: "Invalid id parameter",
				},
				http.StatusBadRequest)
			return
		}

		data, ok := bd.Users[id]
		if !ok {
			tools.SendJson(
				w,
				tools.Response{
					Error: "invalid-url",
				},
				http.StatusNotFound,
			)
			return
		}

		tools.SendJson(
			w,
			tools.Response{
				Data: data,
			},
			http.StatusOK,
		)
	}
}

func (bd *BD) MethodReadAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []User
		for _, user := range bd.Users {
			users = append(users, user)
		}

		tools.SendJson(
			w,
			tools.Response{
				Data: users,
			},
			http.StatusOK,
		)
	}
}
