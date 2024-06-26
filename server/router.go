package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"taxi_order_service/server/locations"
)

type RouterBuilder struct {
	LocationHandler *locations.Handler
}

func (b *RouterBuilder) Build() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/api/v1/locations", b.LocationHandler.SaveLocation)
	return r
}
