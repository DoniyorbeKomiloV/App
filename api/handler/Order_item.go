package handler

import (
	"app/api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
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
		h.handlerResponse(c, "JSON format is not valid", http.StatusBadRequest, err.Error())
		return
	}

	OrderItemId, err := h.strg.OrderItem().Create(c.Request.Context(), createOrderItem)
	if err != nil {
		h.handlerResponse(c, "Error while creating OrderItem", http.StatusInternalServerError, err.Error())
		return
	}
	OrderItem, err := h.strg.OrderItem().GetById(c.Request.Context(), &models.OrderItemPrimaryKey{ItemId: OrderItemId})
	if err != nil {
		h.handlerResponse(c, "Error while getting OrderItem", http.StatusInternalServerError, err.Error())
		return

	}

	h.handlerResponse(c, "OrderItem successfully created", http.StatusCreated, OrderItem)
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
	var orderItem models.UpdateOrderItem
	err := c.ShouldBindJSON(&orderItem)
	if err != nil {
		h.handlerResponse(c, "JSON format is not valid", http.StatusBadRequest, err.Error())
		return
	}
	_, err = h.strg.OrderItem().GetById(c.Request.Context(), &models.OrderItemPrimaryKey{ItemId: orderItem.ItemId})
	if err != nil {
		if err.Error() == fmt.Errorf("no rows in result set").Error() {
			h.handlerResponse(c, "OrderItem does not exist", http.StatusNotFound, nil)
			return
		}
		h.handlerResponse(c, "Error while getting OrderItem", http.StatusInternalServerError, err.Error())
		return
	}
	resp, err := h.strg.OrderItem().Update(c.Request.Context(), &orderItem)
	if err != nil {
		h.handlerResponse(c, "Error while updating OrderItem", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "OrderItem successfully updated", http.StatusCreated, resp)
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
	if _, err := uuid.Parse(id); err != nil {
		h.handlerResponse(c, "Bad Request", http.StatusBadRequest, err.Error())
		return
	}

	orderItem, err := h.strg.OrderItem().GetById(c.Request.Context(), &models.OrderItemPrimaryKey{ItemId: id})
	if err != nil {
		if err.Error() == fmt.Errorf("no rows in result set").Error() {
			h.handlerResponse(c, "OrderItem does not exist", http.StatusNotFound, err.Error())
			return
		}
		h.handlerResponse(c, "Error while getting OrderItem", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "OrderItem successfully retrieved", http.StatusOK, orderItem)
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
	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "Error while parsing offset", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "Error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.strg.OrderItem().GetList(c.Request.Context(), &models.OrderItemGetListRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		h.handlerResponse(c, "Error while getting Categories", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "OrderItem successfully retrieved", http.StatusOK, resp)
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
	if _, err := uuid.Parse(id); err != nil {
		h.handlerResponse(c, "Bad Request", http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.strg.OrderItem().GetById(c.Request.Context(), &models.OrderItemPrimaryKey{ItemId: id})
	if err != nil {
		if err.Error() == fmt.Errorf("no rows in result set").Error() {
			h.handlerResponse(c, "OrderItem does not exist", http.StatusNotFound, nil)
			return
		}
		h.handlerResponse(c, "Error while getting OrderItem", http.StatusInternalServerError, err.Error())
		return
	}

	err = h.strg.OrderItem().Delete(c.Request.Context(), &models.OrderItemPrimaryKey{ItemId: id})
	if err != nil {
		h.handlerResponse(c, "Error while deleting OrderItem", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "OrderItem deleted successfully", http.StatusOK, nil)
}
