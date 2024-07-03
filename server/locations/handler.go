package locations

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"taxi_order_service/models"
)

type Handler struct {
	LocationService ILocationService
}

type location struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type ILocationService interface {
	StoreLocation(
		ctx context.Context,
		point models.Point,
		userId int,
	) (err error)
}

func (h *Handler) SaveLocation(w http.ResponseWriter, r *http.Request) {
	var body location

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Errorf("bad request: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	point := models.NewPoint(body.Latitude, body.Longitude)
	//TODO прокинуть юзерайди в StoreLocation
	userId, err := strconv.Atoi(r.Header.Get("X-User-Id"))
	if err != nil {
		fmt.Errorf("no user: %w", err)
	}
	if err := h.LocationService.StoreLocation(r.Context(), point, userId); err != nil {
		fmt.Errorf("internal server error: %w", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
