package controlle

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go_shop/model"
	"go_shop/service/User"
	"go_shop/util"
)

type UserHandler struct {
	User *User.NewUser
}

func (u *UserHandler) Register(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return util.Resp401(c, fmt.Sprintf("参数存在问题 : %v", err))
	}
	token, err := u.User.UserCreate(user)
	if err != nil {
		return util.Resp400(c, fmt.Sprintf("注册存在错误: %v", err))
	}
	response := map[string]interface{}{
		"token": token,
	}
	return util.Resp200(c, response)
}
