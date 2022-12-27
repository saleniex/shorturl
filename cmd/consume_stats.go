package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"shorturl/internal/amqp"
	"shorturl/internal/params"
	"shorturl/internal/shorturl"
)

type ConsumeStatsCmd struct {
	params params.Params
	logger *zap.Logger
	repo   shorturl.Repo
}

var defaultQueueName = "shorturl_stats"

func NewConsumeStatsCmd(params params.Params, logger *zap.Logger, repo shorturl.Repo) *ConsumeStatsCmd {
	return &ConsumeStatsCmd{
		params: params,
		logger: logger,
		repo:   repo,
	}
}

func (csc ConsumeStatsCmd) Exec(_ *cobra.Command, _ []string) {
	channel := amqp.NewChannel(csc.params.Get("AMQP_URL"))
	queue := amqp.NewQueue(channel, csc.params.GetWithDefault("AMQP_QUEUE_NAME", defaultQueueName))

	var forever chan struct{}

	if err := queue.Consume(csc.consumeCallback); err != nil {
		csc.logger.Error("cannot start stats consumer", zap.Error(err))
		return
	}

	<-forever
}

// consumeCallback process incoming message
// Callback processes string which is expected to be serialized JSON with structure amqp.AccessMessage
func (csc ConsumeStatsCmd) consumeCallback(s string) {
	var message shorturl.AccessMessage
	if err := json.Unmarshal([]byte(s), &message); err != nil {
		csc.logger.Error("cannot unmarshal consumed message",
			zap.Error(err),
			zap.String("message", s))
		return
	}
	if err := csc.repo.LogAccess(message.ShortId, message.Ip); err != nil {
		csc.logger.Error(
			"error while logging access",
			zap.Error(err),
			zap.String("message", s))
	}
}
