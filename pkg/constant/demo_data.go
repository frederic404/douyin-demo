package constant

import (
	"github.com/bytedance/douyin-demo/api/protoc_gen"
	"google.golang.org/protobuf/proto"
)

var DemoUser = &protoc_gen.User{
	Id:            proto.Int64(1),
	Name:          proto.String("DemoUser"),
	FollowCount:   proto.Int64(1),
	FollowerCount: proto.Int64(1),
	IsFollow:      proto.Bool(true),
}

var DemoVideo = &protoc_gen.Video{
	Id:            proto.Int64(1),
	Author:        DemoUser,
	PlayUrl:       proto.String("http://10.93.233.112:8080/static/bear.mp4"),
	CoverUrl:      proto.String("http://10.93.233.112:8080/static/bear-1283347_1280.jpeg"),
	FavoriteCount: proto.Int64(0),
	CommentCount:  proto.Int64(0),
	IsFavorite:    proto.Bool(false),
	Title:         proto.String("DemoVideo"),
}
