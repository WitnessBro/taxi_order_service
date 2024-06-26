package locations

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"taxi_order_service/models"
)

type Handler struct {
	LocationService iLocationService
}

type location struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type iLocationService interface {
	StoreLocation(
		ctx context.Context,
		point models.Point,
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
	if err := h.LocationService.StoreLocation(context.Background(), point); err != nil {
		fmt.Errorf("internal server error: %w", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
