package shorturl

import (
	"encoding/json"
	"fmt"
	rabbit "github.com/rabbitmq/amqp091-go"
	"shorturl/internal/amqp"
)

// DistributedRepo is repository optimized for performance
// Reading is done via Redis cache while new URL registration and stats writing operations are performed asynchronous
// using AMQP queue
type DistributedRepo struct {
	repo  *CachedRepo
	queue *amqp.Queue
}

// NewDistributedRepo create new distributed repository
// repo
func NewDistributedRepo(repo *CachedRepo, queue *amqp.Queue) *DistributedRepo {
	return &DistributedRepo{
		repo:  repo,
		queue: queue,
	}
}

func (d *DistributedRepo) StoreUrl(shortUrl ShortUrl) error {
	// TODO perform URL storing using AMQP in order to reduce load on gateway
	return d.repo.StoreUrl(shortUrl)
}

func (d *DistributedRepo) Find(shortId string) string {
	return d.repo.Find(shortId)
}

func (d *DistributedRepo) LogAccess(shortId, remoteIp string) error {
	accessMessage := AccessMessage{
		Ip:      remoteIp,
		ShortId: shortId,
	}
	serializedAccessMessage, err := json.Marshal(accessMessage)
	if err != nil {
		return fmt.Errorf("error while serializing log access message: %w", err)
	}

	publishMessage := rabbit.Publishing{
		ContentType: "application/json",
		Body:        serializedAccessMessage,
	}

	return d.queue.Publish(publishMessage)
}

func (d *DistributedRepo) ShortUrlAccessStats(shortId string) (*AccessStats, error) {
	return d.repo.ShortUrlAccessStats(shortId)
}
