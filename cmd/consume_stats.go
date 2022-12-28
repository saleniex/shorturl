package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"shorturl/internal/amqp"
	"shorturl/internal/params"
	"shorturl/internal/shorturl"
	"time"
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
	failCount := 0
	maxFailCount := 5

	for {
		err := queue.Consume(csc.consumeCallback)
		if err == nil {
			<-forever
		} else {
			csc.logger.Error("Cannot start stats consumer.", zap.Error(err))
			failCount++
			if failCount > maxFailCount {
				csc.logger.Error("Consumer start failed. Backing off.")
				return
			}
			time.Sleep(3 * time.Second)
		}
	}
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
