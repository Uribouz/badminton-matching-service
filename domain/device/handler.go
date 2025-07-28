package device

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	s Service
}

func NewHandler(service Service) Handler {
	return Handler{service}
}

func (h Handler) Get(c *gin.Context) {
	id := c.Params.ByName("id")

	data, err := h.s.GetDevice(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"id": id, "desc": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"id": id, "value": data})
}

func (h Handler) Post(c *gin.Context) {
	var device Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"desc": err.Error()})
		return
	}

	err := h.s.SaveDevice(device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"desc": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"desc": "done"})
}
