package handler

import (
	"app/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateBook godoc
// @ID create_book
// @Router /books [POST]
// @Summary Create Book
// @Description Create Book
// @Tags Book
// @Accept json
// @Procedure json
// @Param user body models.CreateBook true "CreateBookRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) CreateBook(c *gin.Context) {
	var createBook *models.CreateBook
	err := c.ShouldBindJSON(&createBook)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "Error",
			"message": "Bad Request",
			"data":    err.Error(),
		})
		return
	}

	BookId, err := h.strg.Books().Create(c.Request.Context(), createBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return
	}
	Book, err := h.strg.Books().GetById(c.Request.Context(), &models.BookPrimaryKey{Id: BookId})
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
		"data":    Book,
	})
}

// UpdateBook godoc
// @ID update_book
// @Router /books [PUT]
// @Summary Update Book
// @Description Update Book
// @Tags Book
// @Accept json
// @Procedure json
// @Param user body models.Book true "UpdateBookRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) UpdateBook(c *gin.Context) {
	var book models.UpdateBook
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(401, map[string]interface{}{
			"status":  "error",
			"message": "Bad request",
			"data":    err.Error(),
		})
		return
	}
	resp, err := h.strg.Books().Update(c.Request.Context(), &book)
	if err != nil {
		c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Error while UpdateBook",
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

// GetByIdBook godoc
// @ID get_by_id_book
// @Router /books/{id} [GET]
// @Summary Get By ID Book
// @Description Get By ID Book
// @Tags Book
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetByIdBook(c *gin.Context) {
	var id = c.Param("id")
	book, err := h.strg.Books().GetById(c.Request.Context(), &models.BookPrimaryKey{Id: id})
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
		"message": "Book found",
		"data":    book,
	})
}

// GetListBooks godoc
// @ID get_list_book
// @Router /books [GET]
// @Summary Get List Books
// @Description Get List Books
// @Tags Book
// @Accept json
// @Procedure jsonUser
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetListBooks(c *gin.Context) {
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	resp, err := h.strg.Books().GetList(c.Request.Context(), &models.BookGetListRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Error while GetListBooks",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "get list book response",
		"data":    resp,
	})
}

// DeleteBook godoc
// @ID delete_book
// @Router /books/{id} [DELETE]
// @Summary Delete Book
// @Description Delete Book
// @Tags Book
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) DeleteBook(c *gin.Context) {
	var id = c.Param("id")

	err := h.strg.Books().Delete(c.Request.Context(), &models.BookPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Error while DeleteBook",
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
