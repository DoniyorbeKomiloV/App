package handler

import (
	"app/api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
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
		h.handlerResponse(c, "JSON format is not valid", http.StatusBadRequest, err.Error())
		return
	}

	CategoryId, err := h.strg.Category().Create(c.Request.Context(), createCategory)
	if err != nil {
		h.handlerResponse(c, "Error while creating Category", http.StatusInternalServerError, err.Error())
		return
	}
	Category, err := h.strg.Category().GetById(c.Request.Context(), &models.CategoryPrimaryKey{Id: CategoryId})
	if err != nil {
		h.handlerResponse(c, "Error while getting Category", http.StatusInternalServerError, err.Error())
		return

	}

	h.handlerResponse(c, "Category successfully created", http.StatusCreated, Category)
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
		h.handlerResponse(c, "JSON format is not valid", http.StatusBadRequest, err.Error())
		return
	}
	_, err = h.strg.Category().GetById(c.Request.Context(), &models.CategoryPrimaryKey{Id: category.Id})
	if err != nil {
		if err.Error() == fmt.Errorf("no rows in result set").Error() {
			h.handlerResponse(c, "Category does not exist", http.StatusNotFound, nil)
			return
		}
		h.handlerResponse(c, "Error while getting Category", http.StatusInternalServerError, err.Error())
		return
	}
	resp, err := h.strg.Category().Update(c.Request.Context(), &category)
	if err != nil {
		h.handlerResponse(c, "Error while updating Category", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "Category successfully updated", http.StatusCreated, resp)
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
	if _, err := uuid.Parse(id); err != nil {
		h.handlerResponse(c, "Bad Request", http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.strg.Category().GetById(c.Request.Context(), &models.CategoryPrimaryKey{Id: id})
	if err != nil {
		if err.Error() == fmt.Errorf("no rows in result set").Error() {
			h.handlerResponse(c, "Category does not exist", http.StatusNotFound, err.Error())
			return
		}
		h.handlerResponse(c, "Error while getting Category", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "Category successfully retrieved", http.StatusOK, category)
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
	resp, err := h.strg.Category().GetList(c.Request.Context(), &models.CategoryGetListRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		h.handlerResponse(c, "Error while getting Categories", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "Category successfully retrieved", http.StatusOK, resp)
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
	if _, err := uuid.Parse(id); err != nil {
		h.handlerResponse(c, "Bad Request", http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.strg.Category().GetById(c.Request.Context(), &models.CategoryPrimaryKey{Id: id})
	if err != nil {
		if err.Error() == fmt.Errorf("no rows in result set").Error() {
			h.handlerResponse(c, "Category does not exist", http.StatusNotFound, nil)
			return
		}
		h.handlerResponse(c, "Error while getting Category", http.StatusInternalServerError, err.Error())
		return
	}

	err = h.strg.Category().Delete(c.Request.Context(), &models.CategoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "Error while deleting Category", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Category deleted successfully", http.StatusOK, nil)
}
