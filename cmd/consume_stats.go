package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"shorturl/internal/amqp"
	"shorturl/internal/params"
)

type ConsumeStatsCmd struct {
	params params.Params
	logger *zap.Logger
}

var defaultQueueName = "shorturl_stats"

func NewConsumeStatsCmd(params params.Params, logger *zap.Logger) *ConsumeStatsCmd {
	return &ConsumeStatsCmd{
		params: params,
		logger: logger,
	}
}

func (csc ConsumeStatsCmd) Exec(_ *cobra.Command, _ []string) {
	channel := amqp.NewChannel(csc.params.Get("AMQP_URL"))
	queue := amqp.NewQueue(channel, csc.params.GetWithDefault("AMQP_STATS_QUEUE_NAME", defaultQueueName))

	var forever chan struct{}

	if err := queue.Consume(csc.consumeCallback); err != nil {
		csc.logger.Error("cannot start stats consumer", zap.Error(err))
		return
	}

	<-forever
}

func (csc ConsumeStatsCmd) consumeCallback(_ string) {
	// assume string is shortId so bump up access counter for this short id
}
