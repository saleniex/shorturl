package amqp

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

var DefaultQueueName = "shorturl_stats"

// Queue represents AMQP queue
type Queue struct {
	queueName string
	channel   *Channel
	queue     *amqp.Queue
}

// NewQueue creates new queue instance on given channel with given name
func NewQueue(channel *Channel, queueName string) *Queue {
	return &Queue{
		channel:   channel,
		queueName: queueName,
	}
}

// Publish message to queue
// Parameter is Publishing struct which contains all message properties
func (q *Queue) Publish(message amqp.Publishing) error {
	if err := q.declare(); err != nil {
		return err
	}
	channel, channelErr := q.channel.getOpened()
	if channelErr != nil {
		return fmt.Errorf("publishing failed while getting opened channel: %w", channelErr)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	publishErr := channel.PublishWithContext(
		ctx,
		"",
		q.queue.Name,
		false,
		false,
		message,
	)
	if publishErr != nil {
		return fmt.Errorf("publishing failed: %w", publishErr)
	}

	return nil
}

// Consume messages from queue
// Parameter is callback function type which is executed on retrieval of each message
func (q *Queue) Consume(callback ConsumeCallback) error {
	if err := q.declare(); err != nil {
		return err
	}
	channel, channelErr := q.channel.getOpened()
	if channelErr != nil {
		return fmt.Errorf("consuming failed while getting opened channel: %w", channelErr)
	}
	delivery, consumeErr := channel.Consume(
		q.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	if consumeErr != nil {
		return fmt.Errorf("consuming failed: %w", consumeErr)
	}

	go func(deliveryChan <-chan amqp.Delivery) {
		for d := range deliveryChan {
			callback(string(d.Body))
		}
	}(delivery)

	return nil
}

func (q *Queue) declare() error {
	if q.queue != nil {
		return nil
	}
	channel, chanErr := q.channel.getOpened()
	if chanErr != nil {
		return fmt.Errorf("queue declaring failed while getting opened channel: %w", chanErr)
	}
	queue, queueErr := channel.QueueDeclare(
		q.queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if queueErr != nil {
		return fmt.Errorf("queue declaring failed: %w", queueErr)
	}
	q.queue = &queue

	return nil
}
