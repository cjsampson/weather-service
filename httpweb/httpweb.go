package httpweb

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cjsampson/weather-service/api"
	"github.com/cjsampson/weather-service/app/conf"
)

type Handler struct {
	WeatherAPI api.WeatherAPI
}

func NewHandler(config conf.WeatherConfig) http.Handler {
	wapi := api.NewWeatherAPI(config)

	return &Handler{wapi}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	resp, err := h.WeatherAPI.Get(context.Background(), req)
	if err != nil {
		errResp := ErrorResponse{Error: "bad request"}
		respBytes, rerr := json.Marshal(errResp)
		if rerr != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(respBytes)
		return
	}

	successResp := SuccessResponse{Message: "success", Data: resp}
	respBytes, rerr := json.Marshal(successResp)
	if rerr != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}
