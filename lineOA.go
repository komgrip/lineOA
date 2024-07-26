package lineOA

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type serviceLine struct {
	bot *linebot.Client
}

func NewServicesLine(token, secret string) (*serviceLine, error) {
	bot, err := linebot.New(
		secret,
		token,
	)
	if err != nil {
		return nil, err
	}
	return &serviceLine{bot: bot}, nil
}

func (s *serviceLine) LineService(c *gin.Context) ([]*linebot.Event, error) {
	events, err := s.bot.ParseRequest(c.Request)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *serviceLine) ReplyText(replyToken string, message string) error {
	messageReply := linebot.NewTextMessage(message)
	_, err := s.bot.ReplyMessage(replyToken, messageReply).Do()
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceLine) ReplyFlex(replyToken string, textReply string, message string) error {
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(message))
	if err != nil {
		return err
	}

	flexMessage := linebot.NewFlexMessage(textReply, flexContainer)

	_, err = s.bot.ReplyMessage(replyToken, flexMessage).Do()
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceLine) PushText(userLineID string, message string) error {
	messageReply := linebot.NewTextMessage(message)
	_, err := s.bot.PushMessage(userLineID, messageReply).Do()
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceLine) PushFlex(userLineID string, textReply string, message string) error {
	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(message))
	if err != nil {
		return err
	}

	flexMessage := linebot.NewFlexMessage(textReply, flexContainer)

	_, err = s.bot.PushMessage(userLineID, flexMessage).Do()
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceLine) GetProfile(userLineID string) (*linebot.UserProfileResponse, error) {
	user, err := s.bot.GetProfile(userLineID).Do()
	if err != nil {
		return nil, err
	}

	return user, nil
}
