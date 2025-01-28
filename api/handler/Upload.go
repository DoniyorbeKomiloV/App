package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) HandleUpload(c *gin.Context) {
	file, err := c.FormFile("image")

	if err != nil {
		h.handlerResponse(c, "Handle upload", http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Bad Request",
			"data":    err.Error(),
		})
		return
	}

	err = c.SaveUploadedFile(file, "./uploads/"+file.Filename)
	if err != nil {
		h.handlerResponse(c, "Save upload", http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Error while saving the image",
			"data":    err.Error(),
		})
		return
	}

	h.handlerResponse(c, "Handle upload", http.StatusCreated, map[string]interface{}{
		"status":  "OK",
		"message": "Upload successfully",
		"data":    "uploads/" + file.Filename,
	})
}
