package health

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	repo VersionGetterRepository
}

func (h RequestHandler) check(context *gin.Context) {
	version, err := h.repo.GetVersion()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("repo.GetVersion: %w", err),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message":      "System is healthy",
		"mysqlVersion": version,
	})
}
