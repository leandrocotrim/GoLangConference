package main

import (
	"context"
	"flag"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()
	projectid := flag.String("projectid", "", "project id")
	topicname := flag.String("topicname", "", "topic name")
	message := flag.String("message", "", "message")
	flag.Parse()

	c, err := pubsub.NewClient(ctx, *projectid)

	if err != nil {
		panic(err)
	}

	defer c.Close() // fecha o client no final da execução da função

	topic := c.Topic(*topicname)
	exists, err := topic.Exists(ctx)
	if err != nil {
		panic(err)
	}

	if !exists {
		topic, err = c.CreateTopic(ctx, *topicname)
		if err != nil {
			panic(err)
		}
	}

	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(*message),
	})

	_, err = result.Get(ctx) // espera a publicação
	if err != nil {
		panic(err)
	}

}

// não finalizou....
