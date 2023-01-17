package rabbitmq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"post-service/pkg/config"
	"post-service/pkg/post"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ch *amqp.Channel
	q  amqp.Queue
)

func StartServer(cfg config.RabbitMQ, logic post.PostService) {
	fmt.Println("Starting rabbitmq")
	fmt.Println(cfg.Host + ":" + cfg.Port)
	conn, err := amqp.Dial("amqp://guest:guest@" + cfg.Host + ":" + cfg.Port)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	failOnError(err, "Failed to declare a queue")
	createPosts, err := ch.Consume(
		"posts.POST", // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	deletePosts, err := ch.Consume(
		"posts.DELETE", // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	updateTimes, err := ch.Consume(
		"updatePostTime", // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range createPosts {
			log.Printf("Received a message: %s", d.Body)
			fmt.Printf("d.UserId: %v\n", d.UserId)
			var post post.CreatePost
			if err := json.NewDecoder(bytes.NewReader(d.Body)).Decode(&post); err != nil {
				fmt.Println("Unmarshal went wrong")
				return
			}
			post.Uuid = uuid.NewString()
			logic.CreatePost(context.Background(), post)
		}
	}()
	go func() {
		for d := range updateTimes {
			log.Printf("Received a message: %s", d.Body)
			logic.UpdateTime(context.Background(), string(d.Body))
		}
	}()
	go func() {
		for d := range deletePosts {
			log.Printf("Received a delete message: %s", d.Body)
			logic.DeletePost(context.Background(), string(d.Body))
			RemoveResponses(string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func RemoveResponses(postId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	bytes := []byte(postId)
	fmt.Printf("\" remvoing post\": %v\n", " remvoing post")
	err := ch.PublishWithContext(ctx,
		"",                // exchange
		"response.DELETE", // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s\n", postId)
}
