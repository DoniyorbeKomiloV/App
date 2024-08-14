package handler

import (
	"app/config"
	"app/pkg/logger"
	"app/storage"
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
