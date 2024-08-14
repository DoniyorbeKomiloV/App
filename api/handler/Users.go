package handler

import (
	"app/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "Error",
			"message": "Bad Request",
			"data":    err.Error(),
		})
		return
	}

	UserId, err := h.strg.Users().Create(c.Request.Context(), createUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Server internal",
			"data":    err.Error(),
		})
		return
	}
	User, err := h.strg.Users().GetById(c.Request.Context(), &models.UserPrimaryKey{Id: UserId})
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
		"data":    User,
	})
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
		c.JSON(401, map[string]interface{}{
			"status":  "error",
			"message": "Bad request",
			"data":    err.Error(),
		})
		return
	}
	resp, err := h.strg.Users().Update(c.Request.Context(), &user)
	if err != nil {
		c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Error while UpdateUser",
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
	user, err := h.strg.Users().GetById(c.Request.Context(), &models.UserPrimaryKey{Id: id})
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
		"message": "User found",
		"data":    user,
	})
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
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	resp, err := h.strg.Users().GetList(c.Request.Context(), &models.UserGetListRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Error while GetListUsers",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "get list user response",
		"data":    resp,
	})
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

	err := h.strg.Users().Delete(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Error while DeleteUser",
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
