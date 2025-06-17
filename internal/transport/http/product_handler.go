package http

import (
	"CADET_PROJECT/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, svc *service.ProductService) {
	g := r.Group("/products")
	g.POST("", func(c *gin.Context) {
		var dto service.ProductDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		p, err := svc.Create(dto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, p)
	})

	g.GET("", func(c *gin.Context) {
		list, _ := svc.List()
		c.JSON(http.StatusOK, list)
	})

	g.GET("/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		p, err := svc.Get(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, p)
	})

	g.DELETE("/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := svc.Delete(uint(id)); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	})
}
