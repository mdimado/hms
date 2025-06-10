package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"hospital-management/internal/auth"
	"hospital-management/internal/config"
	"hospital-management/internal/database"
	"hospital-management/internal/patient"
	"hospital-management/internal/user"
)

// @title Hospital Management System API
// @version 1.0
// @description A simple hospital management system with receptionist and doctor portals
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize config
	cfg := config.Load()

	// Initialize database
	db, err := database.Initialize(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Initialize repositories
	userRepo := user.NewRepository(db)
	patientRepo := patient.NewRepository(db)

	// Initialize services
	userService := user.NewService(userRepo)
	patientService := patient.NewService(patientRepo)
	authService := auth.NewService(userService, cfg.JWTSecret)

	// Initialize handlers
	authHandler := auth.NewHandler(authService)
	userHandler := user.NewHandler(userService)
	patientHandler := patient.NewHandler(patientService)

	// Setup router
	router := gin.Default()

	// CORS middleware
	router.Use(cors.Default())

	// // Static files
	// router.Static("/static", "./web/static")
	// router.LoadHTMLGlob("web/templates/*")

	// // Frontend routes
	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "login.html", nil)
	// })

	// router.GET("/dashboard", auth.RequireAuth(cfg.JWTSecret), func(c *gin.Context) {
	// 	userInterface := c.MustGet("user")
	// 	user := userInterface.(map[string]interface{})
	// 	c.HTML(http.StatusOK, "dashboard.html", gin.H{
	// 		"user": user,
	// 	})
	// })

	// API routes
	v1 := router.Group("/api/v1")
	{
		// Auth routes
		v1.POST("/login", authHandler.Login)
		v1.POST("/register", authHandler.Register)

		// Protected routes
		protected := v1.Group("/")
		protected.Use(auth.RequireAuth(cfg.JWTSecret))
		{
			// User routes
			protected.GET("/profile", userHandler.GetProfile)
			protected.PUT("/profile", userHandler.UpdateProfile)

			// Patient routes (Receptionist only)
			receptionist := protected.Group("/patients")
			receptionist.Use(auth.RequireRole("receptionist"))
			{
				receptionist.POST("/", patientHandler.CreatePatient)
				receptionist.GET("/", patientHandler.GetPatients)
				receptionist.GET("/:id", patientHandler.GetPatient)
				receptionist.PUT("/:id", patientHandler.UpdatePatient)
				receptionist.DELETE("/:id", patientHandler.DeletePatient)
			}

			// Doctor routes
			doctor := protected.Group("/doctor")
			doctor.Use(auth.RequireRole("doctor"))
			{
				doctor.GET("/patients", patientHandler.GetPatients)
				doctor.GET("/patients/:id", patientHandler.GetPatient)
				doctor.PUT("/patients/:id/medical-info", patientHandler.UpdateMedicalInfo)
			}
		}
	}

	// Swagger documentation
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}