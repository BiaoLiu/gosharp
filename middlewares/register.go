package middlewares

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gosharp/config"
)

func Register(engine *gin.Engine) {
	//设置cors
	engine.Use(CorsMiddleware())
	//全局错误处理
	engine.Use(ExceptionMiddleware())

	//设置session
	//store := sessionStore()
	//engine.Use(sessions.Sessions(auth.SessionName, store))
	//自定义中间件
	//engine.Use(SessionMiddleware(store))
}

func sessionStore() sessions.Store {
	fmt.Println("REDIS_ADDRESS:", config.Viper.GetString("REDIS_ADDRESS"), "REDIS_PASSWORD:", config.Viper.GetString("REDIS_PASSWORD"))
	store, err := sessions.NewRedisStore(10, "tcp", config.Viper.GetString("REDIS_ADDRESS"), config.Viper.GetString("REDIS_PASSWORD"), []byte("secret"))
	if err != nil {
		panic(fmt.Sprintf("session store init error. %s", err.Error()))
	}
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 3, //过期时间3天
	})
	return store
}
