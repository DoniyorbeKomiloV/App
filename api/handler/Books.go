package handler

import (
	"app/api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
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
		h.handlerResponse(c, "JSON format is not valid", http.StatusBadRequest, err.Error())
		return
	}

	BookId, err := h.strg.Books().Create(c.Request.Context(), createBook)
	if err != nil {
		h.handlerResponse(c, "Error while creating Book", http.StatusInternalServerError, err.Error())
		return
	}
	Book, err := h.strg.Books().GetById(c.Request.Context(), &models.BookPrimaryKey{Id: BookId})
	if err != nil {
		h.handlerResponse(c, "Error while getting Book", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Book successfully created", http.StatusCreated, Book)
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
		h.handlerResponse(c, "JSON format is not valid", http.StatusBadRequest, err.Error())
		return
	}
	_, err = h.strg.Books().GetById(c.Request.Context(), &models.BookPrimaryKey{Id: book.Id})
	if err != nil {
		if err.Error() == fmt.Errorf("no rows in result set").Error() {
			h.handlerResponse(c, "Book does not exist", http.StatusNotFound, nil)
			return
		}
		h.handlerResponse(c, "Error while getting Book", http.StatusInternalServerError, err.Error())
		return
	}
	resp, err := h.strg.Books().Update(c.Request.Context(), &book)
	if err != nil {
		h.handlerResponse(c, "Error while updating Book", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "Book successfully updated", http.StatusCreated, resp)
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
	if _, err := uuid.Parse(id); err != nil {
		h.handlerResponse(c, "Bad Request", http.StatusBadRequest, err.Error())
		return
	}

	book, err := h.strg.Books().GetById(c.Request.Context(), &models.BookPrimaryKey{Id: id})
	if err != nil {
		if err.Error() == fmt.Errorf("no rows in result set").Error() {
			h.handlerResponse(c, "Book does not exist", http.StatusNotFound, err.Error())
			return
		}
		h.handlerResponse(c, "Error while getting Book", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "Book successfully retrieved", http.StatusOK, book)
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
	resp, err := h.strg.Books().GetList(c.Request.Context(), &models.BookGetListRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		h.handlerResponse(c, "Error while getting Books", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "Book successfully retrieved", http.StatusOK, resp)
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
	if _, err := uuid.Parse(id); err != nil {
		h.handlerResponse(c, "Bad Request", http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.strg.Books().GetById(c.Request.Context(), &models.BookPrimaryKey{Id: id})
	if err != nil {
		if err.Error() == fmt.Errorf("no rows in result set").Error() {
			h.handlerResponse(c, "Book does not exist", http.StatusNotFound, nil)
			return
		}
		h.handlerResponse(c, "Error while getting Book", http.StatusInternalServerError, err.Error())
		return
	}

	err = h.strg.Books().Delete(c.Request.Context(), &models.BookPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "Error while deleting Book", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Book deleted successfully", http.StatusOK, nil)
}
