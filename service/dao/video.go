package dao

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IVideoWriter interface {
	UpdateVideo(req UpdateVideoInfo) (err error)
	CreateVideo(Video *Video) (err error)
}

type IVideoReader interface {
	MGetVideoByIDs(IDs []int64) (videos []*Video, err error)
	GetVideosListByCreateTime(timestamp int64, count int) (videos []*Video, err error)
}

func (m *mysqlVideoWriter) GetVideosListByCreateTime(timestamp int64, count int) (videos []*Video, err error) {
	err = m.client.Table(Video{}.TableName()).
		Where("create_time < ?", timestamp).
		Order("create_time").
		Limit(count).
		Find(&videos).
		Error
	return
}

type mysqlVideoReader struct {
	client *gorm.DB
}

type mysqlVideoWriter struct {
	client *gorm.DB
}

type UpdateVideoInfo struct {
	ID            *int64
	AuthorID      *int64
	PlayURL       *string
	CoverURL      *string
	FavoriteCount *int
	CommentCount  *int
	Title         *string
	IsDelete      *bool
}

// gorm 默认ID主键，自动转换蛇形拼写
type Video struct {
	ID            int64
	AuthorID      int64
	PlayURL       string
	CoverURL      string
	FavoriteCount int
	CommentCount  int
	Title         string
	CreateTime    int64
	IsDelete      bool
}

func (Video) TableName() string {
	return "video"
}

func (m *mysqlVideoWriter) UpdateVideo(req UpdateVideoInfo) (err error) {
	if req.ID == nil || *req.ID == 0 {
		log.Error("Video id must exist and not be zero.")
		return nil
	}

	model := &Video{ID: *req.ID}
	fields := make(map[string]interface{}, 0)
	if req.AuthorID != nil {
		fields["author_id"] = *req.AuthorID
	}
	if req.PlayURL != nil {
		fields["play_url"] = *req.PlayURL
	}
	if req.CoverURL != nil {
		fields["cover_url"] = *req.CoverURL
	}
	if req.FavoriteCount != nil {
		fields["favorite_count"] = *req.FavoriteCount
	}
	if req.CommentCount != nil {
		fields["comment_count"] = *req.CommentCount
	}
	if req.Title != nil {
		fields["title"] = *req.Title
	}
	if req.IsDelete != nil {
		fields["is_delete"] = *req.IsDelete
	}
	fields["update_time"] = time.Now().Unix()

	err = m.client.Table(Video{}.TableName()).Model(model).
		Updates(fields).
		Error
	return
}

func (m *mysqlVideoWriter) CreateVideo(video *Video) (err error) {
	return m.client.Table(Video{}.TableName()).
		Create(video).
		Error
}

func (m *mysqlVideoWriter) MGetVideoByIDs(IDs []int64) (videos []*Video, err error) {
	err = m.client.Table(Video{}.TableName()).
		Where("id in (?)", IDs).
		Find(&videos).
		Error
	return
}
