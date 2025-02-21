package handler

import (
	"AyanTestProject/internal/model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) EventNotification(c *gin.Context) {
	var event model.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		slog.ErrorContext(c, "failed to read request body", "details", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	if err := h.srv.EventNotification(c, &event); err != nil {
		slog.ErrorContext(c, "failed to process event", "details", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
