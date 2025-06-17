package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"cadet_project/models"
	"cadet_project/services"
)

type EquipmentController struct {
	svc *services.EquipmentService
}

func NewEquipmentController(r *gin.Engine, svc *services.EquipmentService) {
	c := &EquipmentController{svc: svc}
	g := r.Group("/equipment")
	g.GET("", c.list)
	g.POST("", c.add)
	g.PUT("/:id", c.update)
	g.DELETE("/:id", c.delete)
}

func (c *EquipmentController) list(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.svc.List())
}

func (c *EquipmentController) add(ctx *gin.Context) {
	var e models.Equipment
	if err := ctx.ShouldBindJSON(&e); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, c.svc.Add(e))
}

func (c *EquipmentController) update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var e models.Equipment
	if err := ctx.ShouldBindJSON(&e); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.svc.Update(id, e); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *EquipmentController) delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.svc.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
