package locations_test

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"taxi_order_service/models"
	"taxi_order_service/server"
	"taxi_order_service/server/locations"
	"taxi_order_service/server/locations/mocks"
	"testing"
)

func TestHandler_SaveLocation(t *testing.T) {
	iLocationService := mocks.NewILocationService(t)
	iLocationService.EXPECT().StoreLocation(mock.Anything, models.NewPoint(50.452, 45.3434), 42).Return(nil)
	locationHandler := locations.Handler{
		LocationService: iLocationService,
	}
	routerBuilder := server.RouterBuilder{
		LocationHandler: &locationHandler,
	}
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/locations",
		strings.NewReader(`{"latitude":50.452, "longitude": 45.3434}`))
	req.Header.Set("X-User-Id", "42")
	w := httptest.NewRecorder()
	routerBuilder.Build().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	require.EqualValues(t, http.StatusOK, resp.StatusCode)
	require.EqualValues(t, "", body)
}

func TestHandler_SaveLocation_Internal_Server_Error(t *testing.T) {
	iLocationService := mocks.NewILocationService(t)
	iLocationService.EXPECT().StoreLocation(mock.Anything, models.NewPoint(50.452, 45.3434), 42).Return(errors.New("lolkekcheburek"))
	locationHandler := locations.Handler{
		LocationService: iLocationService,
	}
	routerBuilder := server.RouterBuilder{
		LocationHandler: &locationHandler,
	}
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/locations",
		strings.NewReader(`{"latitude":50.452, "longitude": 45.3434}`))
	req.Header.Set("X-User-Id", "42")
	w := httptest.NewRecorder()
	routerBuilder.Build().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	require.EqualValues(t, http.StatusInternalServerError, resp.StatusCode)
	require.EqualValues(t, "internal server error\n", string(body))
}
