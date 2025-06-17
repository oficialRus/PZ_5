package main

import (
	"github.com/gin-gonic/gin"

	"cadet_project/controllers"
	"cadet_project/services"
)

func main() {
	r := gin.Default()

	// Cadet
	cadetSvc := services.NewCadetService()
	controllers.NewCadetController(r, cadetSvc)

	// Equipment – новая сущность
	equipSvc := services.NewEquipmentService()
	controllers.NewEquipmentController(r, equipSvc)

	r.Run(":8080") // listen :8080
}
