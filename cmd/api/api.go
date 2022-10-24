package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	repository2 "github.com/nuriansyah/log-mbkm-unpas/src/repository"
	"time"
)

type API struct {
	userRepo repository2.UserRepository
	postRepo repository2.PostRepository
	router   *gin.Engine
}

func NewAPi(userRepo repository2.UserRepository, postRepo repository2.PostRepository) API {
	router := gin.Default()
	api := API{
		userRepo: userRepo,
		postRepo: postRepo,
		router:   router,
	}

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.POST("/login", api.login)
	router.POST("/register", api.register)

	//router.GET("/api/post/:id", api.readPost)
	postRouter := router.Group("/api/post", AuthMiddleware())
	{
		postRouter.POST("/create", api.createPost)
	}
	router.GET("/users", api.getProfile)
	profileRouter := router.Group("/api/profile", AuthMiddleware())
	{
		//profileRouter.GET("/", api.getProfile)
		profileRouter.PATCH("/", api.updateProfile)
		//profileRouter.PUT("/avatar", api.changeAvatar)
	}
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
			param.Keys,
		)
	}))
	router.Use(gin.Recovery())

	return api

}

func (api *API) Handler() *gin.Engine {
	return api.router
}

func (api *API) Start() {
	api.Handler().Run()
}
