package account

import "net/http"

const (
	RetrieveUserByUsernamePath         = "/users/%s"
	RetrieveFollowagePath              = "/followage/%s/%s"
	RetrieveChannelChattersPath        = "/chatters/%s"
	RetrieveChannelSubscriberCountPath = "/subs/%s/count?token=%s"
	RetrieveUserStreamPath             = "/streams/%s"
	RetrieveChannelEmotesPath          = "/emotes/%s"
)

type Service struct {
	client http.Client
}
