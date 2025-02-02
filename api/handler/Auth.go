package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
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
		h.handlerResponse(c, "JSON format is not valid", http.StatusBadRequest, err.Error())
		return
	}
	if !helper.IsValidLogin(login.Username) {
		h.handlerResponse(c, "A Username must start with a letter, the length of a Username should be from "+
			"6 to 30 (inclusive) and may contain uppercase and lowercase letters, underscores and digits",
			http.StatusBadRequest, "Username is not valid")
		return
	}
	resp, err := h.strg.Users().GetById(context.Background(), &models.UserPrimaryKey{Username: login.Username})
	if err != nil {
		if err.Error() == fmt.Errorf("no rows in result set").Error() {
			h.handlerResponse(c, "User does not exist", http.StatusBadRequest, err.Error())
			return
		}
		h.handlerResponse(c, "Error while getting User", http.StatusInternalServerError, err.Error())
		return
	}

	if !helper.CheckPasswordHash(login.Password, resp.Password) {
		h.handlerResponse(c, "Password mismatch", http.StatusBadRequest, "Incorrect password")
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
	if !helper.IsValidLogin(createUser.Username) {
		h.handlerResponse(c, "A Username must start with a letter, the length of a Username should be from "+
			"6 to 30 (inclusive) and may contain uppercase and lowercase letters, underscores and digits",
			http.StatusBadRequest, "Username is not valid")
		return
	}

	// TO-DO
	// implement better password checker !!!

	if len(createUser.Password) < 7 {
		h.handlerResponse(c, "Password should include more than 7 elements",
			http.StatusBadRequest, errors.New("Password len should include more than 7 elements"))
		return
	}

	hashedPassword, err := helper.HashPassword(createUser.Password)
	if err != nil {
		h.handlerResponse(c, "Error hashing password", http.StatusInternalServerError, err.Error())
		return
	}

	createUser.Password = hashedPassword

	resp, err := h.strg.Users().GetById(context.Background(), &models.UserPrimaryKey{Username: createUser.Username})
	if err != nil {
		if err.Error() == fmt.Errorf("no rows in result set").Error() {
			id, err = h.strg.Users().Create(context.Background(), &createUser)
			if err != nil {
				h.handlerResponse(c, "Error while creating User", http.StatusInternalServerError, err.Error())
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

	h.handlerResponse(c, "User successfully created", http.StatusCreated, resp)
}

func (h *Handler) Validate(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		h.handlerResponse(c, "Authorization header not present", http.StatusUnauthorized, err.Error())
		return
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
		h.handlerResponse(c, "Invalid", http.StatusUnauthorized, err.Error())
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) >= claims["exp"].(float64) {
			h.handlerResponse(c, "Token expired", http.StatusUnauthorized, "Token expired")
			return
		}
		var userId = cast.ToString(claims["user_id"])
		user, err := h.strg.Users().GetById(c, &models.UserPrimaryKey{Id: userId})
		if err != nil {
			h.handlerResponse(c, "User does not exist", http.StatusNotFound, "User does not exist")
			return
		}
		c.Set("user_id", userId)
		c.Set("user", user)
		c.Next()
	} else {
		h.handlerResponse(c, "Authorization header not present", http.StatusUnauthorized, err.Error())
		return
	}
}
