package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminHandler(c *gin.Context) {
	// user := c.MustGet(gin.AuthUserKey).(string)

	// Parse JSON
	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if c.Bind(&json) == nil {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
