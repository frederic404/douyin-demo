package domain

import (
	"time"

	"github.com/bytedance/douyin-demo/api/protoc_gen"
	"github.com/bytedance/douyin-demo/pkg/constant"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func Feed(c *gin.Context, req *protoc_gen.DouyinFeedRequest) (resp *protoc_gen.DouyinFeedResponse,
	err error) {

	videoList := make([]*protoc_gen.Video, 0)
	videoList = append(videoList, constant.DemoVideo)

	resp = &protoc_gen.DouyinFeedResponse{
		StatusCode: proto.Int32(constant.SUCCESS),
		StatusMsg:  proto.String("success"),
		VideoList:  videoList,
		NextTime:   proto.Int64(time.Unix(100, 100).Unix()),
	}

	return resp, nil

}
