package lineOA

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
)

type ServiceLine struct {
	bot *linebot.Client
}

func NewServicesLine(token, secret string) (*ServiceLine, error) {
	bot, err := linebot.New(
		secret,
		token,
	)
	if err != nil {
		return nil, err
	}
	return &ServiceLine{bot: bot}, nil
}

func (s *ServiceLine) LineService(c *gin.Context) ([]*linebot.Event, error) {
	events, err := s.bot.ParseRequest(c.Request)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *ServiceLine) ReplyText(replyToken string, message string) error {
	messageReply := linebot.NewTextMessage(message)
	_, err := s.bot.ReplyMessage(replyToken, messageReply).Do()
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceLine) ReplyFlex(replyToken string, textReply string, message string) error {
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

func (s *ServiceLine) ReplyTextEmoji(replyToken string, message string, emojis []*linebot.Emoji) error {
	messageReply := linebot.NewTextMessage(message)
	messageReply.Emojis = emojis
	_, err := s.bot.ReplyMessage(replyToken, messageReply).Do()
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceLine) PushText(userLineID string, message string) error {
	messageReply := linebot.NewTextMessage(message)
	_, err := s.bot.PushMessage(userLineID, messageReply).Do()
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceLine) PushTextEmoji(userLineID string, message string, emojis []*linebot.Emoji) error {
	messageReply := linebot.NewTextMessage(message)
	messageReply.Emojis = emojis
	_, err := s.bot.PushMessage(userLineID, messageReply).Do()
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceLine) PushFlex(userLineID string, textReply string, message string) error {
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

func (s *ServiceLine) GetProfile(userLineID string) (*linebot.UserProfileResponse, error) {
	user, err := s.bot.GetProfile(userLineID).Do()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *ServiceLine) ReplyImage(replyToken string, originalContentURL string, previewImageURL string) error {
	messageReply := linebot.NewImageMessage(originalContentURL, previewImageURL)
	_, err := s.bot.ReplyMessage(replyToken, messageReply).Do()
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceLine) PushImage(userLineID string, originalContentURL string, previewImageURL string) error {
	messageReply := linebot.NewImageMessage(originalContentURL, previewImageURL)
	_, err := s.bot.PushMessage(userLineID, messageReply).Do()
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceLine) GetProfileLineGroup(groupID string) (*linebot.GroupSummaryResponse, error) {
	profile, err := s.bot.GetGroupSummary(groupID).Do()
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *ServiceLine) HealthCheck() (*linebot.TestWebhookResponse, error) {
	response, err := s.bot.TestWebhook().Do()
	if err != nil {
		return nil, err
	}

	return response, nil
}
