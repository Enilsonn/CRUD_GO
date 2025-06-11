package api

import (
	"net/http"

	"github.com/Enilsonn/CRUD_GO/api/handlers"
	"github.com/Enilsonn/CRUD_GO/api/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewMultiplier(bd *(handlers.BD)) http.Handler {

	r := chi.NewMux()

	r.Use(middlewares.SetApplicationJson)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post(
		"/",
		bd.MethodCreate(),
	)

	r.Get(
		"/{id}",
		bd.MethodReadByID(),
	)

	r.Get(
		"/",
		bd.MethodReadAll(),
	)

	r.Put(
		"/{id}",
		bd.MethodUpdateByID(),
	)

	r.Delete(
		"/{id}",
		bd.MethodDeleteByID(),
	)

	return r
}
