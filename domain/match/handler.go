package match

import (
	"net/http"
	"strconv"

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
	courtNoStr := c.Params.ByName("court_no")
	dateTime := c.Params.ByName("date_time")

	courtNo, err := strconv.Atoi(courtNoStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"desc": "Invalid court_no format"})
		return
	}

	data, err := h.s.GetMatch(eventId, courtNo, dateTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"event_id": eventId, "court_no": courtNo, "date_time": dateTime, "desc": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"event_id": eventId, "court_no": courtNo, "date_time": dateTime, "value": data})
}

func (h Handler) Post(c *gin.Context) {
	var match Match
	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"desc": err.Error()})
		return
	}

	err := h.s.SaveMatch(match)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"desc": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"desc": "done"})
}