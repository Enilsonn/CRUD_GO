package handlers

import (
	"net/http"
	"strconv"

	"github.com/Enilsonn/CRUD_GO/api/tools"
	"github.com/go-chi/chi/v5"
)

func (bd *BD) MethodDeleteByID() http.HandlerFunc {
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
		}

		if _, ok := bd.Users[id]; ok {
			delete(bd.Users, id)
			tools.SendJson(
				w,
				tools.Response{
					Data: "User deleted",
				},
				http.StatusOK,
			)
			return
		}

		tools.SendJson(
			w,
			tools.Response{
				Error: "User doesn't exist",
			},
			http.StatusNotFound,
		)
	}
}
