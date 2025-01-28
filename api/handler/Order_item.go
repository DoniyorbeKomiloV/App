package handler

import (
	"app/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateOrderItem godoc
// @ID create_OrderItem
// @Router /order_items [POST]
// @Summary Create OrderItem
// @Description Create OrderItem
// @Tags OrderItem
// @Accept json
// @Procedure json
// @Param user body models.CreateOrderItem true "CreateOrderItemRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) CreateOrderItem(c *gin.Context) {
	var createOrderItem *models.CreateOrderItem
	err := c.ShouldBindJSON(&createOrderItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Bad Request",
			"data":    err.Error(),
		})
		return
	}

	OrderItemId, err := h.strg.OrderItem().Create(c.Request.Context(), createOrderItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return
	}
	OrderItem, err := h.strg.OrderItem().GetById(c.Request.Context(), &models.OrderItemPrimaryKey{ItemId: OrderItemId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "OK",
		"message": "User created",
		"data":    OrderItem,
	})
}

// UpdateOrderItem godoc
// @ID update_OrderItem
// @Router /order_items [PUT]
// @Summary Update OrderItem
// @Description Update OrderItem
// @Tags OrderItem
// @Accept json
// @Procedure json
// @Param user body models.OrderItem true "UpdateOrderItemRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) UpdateOrderItem(c *gin.Context) {
	var OrderItem models.UpdateOrderItem
	err := c.ShouldBindJSON(&OrderItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Bad request",
			"data":    err.Error(),
		})
		return
	}
	resp, err := h.strg.OrderItem().Update(c.Request.Context(), &OrderItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Error while UpdateOrderItem",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "OK",
		"message": "Success",
		"data":    resp,
	})
}

// GetByIdOrderItem godoc
// @ID get_by_id_OrderItem
// @Router /order_items/{id} [GET]
// @Summary Get By ID OrderItem
// @Description Get By ID OrderItem
// @Tags OrderItem
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetByIdOrderItem(c *gin.Context) {
	var id = c.Param("id")
	OrderItem, err := h.strg.OrderItem().GetById(c.Request.Context(), &models.OrderItemPrimaryKey{ItemId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Server internal Error",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "OrderItem found",
		"data":    OrderItem,
	})
}

// GetListOrderItems godoc
// @ID get_list_OrderItem
// @Router /order_items [GET]
// @Summary Get List OrderItems
// @Description Get List OrderItems
// @Tags OrderItem
// @Accept json
// @Procedure jsonUser
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetListOrderItems(c *gin.Context) {
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	resp, err := h.strg.OrderItem().GetList(c.Request.Context(), &models.OrderItemGetListRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Error while GetListOrderItems",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "get list OrderItem response",
		"data":    resp,
	})
}

// DeleteOrderItem godoc
// @ID delete_OrderItem
// @Router /order_items/{id} [DELETE]
// @Summary Delete OrderItem
// @Description Delete OrderItem
// @Tags OrderItem
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) DeleteOrderItem(c *gin.Context) {
	var id = c.Param("id")

	err := h.strg.OrderItem().Delete(c.Request.Context(), &models.OrderItemPrimaryKey{ItemId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Error while DeleteOrderItem",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusNoContent, map[string]interface{}{
		"status":  "OK",
		"message": "Success",
		"data":    nil,
	})
}
