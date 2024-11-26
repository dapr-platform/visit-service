package main

import (
	"net/http"
	"os"
	"strconv"
	"time"
	"visit-service/api"
	_ "visit-service/docs"
	"visit-service/service"

	"github.com/dapr-platform/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
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
	api.InitRoute(mux)
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/swagger*", httpSwagger.WrapHandler)
	s := daprd.NewServiceWithMux(":"+strconv.Itoa(PORT), mux)

	// 启动一个goroutine来初始化配置
	go func() {
		// 等待服务启动
		time.Sleep(time.Second * 2)

		// 初始化系统配置
		if err := service.InitSystemConfig(); err != nil {
			common.Logger.Errorf("Failed to initialize system config: %v", err)
		} else {
			common.Logger.Info("System config initialized successfully")
		}
	}()

	common.Logger.Debug("server start")
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		common.Logger.Fatalf("error: %v", err)
	}
}
