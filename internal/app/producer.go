package app

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	kfk "github.com/segmentio/kafka-go"
	"github.com/tavvfiq/go-kafka-testbed/internal/pkg/kafka"
)

type producerApp struct {
	kafka kafka.KafkaWriter
}

func StartProducer() {
	writer := kafka.NewWriter([]string{"localhost:9092"}, "producer1", "test-topic", 10*time.Second)
	app := &producerApp{
		kafka: writer,
	}
	http.HandleFunc("/v1/publish", app.publishData)
	fmt.Println("server started at localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func (app *producerApp) publishData(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = app.publishToKafka(r.Context(), nil, body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (app *producerApp) publishToKafka(ctx context.Context, key, value []byte) error {
	message := kfk.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}
	return app.kafka.Push(ctx, message)
}
