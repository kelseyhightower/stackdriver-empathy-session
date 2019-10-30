package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

var (
	subscriptionId string
)

func main() {
	flag.StringVar(&subscriptionId, "subscription", "empathy", "The Pub/Sub subscription ID")
	flag.Parse()

	projectId := os.Getenv("PROJECT_ID")
	if projectId == "" {
		log.Fatal("The PROJECT_ID env var must be set")
	}

	log.Printf("starting backend service...")

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("error creating the pubsub client: %v", err)
	}

	log.Printf("reading from the %s pubsub subscription", subscriptionId)

	subscription := client.Subscription(subscriptionId)
	err = subscription.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		log.Printf("processing message %s...", m.ID)

		// Simulate some real work. Nothing fancy.
		time.Sleep(500 * time.Millisecond)

		log.Printf("message data: %s", m.Data)

		m.Ack()

		log.Printf("successfully processed message %s", m.ID)
	})

	if err != nil {
		log.Fatalf("error reading from the %s subscription: %v", subscriptionId, err)
	}
}
