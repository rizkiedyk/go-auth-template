package handler

import (
	"go-auth/domain/dto"
	"go-auth/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) SetRole(c *gin.Context) {
	var req dto.SetRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, dto.Resp{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	user, err := u.userService.SetRole(req)
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
		Message: "user role updated",
		Data:    user,
	})
}