package handler

import (
	"app/api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// CreateUser godoc
// @ID create_user
// @Router /users [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Procedure json
// @Param user body models.CreateUser true "CreateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) CreateUser(c *gin.Context) {
	var createUser *models.CreateUser
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		h.handlerResponse(c, "JSON format is not valid", http.StatusBadRequest, err.Error())
		return
	}

	UserId, err := h.strg.Users().Create(c.Request.Context(), createUser)
	if err != nil {
		h.handlerResponse(c, "Error while creating User", http.StatusInternalServerError, err.Error())
		return
	}
	User, err := h.strg.Users().GetById(c.Request.Context(), &models.UserPrimaryKey{Id: UserId})
	if err != nil {
		h.handlerResponse(c, "Error while getting User", http.StatusInternalServerError, err.Error())
		return

	}

	h.handlerResponse(c, "User successfully created", http.StatusCreated, User)
}

// UpdateUser godoc
// @ID update_user
// @Router /users [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Procedure json
// @Param user body models.User true "UpdateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) UpdateUser(c *gin.Context) {
	var user models.UpdateUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handlerResponse(c, "JSON format is not valid", http.StatusBadRequest, err.Error())
		return
	}
	_, err = h.strg.Users().GetById(c.Request.Context(), &models.UserPrimaryKey{Id: user.Id})
	if err != nil {
		if err.Error() == fmt.Errorf("no rows in result set").Error() {
			h.handlerResponse(c, "User does not exist", http.StatusNotFound, nil)
			return
		}
		h.handlerResponse(c, "Error while getting User", http.StatusInternalServerError, err.Error())
		return
	}
	resp, err := h.strg.Users().Update(c.Request.Context(), &user)
	if err != nil {
		h.handlerResponse(c, "Error while updating User", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "User successfully updated", http.StatusCreated, resp)
}

// GetByIdUser godoc
// @ID get_by_id_user
// @Router /users/{id} [GET]
// @Summary Get By ID User
// @Description Get By ID User
// @Tags User
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetByIdUser(c *gin.Context) {
	var id = c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		h.handlerResponse(c, "Bad Request", http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.strg.Users().GetById(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		if err.Error() == fmt.Errorf("no rows in result set").Error() {
			h.handlerResponse(c, "User does not exist", http.StatusNotFound, err.Error())
			return
		}
		h.handlerResponse(c, "Error while getting User", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "User successfully retrieved", http.StatusOK, user)
}

// GetListUsers godoc
// @ID get_list_user
// @Router /users [GET]
// @Summary Get List Users
// @Description Get List Users
// @Tags User
// @Accept json
// @Procedure jsonUser
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) GetListUsers(c *gin.Context) {
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
	resp, err := h.strg.Users().GetList(c.Request.Context(), &models.UserGetListRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		h.handlerResponse(c, "Error while getting Users", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "User successfully retrieved", http.StatusOK, resp)
}

// DeleteUser godoc
// @ID delete_user
// @Router /users/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) DeleteUser(c *gin.Context) {
	var id = c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		h.handlerResponse(c, "Bad Request", http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.strg.Users().GetById(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		if err.Error() == fmt.Errorf("no rows in result set").Error() {
			h.handlerResponse(c, "User does not exist", http.StatusNotFound, nil)
			return
		}
		h.handlerResponse(c, "Error while getting User", http.StatusInternalServerError, err.Error())
		return
	}

	err = h.strg.Users().Delete(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "Error while deleting User", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "User deleted successfully", http.StatusOK, nil)
}
