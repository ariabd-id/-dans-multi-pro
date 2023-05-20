package main

import (
	"dans-multi-pro/config"
	"dans-multi-pro/controllers"
	"dans-multi-pro/middlewares"
	"dans-multi-pro/repositories"
	"dans-multi-pro/services"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	route := gin.Default()

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// user
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// job
	jobRepo := repositories.NewJobRepo(db)
	jobService := services.NewJobService(jobRepo)
	jobController := controllers.NewJobController(jobService)

	// route
	route.POST("/register", userController.UserRegister)
	route.POST("/login", userController.Login)

	authRouter := route.Group("/auth")
	{
		authRouter.Use(middlewares.CORSMiddleware())
		authRouter.POST("/register", userController.UserRegister)
		authRouter.POST("/login", userController.Login)
	}

	jobRouter := route.Group("/job")
	{
		jobRouter.Use(middlewares.Auth())
		jobRouter.GET("/list", jobController.GetJobList)
		jobRouter.GET("/detail/:detailID", jobController.GetJobDetail)
	}

	route.Run(config.APP_PORT)
}
