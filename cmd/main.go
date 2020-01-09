package main

import (
	"context"
	"flag"
	"gosharp/internal/server/http"
	"gosharp/internal/service"
	"gosharp/library/conf/viper"
	"gosharp/library/log"
	"gosharp/library/validation"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9000
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	flag.Parse()
	config.Init("configs")
	log.Init("logs") // debug flag: log.dir={path}

	validation.SetValidationMessage()
	//defer log.Close()
	log.Logger.Info("gosharp start")
	//ecode.Init(nil)
	svc := service.New()
	httpSrv := http.New(svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Logger.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			defer cancel()
			httpSrv.Shutdown(ctx)
			log.Logger.Info("gosharp exit")
			svc.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
