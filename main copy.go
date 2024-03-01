// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	// Initialize the database connection
	initDB()
	router := gin.Default()

	// Admin routes
	adminRoutes := router.Group("/admin")
	{
		adminRoutes.POST("/register", registerAdmin)
		adminRoutes.POST("/login", loginAdmin)

		patientsRoutes := adminRoutes.Group("/patients")
		{
			patientsRoutes.GET("", listPatients)
			patientsRoutes.POST("", createPatient)
			patientsRoutes.PUT("/:patientID", updatePatient)
			patientsRoutes.DELETE("/:patientID", deletePatient)
		}
	}

	// Patient routes
	router.POST("/patient/login", loginPatient)

	// Run the server
	router.Run(":8080")
}

// Initialize the database connection
func initDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/hospitals?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database")
		panic(err)
	}

	// Auto Migrate
	db.AutoMigrate(&Admin{}, &Patient{})
}

// Model definitions
type Admin struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Patient struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Handlers
func registerAdmin(c *gin.Context) {
	// Implement admin registration logic
	c.JSON(http.StatusOK, gin.H{"message": "Admin registered successfully"})
}

func loginAdmin(c *gin.Context) {
	// Implement admin login logic
	c.JSON(http.StatusOK, gin.H{"message": "Admin logged in successfully"})
}

func listPatients(c *gin.Context) {
	// Implement logic to list patients
	c.JSON(http.StatusOK, gin.H{"message": "List of patients"})
}

func createPatient(c *gin.Context) {
	// Implement logic to create a new patient
	c.JSON(http.StatusOK, gin.H{"message": "Patient created successfully"})
}

func updatePatient(c *gin.Context) {
	// Implement logic to update a patient
	c.JSON(http.StatusOK, gin.H{"message": "Patient updated successfully"})
}

func deletePatient(c *gin.Context) {
	// Implement logic to delete a patient
	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}

func loginPatient(c *gin.Context) {
	// Implement patient login logic
	c.JSON(http.StatusOK, gin.H{"message": "Patient logged in successfully"})
}
