package event

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
	eventId := c.Params.ByName("event_id")

	data, err := h.s.GetEvent(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"event_id": eventId, "desc": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"event_id": eventId, "value": data})
}

func (h Handler) Post(c *gin.Context) {
	var event Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"desc": err.Error()})
		return
	}

	err := h.s.SaveEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"desc": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"desc": "done"})
}