package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Procedure json
// @Param login body models.UserLoginRequest true "UserLoginRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) Login(c *gin.Context) {
	var login models.UserLoginRequest

	err := c.ShouldBindJSON(&login) // parse req body to given type struct
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Bad Request",
			"data":    err.Error(),
		})
		return
	}

	resp, err := h.strg.Users().GetById(context.Background(), &models.UserPrimaryKey{Username: login.Username})
	if err != nil {
		if err.Error() == "no rows in result set" {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status":  "Error",
				"message": "User does not exist",
				"data":    err.Error(),
			})
			return
		}
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(resp)

	if resp.Password != login.Password {
		h.handlerResponse(c, "Wrong password", http.StatusBadRequest, "Wrong password")
		return
	}

	token, err := helper.GenerateJWT(map[string]interface{}{
		"user_id": resp.Id,
	}, time.Hour*72, h.cfg.SecretKey)
	c.SetCookie("Authorization", token, 3600*72, "/", "localhost", false, true)
	h.handlerResponse(c, "token", http.StatusCreated, token)
}

// Register godoc
// @ID register
// @Router /register [POST]
// @Summary Register
// @Description Register
// @Tags Auth
// @Accept json
// @Procedure json
// @Param register body models.CreateUser true "CreateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *Handler) Register(c *gin.Context) {

	var createUser models.CreateUser
	var id string
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		h.handlerResponse(c, "error user should bind json", http.StatusBadRequest, err.Error())
		return
	}

	if len(createUser.Password) < 7 {
		h.handlerResponse(c, "Password should include more than 7 elements",
			http.StatusBadRequest, errors.New("Password len should include more than 7 elements"))
		return
	}

	resp, err := h.strg.Users().GetById(context.Background(), &models.UserPrimaryKey{Username: createUser.Username})
	if err != nil {
		if err.Error() == "no rows in result set" {
			id, err = h.strg.Users().Create(context.Background(), &createUser)
			if err != nil {
				h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			h.handlerResponse(c, "User already exists", http.StatusConflict, err.Error())
			return
		}
	} else if err == nil {
		h.handlerResponse(c, "User already exist", http.StatusConflict, nil)
		return
	}
	resp, err = h.strg.Users().GetById(context.Background(), &models.UserPrimaryKey{Id: id})

	h.handlerResponse(c, "create user response", http.StatusCreated, resp)
}

func (h *Handler) Validate(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		h.handlerResponse(c, "Authorization header not present", http.StatusUnauthorized, err.Error())
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(h.cfg.SecretKey), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) >= claims["exp"].(float64) {
			h.handlerResponse(c, "Token expired", http.StatusUnauthorized, "Token expired")
		}
		var userId = cast.ToString(claims["user_id"])
		user, err := h.strg.Users().GetById(c, &models.UserPrimaryKey{Id: userId})
		if err != nil {
			h.handlerResponse(c, "User does not exist", http.StatusNotFound, "User does not exist")
		}
		c.Set("user_id", userId)
		c.Set("user", user)
		c.Next()
	} else {
		h.handlerResponse(c, "Authorization header not present", http.StatusUnauthorized, err.Error())
	}

}
