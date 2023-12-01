package messaging

import (
	"github.com/slack-go/slack"
)

// SendMessage sends a message to a Slack channel using the provided token.
func SendMessage(token, channelID, message string) error {
	client := slack.New(token)
	_, _, err := client.PostMessage(channelID, slack.MsgOptionText(message, false))
	return err
}
