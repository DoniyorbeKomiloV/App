package handler

import (
	"app/config"
	"app/pkg/logger"
	"app/storage"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"strconv"
)

type Handler struct {
	cfg    *config.Config
	logger logger.LoggerI
	strg   storage.StorageInterface
}

type Response struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func NewHandler(cfg *config.Config, storage storage.StorageInterface, logger logger.LoggerI) *Handler {
	return &Handler{
		cfg:    cfg,
		logger: logger,
		strg:   storage,
	}
}

func (h *Handler) getOffsetQuery(offset string) (int, error) {

	if len(offset) <= 0 {
		return h.cfg.DefaultOffset, nil
	}

	return strconv.Atoi(offset)
}

func (h *Handler) getLimitQuery(limit string) (int, error) {

	if len(limit) <= 0 {
		return h.cfg.DefaultLimit, nil
	}

	return strconv.Atoi(limit)
}

func (h *Handler) handlerResponse(c *gin.Context, path string, code int, message interface{}) {
	response := Response{
		Status:      code,
		Data:        message,
		Description: path,
	}

	switch {
	case code < 300:
		h.logger.Info(path, logger.Any("info", response))
	case code >= 400:
		h.logger.Error(path, logger.Any("error", response))
	}

	c.JSON(code, response)
}

func (h *Handler) SendMessageToMail(subject string, sEmail string, sPassword string, To string, Message string) error {
	// Set up the email message
	messageBody := Message

	// Create a new email message using gomail
	m := gomail.NewMessage()
	m.SetHeader("From", sEmail)
	m.SetHeader("To", To)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", messageBody)

	// Set up the email sending configuration
	d := gomail.NewDialer("smtp.gmail.com", 587, sEmail, sPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
