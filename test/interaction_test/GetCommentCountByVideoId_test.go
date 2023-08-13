package interactiontest

import (
	"TikTok/apps/interaction/rpc/interaction"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommentCountByVideoId(t *testing.T){
	UserComId := make(map[int64]int64)
	//测试增加评论
	{
		for i := 0; i < 100; i++ {
			cmt := fmt.Sprintf("user = %d comment videoId = %d ", i, i/10)
			l := new(string)
			*l = cmt
			resp, err := logic.SendCommentAction(context.Background(), &interaction.CommentActionReq{
				UserId:      int64(i),
				VideoId:     int64(i / 10),
				ActionType:  1,
				CommentText: l,
			})
			assert.Equal(t, err, nil)
			assert.Equal(t, resp.Comment.UserId, int64(i))
			assert.Equal(t, resp.Comment.Content, cmt)
			UserComId[int64(i)] = resp.Comment.Id
		}
	}
	//每个视频都有10个评论
	{
		for i := 0; i < 10; i++{
			count , err := logic.GetCommentCountByVideoId(context.Background() , &interaction.CommentCountByVideoIdReq{
				VideoId: int64(i),
			})
			assert.Equal(t , nil , err)
			assert.Equal(t , int64(10) , count.CommentCount)
		}

	}

	//测试根据id删除评论
	{
		for i := 0; i < 100; i++ {
			tmp := UserComId[int64(i)]
			resp, err := logic.SendCommentAction(context.Background(), &interaction.CommentActionReq{
				UserId:      int64(i),
				VideoId:     int64(i / 10),
				ActionType:  2,
				CommentId: &tmp,
				
			})
			var expected  *interaction.Comment
			assert.Equal(t, nil , err)
			assert.Equal(t , expected  , resp.Comment)
		}
	}
	//每个视频没有评论
	{
		for i := 0; i < 10; i++{
			count , err := logic.GetCommentCountByVideoId(context.Background() , &interaction.CommentCountByVideoIdReq{
				VideoId: int64(i),
			})
			assert.Equal(t , nil , err)
			assert.Equal(t , int64(0) , count.CommentCount)
		}

	}
}
