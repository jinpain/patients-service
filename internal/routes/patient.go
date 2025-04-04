package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/his-vita/patients-service/internal/controller"
)

func PatientRoutes(rg *gin.RouterGroup, patientCtrl *controller.PatientController) {
	pg := rg.Group("/patients")
	{
		pg.GET("/", patientCtrl.GetPatient)          // GET /api/v1/patients/
		pg.POST("/", patientCtrl.CreatePatient)      // POST /api/v1/patients/
		pg.PUT("/:id", patientCtrl.UpdatePatient)    // PUT /api/v1/patients/:id
		pg.DELETE("/:id", patientCtrl.DeletePatient) // DELETE /api/v1/patients/:id
	}
}
