package main

import (
  "context"
  //"time"
  "github.com/segmentio/kafka-go"
)

func main() {
  // to produce messages
 // topic := "foo"
 // partition := 0

  w := kafka.NewWriter(kafka.WriterConfig{
      Brokers: []string{"10.0.0.16"},
      Topic:   "foo",
  })

  w.WriteMessages(context.Background(),
    kafka.Message{
        Key:   []byte("Key-A"),
        Value: []byte("Hello World!"),
    },
  )

  w.Close()

  //conn, err := kafka.DialLeader(context.Background(), "tcp", "10.0.0.16.", topic, partition)

  //conn.SetWriteDeadline(time.Now().Add(10*time.Second))
  //conn.WriteMessages(
  //    kafka.Message{Value: []byte("one!")},
  //    kafka.Message{Value: []byte("two!")},
  //    kafka.Message{Value: []byte("three!")},
  //)

//  conn.Close()
}
