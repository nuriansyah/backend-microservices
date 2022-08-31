package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nuriansyah/log-mbkm-unpas/repository"
)

type API struct {
	userRepo repository.UserRepository
	postRepo repository.PostRepository
	router   *gin.Engine
}

func NewAPi(userRepo repository.UserRepository, postRepo repository.PostRepository) API {
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

	router.GET("/api/post/:id", api.readPost)
	postRouter := router.Group("/api/post", AuthMiddleware())
	{
		postRouter.POST("/", api.createPost)
	}
	router.Use()

	return api

}

func (api *API) Handler() *gin.Engine {
	return api.router
}

func (api *API) Start() {
	api.Handler().Run()
}
