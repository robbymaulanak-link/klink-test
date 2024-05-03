package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"test-k-link-indonesia/dto"
	"test-k-link-indonesia/models"
	pkgbycript "test-k-link-indonesia/packages/bycript"
	pkgjwt "test-k-link-indonesia/packages/jwt"
	"test-k-link-indonesia/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) Register(c *gin.Context) {
	request := new(dto.RegisterRequestDTO)

	fmt.Println(request.Username, request.Password)

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	validation := validator.New()

	err := validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	var user models.User

	password, _ := pkgbycript.HashingPassword(request.Password)

	if request.Username == "admin" {
		user = models.User{
			Username: request.Username,
			Password: password,
			IsAdmin:  true,
		}
	} else {
		user = models.User{
			Username: request.Username,
			Password: password,
			IsAdmin:  false,
		}
	}

	userRegister, err := h.UserRepository.Register(user)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Register Success",
		"data":    userRegister,
	})
}

func (h *handlerUser) Login(c *gin.Context) {
	request := new(dto.LoginRequestDTO)

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	user, err := h.UserRepository.Login(request.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "username not registered",
		})
		return
	}

	validPassword := pkgbycript.CheckPasswordHash(request.Password, user.Password)
	if !validPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "wrong password",
		})
		return
	}

	claims := jwt.MapClaims{}
	claims["userId"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["admin"] = user.IsAdmin

	token, _ := pkgjwt.GenerateToken(&claims)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   user,
		"token":  token,
	})
}
