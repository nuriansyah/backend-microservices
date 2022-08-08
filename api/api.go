package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nuriansyah/log-mbkm-unpas/repository"
)

type API struct {
	mhsRepo   repository.MahasiswaRepository
	dosenRepo repository.DosenRepository
	logRepo   repository.LogRepository
	router    *gin.Engine
}

func NewAPi(mhsRepo repository.MahasiswaRepository, dosenRepo repository.DosenRepository, logRepo repository.LogRepository) API {
	router := gin.Default()
	api := API{
		mhsRepo:   mhsRepo,
		dosenRepo: dosenRepo,
		logRepo:   logRepo,
		router:    router,
	}

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.POST("/login", api.login)
	router.POST("/register", api.register)

	return api

}

func (api *API) Handler() *gin.Engine {
	return api.router
}

func (api *API) Start() {
	api.Handler().Run()
}
