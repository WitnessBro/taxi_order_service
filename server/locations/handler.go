package locations

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"taxi_order_service/models"
	"taxi_order_service/services"
)

type Handler struct {
	LocationService ILocationService
	KafkaProducer   *services.KafkaProducer
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
	userId, err := strconv.Atoi(r.Header.Get("X-User-Id"))
	if err != nil {
		fmt.Errorf("user not authorized: %w", err)
		http.Error(w, "user not authorized", http.StatusUnauthorized)
	}
	if err := h.LocationService.StoreLocation(r.Context(), point, userId); err != nil {
		fmt.Errorf("internal server error: %w", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	message := fmt.Sprintf("%f,%f", body.Latitude, body.Longitude)
	if err := h.KafkaProducer.WriteMessage(r.Context(), []byte(strconv.Itoa(userId)), []byte(message)); err != nil {
		fmt.Errorf("could not write message to Kafka: %w", err)
		http.Error(w, "could not write message to Kafka", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
