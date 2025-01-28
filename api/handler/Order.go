package handler

import (
	"app/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateOrder godoc
// @ID create_order
// @Router /orders [POST]
// @Summary Create Order
// @Description Create Order
// @Tags Order
// @Accept json
// @Procedure json
// @Param user body models.CreateOrder true "CreateOrderRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) CreateOrder(c *gin.Context) {
	var createOrder *models.CreateOrder
	err := c.ShouldBindJSON(&createOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Bad Request",
			"data":    err.Error(),
		})
		return
	}

	OrderId, err := h.strg.Order().Create(c.Request.Context(), createOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return
	}
	Order, err := h.strg.Order().GetById(c.Request.Context(), &models.OrderPrimaryKey{OrderId: OrderId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return

	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "User created",
		"data":    Order,
	})
}

// UpdateOrder godoc
// @ID update_order
// @Router /orders [PUT]
// @Summary Update Order
// @Description Update Order
// @Tags Order
// @Accept json
// @Procedure json
// @Param user body models.Order true "UpdateOrderRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) UpdateOrder(c *gin.Context) {
	var order models.UpdateOrder
	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(401, map[string]interface{}{
			"status":  "error",
			"message": "Bad request",
			"data":    err.Error(),
		})
		return
	}
	resp, err := h.strg.Order().Update(c.Request.Context(), &order)
	if err != nil {
		c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Error while UpdateOrder",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"status":  "OK",
		"message": "Success",
		"data":    resp,
	})
}

// GetByIdOrder godoc
// @ID get_by_id_order
// @Router /orders/{id} [GET]
// @Summary Get By ID Order
// @Description Get By ID Order
// @Tags Order
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetByIdOrder(c *gin.Context) {
	var id = c.Param("id")
	order, err := h.strg.Order().GetById(c.Request.Context(), &models.OrderPrimaryKey{OrderId: id})
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
		"message": "Order found",
		"data":    order,
	})
}

// GetListOrders godoc
// @ID get_list_order
// @Router /orders [GET]
// @Summary Get List Orders
// @Description Get List Orders
// @Tags Order
// @Accept json
// @Procedure jsonUser
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetListOrders(c *gin.Context) {
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	resp, err := h.strg.Order().GetList(c.Request.Context(), &models.OrderGetListRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Error while GetListOrders",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "get list order response",
		"data":    resp,
	})
}

// DeleteOrder godoc
// @ID delete_order
// @Router /orders/{id} [DELETE]
// @Summary Delete Order
// @Description Delete Order
// @Tags Order
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) DeleteOrder(c *gin.Context) {
	var id = c.Param("id")

	err := h.strg.Order().Delete(c.Request.Context(), &models.OrderPrimaryKey{OrderId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Error while DeleteOrder",
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
