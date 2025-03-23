package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

const TopicName = "mytopic"

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)
	go func() {
		<-sig
		cancel()
	}()

	switch os.Args[1] {
	case "producer":
		writer := &kafka.Writer{
			Addr:         kafka.TCP("localhost:9092"),
			Async:        true,
			RequiredAcks: kafka.RequireOne,
			BatchSize:    100,
			BatchTimeout: 1 * time.Second,
			Compression:  kafka.Lz4,
		}
		defer writer.Close()

		for {
			fmt.Print("> ")
			reader := bufio.NewReader(os.Stdin)
			response, err := reader.ReadString('\n')
			if err != nil {
				if errors.Is(err, io.EOF) {
					fmt.Println("\n🔴 Input closed, exiting producer.")
					break
				}

				fmt.Println("❌ Error reading input:", err)
				continue
			}

			response = strings.TrimSuffix(response, "\n")
			if response == "" {
				break
			}

			if err := writer.WriteMessages(ctx, kafka.Message{
				Topic: TopicName,
				Value: []byte(response),
				Key:   nil,
			}); err != nil {
				if errors.Is(err, context.Canceled) {
					fmt.Println("\n🔴 Producer stopped.")
					break
				}

				fmt.Println("❌ Error sending message:", err)
				break
			}
		}
	case "consumer":
		groupId := os.Args[2]
		consumer := kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{"localhost:9092"},
			Topic:   TopicName,
			GroupID: groupId,
		})

		defer consumer.Close()

		for {
			m, err := consumer.FetchMessage(ctx)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					fmt.Println("\n🔴 Consumer stopped.")
					break
				}

				fmt.Println("❌ Error fetching message:", err)
				break
			}

			fmt.Println(string(m.Value))
			err = consumer.CommitMessages(ctx, m)
			if err != nil {
				fmt.Println("❌ Error fetching message:", err)
				break
			}
		}

	}

}
