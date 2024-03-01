// main.go

package main

import (
	"fmt"
	"log"

	"github.com/example/config"
	"github.com/example/controller"
	"github.com/example/middleware"
	"github.com/example/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin engine
	router := gin.Default()

	// Initialize database connection
	db, err := config.SetupDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Initialize services
	adminService := &service.AdminService{DB: db}
	patientService := &service.PatientService{DB: db}

	// Initialize controllers
	adminController := controller.NewAdminController(adminService)
	patientController := controller.NewPatientController(patientService)

	// Apply middleware
	router.Use(middleware.AuthMiddleware())

	// Define routes
	api := router.Group("/api")
	{
		adminRoutes := api.Group("/admin")
		{
			adminRoutes.POST("/register", adminController.RegisterAdmin)
			adminRoutes.POST("/login", adminController.LoginAdmin)
			adminRoutes.GET("/patients", adminController.ListPatients)
			// Add other admin routes as needed
		}

		patientRoutes := api.Group("/patient")
		{
			patientRoutes.POST("/register", patientController.RegisterPatient)
			patientRoutes.POST("/login", patientController.LoginPatient)
			// Add other patient routes as needed
		}
	}

	// Start the server
	port := ":8080"
	fmt.Println("Server is running on port", port)
	err = router.Run(port)
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
