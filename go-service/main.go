package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	// conn, err := amqp.Dial("amqp://rabbitmq:1jj395qu@localhost:5672/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer conn.Close()

	// ch, err := conn.Channel()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer ch.Close()

	// var queue = "q.go.service"

	// SendRabbitMQ(ch, queue)
	// fmt.Println("send message")

	r := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.POST("/", func(ctx *gin.Context) {
		// SendRabbitMQ(ch, "q-go-service")

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/api", func(c *gin.Context) {
		resp, err := http.Get("http://localhost:8000/js")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		type Data struct {
			Message string `json:"message"`
		}

		data := Data{}

		json.Unmarshal(body, &data)

		c.JSON(http.StatusOK, gin.H{
			"status":  resp.Status,
			"message": "Hello World",
			"data":    data,
		})
	})

	r.Run(":8080")
	// ConsumeRabbitMQ(ch, queue)
}

func ConsumeRabbitMQ(ch *amqp.Channel, queue string) {

	msgs, err := ch.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s\n", d.Body)
			fmt.Println("type :: " + time.Now().Format("2006-01-02 15:04:05"))
			// d.Ack(true)
		}
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func SendRabbitMQ(ch *amqp.Channel, queue string) {

	msg := map[string]interface{}{
		"message": "Hello World!",
	}

	b, _ := json.Marshal(msg)

	err := ch.Publish("ex.sing", queue, false, false, amqp.Publishing{
		ContentType:     "application/json",
		Body:            []byte(b),
		Type:            "go",
		ContentEncoding: "utf-8",
	})
	if err != nil {
		log.Fatal(err)
	}
}
