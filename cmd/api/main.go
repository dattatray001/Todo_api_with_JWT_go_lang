package main

import (
	"log"
	"todo_api/internal/config"
	"todo_api/internal/database"
	"todo_api/internal/handlers"
	"todo_api/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Declare a pointer to hold application configuration
	var cfg *config.Config
	var err error

	// Load configuration values (env variables, config files, etc.)
	cfg, err = config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Declare a PostgreSQL connection pool
	var pool *pgxpool.Pool

	// Initialize database connection using DATABASE_URL from config
	pool, err = database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Ensure database connections are properly closed
	// This runs when the main function exits
	defer pool.Close()

	// setup ruoter using gon framework
	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)
	//  method for test server init
	router.GET("/", func(c *gin.Context) {
		// map[string]interface{}
		// map[string]any{}
		c.JSON(200, gin.H{
			"message":  "Todo API is running well!",
			"status":   "success",
			"database": "connected",
		})
	})

	router.POST("/auth/register", handlers.CreateUserHandler(pool))
	router.POST("/auth/login", handlers.LoginHandler(pool, cfg))

	// router.POST("/todos", handlers.CreateTodoHandler(pool))
	// router.GET("/todos", handlers.GetAllTodosHandler(pool))
	// router.GET("/todos/:id", handlers.GetToDoByIDHandler(pool))
	// router.POST("/todos/:id", handlers.UpdateToDoHandler(pool))
	// router.DELETE("/todos/:id", handlers.DeleteToDoHandler(pool))

	protected := router.Group("/todos")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		protected.POST("", handlers.CreateTodoHandler(pool))
		protected.GET("", handlers.GetAllTodosHandler(pool))
		protected.GET("/:id", handlers.GetToDoByIDHandler(pool))
		protected.PUT("/:id", handlers.UpdateToDoHandler(pool))
		protected.DELETE("/:id", handlers.DeleteToDoHandler(pool))
	}

	// Middleware Test Route
	router.GET("/protected-test", middleware.AuthMiddleware(cfg), handlers.TestProtectedHandler())

	// defind the port number
	router.Run(":" + cfg.Port)

}
