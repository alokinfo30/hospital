// admin_controller.go

package Controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/your-gin-project/Helper"
	"github.com/yourusername/your-gin-project/Model"
	"github.com/yourusername/your-gin-project/Services"
)

// AdminController handles requests related to admin operations
type AdminController struct {
	AdminService *Services.AdminService
}

// NewAdminController creates a new instance of AdminController
func NewAdminController(adminService *Services.AdminService) *AdminController {
	return &AdminController{AdminService: adminService}
}

// RegisterAdmin handles the registration of a new admin
func (ac *AdminController) RegisterAdmin(c *gin.Context) {
	var adminModel Model.Admin

	if err := c.ShouldBindJSON(&adminModel); err != nil {
		Helper.RespondWithError(c, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	// Additional validation if needed

	// Call AdminService to register the admin
	ac.AdminService.RegisterAdmin(&adminModel)

	Helper.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Admin registered successfully"})
}

// LoginAdmin handles the login of an admin
func (ac *AdminController) LoginAdmin(c *gin.Context) {
	var loginModel Model.LoginModel

	if err := c.ShouldBindJSON(&loginModel); err != nil {
		Helper.RespondWithError(c, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	// Additional validation if needed

	// Call AdminService to perform login
	token, err := ac.AdminService.LoginAdmin(loginModel.Email, loginModel.Password)
	if err != nil {
		Helper.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	Helper.RespondWithJSON(c, http.StatusOK, gin.H{"token": token})
}

// GetAdminList returns a list of all admins
func (ac *AdminController) GetAdminList(c *gin.Context) {
	adminList := ac.AdminService.GetAdminList()
	Helper.RespondWithJSON(c, http.StatusOK, adminList)
}

// UpdateAdmin updates details of an admin
func (ac *AdminController) UpdateAdmin(c *gin.Context) {
	var updatedAdmin Model.Admin

	if err := c.ShouldBindJSON(&updatedAdmin); err != nil {
		Helper.RespondWithError(c, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	// Additional validation if needed

	// Call AdminService to update admin details
	ac.AdminService.UpdateAdmin(&updatedAdmin)

	Helper.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Admin updated successfully"})
}

// DeleteAdmin deletes an admin
func (ac *AdminController) DeleteAdmin(c *gin.Context) {
	adminID := c.Param("id")

	// Validate adminID (convert to uint, check existence, etc.)

	// Call AdminService to delete admin
	ac.AdminService.DeleteAdmin(adminID)

	Helper.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Admin deleted successfully"})
}
