package player

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
	playerName := c.Params.ByName("player_name")

	data, err := h.s.GetPlayer(eventId, playerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"event_id": eventId, "player_name": playerName, "desc": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"event_id": eventId, "player_name": playerName, "value": data})
}

func (h Handler) Post(c *gin.Context) {
	var player Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"desc": err.Error()})
		return
	}

	err := h.s.SavePlayer(player)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"desc": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"desc": "done"})
}