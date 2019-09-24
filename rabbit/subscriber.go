package rabbit

import (
	"encoding/json"
	"log"
	"os"

	logit "github.com/brettallred/go-logit"
	"github.com/nuvi/go-models"
	"github.com/nuvi/justin-sentiment/compliment"
	"github.com/nuvi/justin-sentiment/multinomial"
	"github.com/nuvi/justin-sentiment/tfidf"
	"github.com/nuvi/justin-sentiment/train"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		logit.Fatal(err.Error(), msg)
	}
}

//Consume consumes one message each time this function is called
func Consume(c *train.Classifier) {
	logit.Info("Connecting to RabbitMQ")
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	args := make(map[string]interface{})
	args["x-max-length"] = 100

	queue, err := ch.QueueDeclare("sentiment_trainer.social_activity_parser.twitter_activity.created", true, false, false, false, args)
	failOnError(err, "Could not declare queue")

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			activity := models.CompactSocialActivity{}
			err := json.Unmarshal(d.Body, &activity)
			if err != nil {
				logit.Error("Invalid incoming activity: " + err.Error())
			}

			sentimentCategory := ""
			score := 0.0
			switch os.Getenv("CLASSIFICATION_METHOD") {
			case "tfidf":
				sentimentCategory, score = tfidf.ClassifyTF(c, activity.RawBodyText)
			case "multinomial":
				sentimentCategory, score = multinomial.ClassifyMulti(c, activity.RawBodyText)
			case "compliment":
				sentimentCategory, score = compliment.ClassifyComp(c, activity.RawBodyText)
			}
			activity.SentimentScore = score
			activity.SentimentCategory = sentimentCategory
			publish(activity)
			d.Ack(true)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func unmarshalActivity(payload []byte) (*models.SocialActivity, error) {
	activity := &models.SocialActivity{}
	err := json.Unmarshal(payload, &activity)
	return activity, err
}
