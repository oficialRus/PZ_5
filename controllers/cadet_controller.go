// controllers/cadet_controller.go
package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"cadet_project/models"
	"cadet_project/services"
)

type CadetController struct{ svc *services.CadetService }

func NewCadetController(r *gin.Engine, svc *services.CadetService) {
	c := &CadetController{svc}
	g := r.Group("/cadets")
	g.GET("", c.list)
	g.POST("", c.add)
	g.PUT("/:id", c.update)
	g.DELETE("/:id", c.delete)
}

func (c *CadetController) list(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.svc.List())
}

func (c *CadetController) add(ctx *gin.Context) {
	var cadet models.Cadet
	if err := ctx.ShouldBindJSON(&cadet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, c.svc.Add(cadet))
}

func (c *CadetController) update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var cadet models.Cadet
	if err := ctx.ShouldBindJSON(&cadet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.svc.Update(id, cadet); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *CadetController) delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.svc.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
