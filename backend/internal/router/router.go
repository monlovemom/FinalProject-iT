package router

import (
	"database/sql"
	"net/http"

	"finalproject-it/backend/internal/handlers"
	"finalproject-it/backend/internal/middleware"
	"finalproject-it/backend/internal/repositories"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://127.0.0.1:3000",
			"http://localhost",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/health", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "Unhealthy"})
			return
		}

		c.JSON(200, gin.H{"status": "GooD"})
	})

	userRepo := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	// public
	r.POST("/api/register", userHandler.Register)
	r.POST("/api/login", userHandler.Login)

	// user
	auth := r.Group("/api")
	auth.Use(middleware.Auth())
	{
		auth.GET("/me", userHandler.Me)
	}

	// admin | owner
	admin := r.Group("/api")
	admin.Use(middleware.Auth(), middleware.RoleRequired("admin", "owner"))
	{
		admin.GET("/users", userHandler.GetAll)
		admin.GET("/users/:id", userHandler.GetByID)
		admin.PUT("/users/:id", userHandler.Update)
		admin.DELETE("/users/:id", userHandler.Delete)
	}

	return r
}
