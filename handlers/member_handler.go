package handlers

import (
	"net/http"
	"strconv"
	"test-k-link-indonesia/dto"
	"test-k-link-indonesia/models"
	pkgbycript "test-k-link-indonesia/packages/bycript"
	"test-k-link-indonesia/repositories"
	"test-k-link-indonesia/utilities"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerMember struct {
	MemberRepository repositories.MemberRepository
}

func HandlerMember(MemberRepository repositories.MemberRepository) *handlerMember {
	return &handlerMember{MemberRepository}
}

func (h *handlerMember) CreateMember(c *gin.Context) {
	userLogin, exists := c.Get("userLogin")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "unauthorized",
		})
		return
	}

	var isAdmin = userLogin.(jwt.MapClaims)["admin"]

	if isAdmin == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "unauthorized",
		})
		return
	}

	request := new(dto.MemberRequestDTO)

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

	var gender models.Gender
	var level models.Level

	getGender, errGender := h.MemberRepository.GetGender(request.Gender)

	if errGender != nil {
		modelGender := models.Gender{
			Gender:    request.Gender,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		createGender, _ := h.MemberRepository.CreateGender(modelGender)

		gender = models.Gender{
			ID:        createGender.ID,
			Gender:    createGender.Gender,
			CreatedAt: createGender.CreatedAt,
			UpdatedAt: createGender.UpdatedAt,
		}
	} else {
		gender = models.Gender{
			ID:        getGender.ID,
			Gender:    getGender.Gender,
			CreatedAt: getGender.CreatedAt,
			UpdatedAt: getGender.UpdatedAt,
		}
	}

	getLevel, errLevel := h.MemberRepository.GetLevel(request.Level)
	if errLevel != nil {
		modelLevel := models.Level{
			Level:     request.Level,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		createLevel, _ := h.MemberRepository.CreateLevel(modelLevel)

		level = models.Level{
			ID:        createLevel.ID,
			Level:     createLevel.Level,
			CreatedAt: createLevel.CreatedAt,
			UpdatedAt: createLevel.UpdatedAt,
		}
	} else {
		level = models.Level{
			ID:        getLevel.ID,
			Level:     getLevel.Level,
			CreatedAt: getLevel.CreatedAt,
			UpdatedAt: getLevel.UpdatedAt,
		}
	}

	password, _ := pkgbycript.HashingPassword(request.NamaBelakang)

	user := models.User{
		Username:  request.NamaDepan,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, _ := h.MemberRepository.CreateUser(user)

	birthday, _ := utilities.ParseTime(request.Birthday)
	joinDate, _ := utilities.ParseTime(request.JoinDate)

	member := models.Member{
		NamaDepan:    request.NamaDepan,
		NamaBelakang: request.NamaBelakang,
		Birthday:     birthday,
		UserID:       createUser.ID,
		GenderID:     gender.ID,
		JoinDate:     joinDate,
		LevelID:      level.ID,
	}

	createMember, errMember := h.MemberRepository.CreateMember(member)
	if errMember != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": errMember.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   createMember,
	})
}

func (h *handlerMember) ShowAll(c *gin.Context) {
	userLogin, exists := c.Get("userLogin")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "unauthorized",
		})
		return
	}

	var isAdmin = userLogin.(jwt.MapClaims)["admin"]

	if isAdmin == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "unauthorized",
		})
		return
	}

	users, err := h.MemberRepository.GetAllMembers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   users,
	})
}

func (h *handlerMember) ShowMemberByAdmin(c *gin.Context) {
	member_id, _ := strconv.Atoi(c.Param("id"))

	userLogin, exists := c.Get("userLogin")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "unauthorized",
		})
		return
	}

	var isAdmin = userLogin.(jwt.MapClaims)["admin"]

	if isAdmin == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "unauthorized",
		})
		return
	}

	getMemberById, err := h.MemberRepository.GetMemberById(member_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   getMemberById,
	})

}

func (h *handlerMember) ShowDataMember(c *gin.Context) {
	userLogin, exists := c.Get("userLogin")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "unauthorized",
		})
		return
	}

	userId := int(userLogin.(jwt.MapClaims)["userId"].(float64))

	user, err := h.MemberRepository.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   user,
	})

}
