package action

import (
	"net/http"
	"strconv"
	"time"

	"github.com/bytedance/douyin-demo/api/protoc_gen"
	"github.com/bytedance/douyin-demo/pkg/constant"
	"github.com/bytedance/douyin-demo/service/domain/feed"
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	latestTime, intError := getLatestTime(c)
	if intError != nil {
		return
	}

	req := &protoc_gen.DouyinFeedRequest{
		LatestTime: &latestTime,
		Token:      nil,
	}
	resp, err := feed.Feed(c, req)
	if err != nil || (resp != nil && resp.StatusCode != nil && *resp.StatusCode != constant.SUCCESS) {
		c.JSON(http.StatusBadRequest, resp)
		return

	}
	c.JSON(http.StatusOK, resp)
	return

}

func getLatestTime(c *gin.Context) (int64, error) {
	raw := c.DefaultQuery("latest_time", "")
	if raw != "" {
		latestTime, intError := strconv.ParseInt(raw, 10, 64)
		return latestTime, intError
	}
	return time.Now().Unix(), nil
}
