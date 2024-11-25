package main

import (
	"github.com/dapr-platform/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
	"strconv"
	_ "visit-service/docs"
)

var (
	PORT = 80
)

func init() {

	if val := os.Getenv("LISTEN_PORT"); val != "" {
		PORT, _ = strconv.Atoi(val)
	}
	common.Logger.Debug("use PORT ", PORT)
}

// @title visit-service API
// @version 1.0
// @description visit-service API
// @BasePath /swagger/visit-service
func main() {
	mux := chi.NewRouter()
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/swagger*", httpSwagger.WrapHandler)
	s := daprd.NewServiceWithMux(":"+strconv.Itoa(PORT), mux)
	common.Logger.Debug("server start")
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		common.Logger.Fatalf("error: %v", err)
	}
}
