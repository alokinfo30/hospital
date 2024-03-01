// patient_controller.go

package Controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/your-gin-project/Helper"
	"github.com/yourusername/your-gin-project/Model"
	"github.com/yourusername/your-gin-project/Services"
)

// PatientController handles requests related to patient operations
type PatientController struct {
	PatientService *Services.PatientService
}

// NewPatientController creates a new instance of PatientController
func NewPatientController(patientService *Services.PatientService) *PatientController {
	return &PatientController{PatientService: patientService}
}

// LoginPatient handles the login of a patient
func (pc *PatientController) LoginPatient(c *gin.Context) {
	var loginModel Model.LoginModel

	if err := c.ShouldBindJSON(&loginModel); err != nil {
		Helper.RespondWithError(c, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	// Additional validation if needed

	// Call PatientService to perform login
	token, err := pc.PatientService.LoginPatient(loginModel.Email, loginModel.Password)
	if err != nil {
		Helper.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	Helper.RespondWithJSON(c, http.StatusOK, gin.H{"token": token})
}

// GetPatientList returns a list of all patients
func (pc *PatientController) GetPatientList(c *gin.Context) {
	patientList := pc.PatientService.GetPatientList()
	Helper.RespondWithJSON(c, http.StatusOK, patientList)
}

// CreatePatient handles the creation of a new patient
func (pc *PatientController) CreatePatient(c *gin.Context) {
	var patientModel Model.Patient

	if err := c.ShouldBindJSON(&patientModel); err != nil {
		Helper.RespondWithError(c, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	// Additional validation if needed

	// Call PatientService to create the patient
	pc.PatientService.CreatePatient(&patientModel)

	Helper.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Patient created successfully"})
}

// UpdatePatient updates details of a patient
func (pc *PatientController) UpdatePatient(c *gin.Context) {
	var updatedPatient Model.Patient

	if err := c.ShouldBindJSON(&updatedPatient); err != nil {
		Helper.RespondWithError(c, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	// Additional validation if needed

	// Call PatientService to update patient details
	pc.PatientService.UpdatePatient(&updatedPatient)

	Helper.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Patient updated successfully"})
}

// DeletePatient deletes a patient
func (pc *PatientController) DeletePatient(c *gin.Context) {
	patientID := c.Param("id")

	// Validate patientID (convert to uint, check existence, etc.)

	// Call PatientService to delete patient
	pc.PatientService.DeletePatient(patientID)

	Helper.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}
