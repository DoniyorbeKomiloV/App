package handler

import (
	"app/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateCategory godoc
// @ID create_category
// @Router /categories [POST]
// @Summary Create Category
// @Description Create Category
// @Tags Category
// @Accept json
// @Procedure json
// @Param user body models.CreateCategory true "CreateCategoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) CreateCategory(c *gin.Context) {
	var createCategory *models.CreateCategory
	err := c.ShouldBindJSON(&createCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Bad Request",
			"data":    err.Error(),
		})
		return
	}

	CategoryId, err := h.strg.Category().Create(c.Request.Context(), createCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return
	}
	Category, err := h.strg.Category().GetById(c.Request.Context(), &models.CategoryPrimaryKey{Id: CategoryId})
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
		"data":    Category,
	})
}

// UpdateCategory godoc
// @ID update_category
// @Router /categories [PUT]
// @Summary Update Category
// @Description Update Category
// @Tags Category
// @Accept json
// @Procedure json
// @Param user body models.UpdateCategory true "UpdateCategoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) UpdateCategory(c *gin.Context) {
	var category models.UpdateCategory
	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(400, map[string]interface{}{
			"status":  "error",
			"message": "Bad request",
			"data":    err.Error(),
		})
		return
	}
	resp, err := h.strg.Category().Update(c.Request.Context(), &category)
	if err != nil {
		c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Error while UpdateCategory",
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

// GetByIdCategory godoc
// @ID get_by_id_category
// @Router /categories/{id} [GET]
// @Summary Get By ID Category
// @Description Get By ID Category
// @Tags Category
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetByIdCategory(c *gin.Context) {
	var id = c.Param("id")
	category, err := h.strg.Category().GetById(c.Request.Context(), &models.CategoryPrimaryKey{Id: id})
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
		"message": "Category found",
		"data":    category,
	})
}

// GetListCategories godoc
// @ID get_list_category
// @Router /categories [GET]
// @Summary Get List Categories
// @Description Get List Categories
// @Tags Category
// @Accept json
// @Procedure jsonUser
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetListCategories(c *gin.Context) {
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	resp, err := h.strg.Category().GetList(c.Request.Context(), &models.CategoryGetListRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Error while GetListCategories",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "get list category response",
		"data":    resp,
	})
}

// DeleteCategory godoc
// @ID delete_category
// @Router /categories/{id} [DELETE]
// @Summary Delete Category
// @Description Delete Category
// @Tags Category
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) DeleteCategory(c *gin.Context) {
	var id = c.Param("id")

	err := h.strg.Category().Delete(c.Request.Context(), &models.CategoryPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Error while DeleteCategory",
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
