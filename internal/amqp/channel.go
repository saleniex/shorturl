package amqp

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Channel represents AMQP channel
type Channel struct {
	url     string
	channel *amqp.Channel
}

// NewChannel creates new channel instance
// Channel is not opened at this moment
func NewChannel(url string) *Channel {
	return &Channel{url: url}
}

func (c *Channel) getOpened() (*amqp.Channel, error) {
	if c.channel == nil {
		if err := c.openChannel(); err != nil {
			return nil, err
		}
	}
	if c.channel.IsClosed() {
		return nil, fmt.Errorf("amqp channel is closed")
	}
	return c.channel, nil
}

func (c *Channel) openChannel() error {
	conn, dialErr := amqp.Dial(c.url)
	if dialErr != nil {
		return dialErr
	}

	channel, channelErr := conn.Channel()
	if channelErr != nil {
		return fmt.Errorf("opening connection channel failed: %s", channelErr)
	}
	c.channel = channel

	return nil
}
