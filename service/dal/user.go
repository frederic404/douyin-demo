package dal

import (
	"time"

	"github.com/bytedance/douyin-demo/service/dao"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	UserName string
	Password string
}

func CreateUser(c gin.Context, req *CreateUserRequest) {
	err := dao.UserWriter.CreateUser(&dao.User{
		UserName:   req.UserName,
		Password:   req.Password,
		CreateTime: time.Now().Unix(),
		UpdateTIme: time.Now().Unix(),
		IsDelete:   false,
	})
	if err != nil {
		return
	}

}
