package dao

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IUserWriter interface {
	UpdateUser(user *UpdateUserInfo) (err error)
	CreateUser(user *User) (err error)
}

type IUserReader interface {
	MGetUserByIDs(IDs []int64) (users []*User, err error)
}

type mysqlUserReader struct {
	client *gorm.DB
}

type mysqlUserWriter struct {
	client *gorm.DB
}

type UpdateUserInfo struct {
	ID       *int64
	UserName *string
	Password *string
	IsDelete *bool
}

// gorm 默认ID主键，自动转换蛇形拼写
type User struct {
	ID         int64
	UserName   string
	Password   string
	CreateTime int64
	UpdateTIme int64
	IsDelete   bool
}

func (User) TableName() string {
	return "user"
}

func (m *mysqlUserWriter) UpdateUser(req *UpdateUserInfo) (err error) {
	if req.ID == nil || *req.ID == 0 {
		log.Error("user id must exist and not be zero.")
		return nil
	}

	model := &User{ID: *req.ID}
	fields := make(map[string]interface{}, 0)
	if req.UserName != nil {
		fields["user_name"] = *req.UserName
	}
	if req.Password != nil {
		fields["password"] = *req.UserName
	}
	if req.IsDelete != nil {
		fields["is_delete"] = *req.IsDelete
	}
	fields["update_time"] = time.Now().Unix()

	err = m.client.Table(User{}.TableName()).Model(model).
		Updates(fields).
		Error
	return
}

func (m *mysqlUserWriter) CreateUser(user *User) (err error) {
	return m.client.Table(user.TableName()).
		Create(user).
		Error
}

func (m *mysqlUserReader) MGetUserByIDs(IDs []int64) (users []*User, err error) {
	err = m.client.Table(User{}.TableName()).
		Where("id in (?)", IDs).
		Find(&users).
		Error
	return
}
