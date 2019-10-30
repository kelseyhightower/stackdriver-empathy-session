package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

var (
	listenAddr string
	topicId    string
)

func main() {
	flag.StringVar(&listenAddr, "http", "127.0.0.1:8080", "The HTTP listen address")
	flag.StringVar(&topicId, "topic", "empathy", "The Pub/Sub topic to publish messages to")
	flag.Parse()

	projectId := os.Getenv("PROJECT_ID")
	if projectId == "" {
		log.Fatal("PROJECT_ID env var must set and non-empty")
	}

	log.Printf("starting frontend service...")
	log.Printf("listening on http://%s", listenAddr)
	log.Printf("publishing messages to the %s pubsub topic", topicId)

	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("error creating the pubsub client: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("handling request for %s...", r.RemoteAddr)

		message := &pubsub.Message{
			Data: []byte("ping"),
		}

		// Simulate some real work. Nothing fancy.
		time.Sleep(250 * time.Millisecond)

		ctx := context.Background()

		log.Printf("publishing message to the %s topic...", topicId)

		_, err = pubsubClient.Topic(topicId).Publish(ctx, message).Get(ctx)
		if err != nil {
			log.Println("error publishing to the %s topic: %v", topicId, err)
			http.Error(w, err.Error(), 500)
			return
		}

		log.Printf("successfully handled request for %s", r.RemoteAddr)
	})

	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
