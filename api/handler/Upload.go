package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) HandleUpload(c *gin.Context) {
	file, err := c.FormFile("image")

	if err != nil {
		h.handlerResponse(c, "Handle upload", http.StatusBadRequest, err.Error())
		return
	}

	err = c.SaveUploadedFile(file, "./uploads/"+file.Filename)
	if err != nil {
		h.handlerResponse(c, "Save upload", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Successfully uploaded", http.StatusCreated, "uploads/"+file.Filename)
}
