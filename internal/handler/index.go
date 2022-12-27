package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"shorturl/internal/amqp"
	"shorturl/internal/params"
)

// HandleIndex handles requests for GET /
func HandleIndex(engine *gin.Engine) {
	engine.GET("/", func(context *gin.Context) {

		p := params.NewEnvParams()
		c := amqp.NewChannel(p.Get("AMQP_URL"))
		q := amqp.NewQueue(c, p.Get("AMQP_QUEUE_NAME"))
		if err := q.Publish(amqp091.Publishing{Body: []byte("Yolo")}); err != nil {
			log.Println("publish failed with error: " + err.Error())
		} else {
			log.Println("message published w/o errors")
		}

		context.JSON(http.StatusOK, gin.H{
			"app": "url-shortener",
		})
	})
}
