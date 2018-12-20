package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)

type mqttInfo struct {
	host string
	port string
}

var cli *client.Client

func _initMqttConnection(info mqttInfo) {
	cli = client.New(&client.Options{
		ErrorHandler: func(err error) {
			fmt.Println(err)
		},
	})

	err := cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address:  os.Getenv("mqtt_host"),
		ClientID: []byte("mqtt_api"),
	})
	if err != nil {
		panic(err)
	}
	log.Println("Connected to MQTT")

	topic_to_sub := []byte("+/" + os.Getenv("answer_topic"))

	err = cli.Subscribe(&client.SubscribeOptions{
		SubReqs: []*client.SubReq{
			&client.SubReq{
				TopicFilter: topic_to_sub,
				QoS:         mqtt.QoS0,
				// Define the processing of the message handler.
				Handler: func(topicName, message []byte) {
					log.Println("Got message !")
					fmt.Println(string(topicName), string(message))
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	log.Println("Topic registered : " + string(topic_to_sub[:len(topic_to_sub)]))
}
func _askMqttPassports(name string) (string, error) {
	var err error = nil
	// topic := buildAskTopicName(name, os.Getenv("question_topic"))
	topic := []byte(name + "/" + os.Getenv("question_topic"))

	// Publish a message.
	err = cli.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS0,
		TopicName: topic,
		Message:   []byte("ASK"),
	})
	if err != nil {
		panic(err)
	}
	log.Println("ask for passport sent to mqtt")
	return "", err
}
