package rabbit

import (
	"encoding/json"
	"os"

	logit "github.com/brettallred/go-logit"
	"github.com/nuvi/go-models"
	"github.com/streadway/amqp"
)

//Consume consumes one message each time this function is called
func publish(activity models.CompactSocialActivity) {
	logit.Info("Connecting to RabbitMQ")
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	result, err := json.Marshal(activity)
	if err != nil {
		logit.Error("activity was not marshaled correctly " + err.Error())
	}

	queue, err := ch.QueueDeclare("sentiment_trainer.social_activity_parser.twitter_activity.created", true, false, false, false, nil)
	failOnError(err, "Could not declare queue")
	err = ch.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        result,
		})
	failOnError(err, "Failed to publish a message")
}
