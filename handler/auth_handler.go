package handler

import (
	"go-auth/domain/dto"
	"go-auth/domain/model"
	"go-auth/service"
	"go-auth/utils/validator"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authService service.IAuthService
}


func NewAuthHandler(authService service.IAuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var reqUser dto.ReqRegister
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.JSON(400, dto.Resp{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	validationErrors := validator.ValidateStruct(reqUser)
    if validationErrors != nil {
		c.JSON(400, dto.Resp{
			Code:    400,
			Message: "invalid request body",
			Data:    validationErrors,
		})
        return
    }

	user := model.User{
		Id: uuid.New().String(),
		Username: reqUser.Username,
		Password: reqUser.Password,
		Email:    reqUser.Email,
		CreatedAt: int(time.Now().Unix()),
		UpdatedAt: int(time.Now().Unix()),
	}

	if err := h.authService.Register(user); err != nil {
		c.JSON(400, dto.Resp{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(200, dto.Resp{
		Code:    200,
		Message: "register success",
		Data:    nil,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.ReqLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, dto.Resp{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	validationErrors := validator.ValidateStruct(req)
    if validationErrors != nil {
		c.JSON(400, dto.Resp{
			Code:    400,
			Message: "invalid request body",
			Data:    validationErrors,
		})
        return
    }

	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(400, dto.Resp{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(200, dto.Resp{
		Code:    200,
		Message: "login success",
		Data:    token,
	})
}