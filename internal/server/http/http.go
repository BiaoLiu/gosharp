package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "gosharp/docs"
	"gosharp/internal/server/http/middleware"
	"gosharp/internal/service"
	"gosharp/library/conf/viper"
	"gosharp/library/log"
	"net/http"
	"time"
)

var (
	svc *service.Service
)

type ApiServer struct {
	server *http.Server
}

// New new a bm server.
func New(s *service.Service) (httpSrv *ApiServer) {
	svc = s
	engine := gin.Default()
	initRouter(engine)
	endPoint := fmt.Sprintf(":%d", config.Viper.GetInt("PORT"))
	server := &http.Server{
		Addr:         endPoint,
		Handler:      engine,
		ReadTimeout:  time.Duration(config.Viper.GetDuration("ReadTimeout") * time.Second),
		WriteTimeout: time.Duration(config.Viper.GetDuration("WriteTimeout") * time.Second),
		//MaxHeaderBytes: maxHeaderBytes,
	}
	httpSrv = &ApiServer{
		server: server,
	}
	if err := httpSrv.Start(); err != nil {
		panic(err)
	}
	return
}

func initRouter(e *gin.Engine) {
	//设置cors
	e.Use(middleware.CorsMiddleware())
	//全局错误处理
	e.Use(middleware.ExceptionMiddleware())

	e.GET("/ping", ping)
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g := e.Group("/x/comment")
	{
		g.GET("/start", howToStart)
	}

	e.POST("/login", Login)

	e.POST("/articles", ArticleCreate)
	e.PUT("/articles/:id", ArticleUpdate)
	e.GET("/articles", ArticleList)
	e.DELETE("/articles", ArticleDelete)
	e.GET("/articles/:id", func(c *gin.Context) {
		path1 := c.Param("id")
		if path1 == "all" {
			ArticleAllList(c)
		} else {
			ArticleRetrieve(c)
		}
	})
}

func (s *ApiServer) Start() error {
	log.Logger.Info("gin: start http listen addr ", s.server.Addr)

	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			if errors.Cause(err) == http.ErrServerClosed {
				log.Logger.Info("gin: server closed")
				return
			}
			panic(errors.Wrapf(err, "gin: engine.ListenServer(%+v)", s.server))
		}
	}()

	return nil
}

func (s *ApiServer) Shutdown(ctx context.Context) error {
	return errors.WithStack(s.server.Shutdown(ctx))
}

func ping(c *gin.Context) {
	if err := svc.Ping(c); err != nil {
		log.Logger.Error("ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *gin.Context) {
	c.String(0, "Golang 大法好 !!!")
}
