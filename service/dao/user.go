package dao

import "gorm.io/gorm"

type IUserWriter interface {
}

type IUserReader interface {
}

type mysqlUserReader struct {
	client *gorm.DB
}

type mysqlUserWriter struct {
	client *gorm.DB
}
