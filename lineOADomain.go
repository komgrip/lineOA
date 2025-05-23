package lineOA

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
)

type LineServiceDomain interface {
	LineService(c *gin.Context) ([]*linebot.Event, error)
	ReplyText(replyToken string, message string) error
	ReplyTextEmoji(replyToken string, message string, emojis []*linebot.Emoji) error
	ReplyFlex(replyToken string, textReply string, message string) error
	PushText(userLineID string, message string) error
	PushTextEmoji(userLineID string, message string, emojis []*linebot.Emoji) error
	PushFlex(userLineID string, textReply string, message string) error
	ReplyImage(replyToken string, originalContentURL string, previewImageURL string) error
	PushImage(userLineID string, originalContentURL string, previewImageURL string) error
	GetProfile(userLineID string) (*linebot.UserProfileResponse, error)
	GetProfileLineGroup(groupID string) (*linebot.GroupSummaryResponse, error)
	HealthCheck() (*linebot.TestWebhookResponse, error)
}
