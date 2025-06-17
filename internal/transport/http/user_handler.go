package http

import (
	"CADET_PROJECT/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, svc *service.UserService) {
	g := r.Group("/users")
	g.POST("", func(c *gin.Context) {
		var dto service.SignUpDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := svc.Create(dto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, user)
	})

	g.GET("", func(c *gin.Context) {
		list, _ := svc.List()
		c.JSON(http.StatusOK, list)
	})

	g.GET("/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		u, err := svc.Get(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, u)
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
