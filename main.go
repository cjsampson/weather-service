package main

import (
	"net/http"

	"github.com/cjsampson/weather-service/app/conf"
	"github.com/cjsampson/weather-service/httpweb"
	"github.com/cjsampson/weather-service/httpweb/server"
	plog "github.com/cjsampson/weather-service/platform/logger"
)

func main() {
	logger := plog.New(conf.LoadLogConfig())

	weatherConf := conf.LoadWeatherConfig()

	handler := httpweb.NewHandler(weatherConf)

	mux := http.NewServeMux()
	mux.HandleFunc("/weather", handler.ServeHTTP)

	serve := server.NewServerState(mux, logger)
	if err := serve.Start(); err != nil {
		logger.Error("server.StartServer error - %v\n", err)
	}
}
