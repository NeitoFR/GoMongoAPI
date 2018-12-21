package main

import (
	"log"
	"net/http"
	"os"

	"github.com/graarh/golang-socketio/transport"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/joho/godotenv"
)

// var server *gosocketio.Server

func _initSocketServer() {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
	log.Println("1")
	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("New client connected")
		c.Emit("hello", "hello from go server")
	})
	log.Println("2")
	serveMux := http.NewServeMux()
	log.Println("3")
	serveMux.Handle("/", server)
	log.Println("4")
	log.Panic(http.ListenAndServe(":"+os.Getenv("socker_server_port"), serveMux))
	log.Println("5")
	log.Println("Listening on port : " + os.Getenv("socker_server_port"))
}

func main() {
	godotenv.Load("./.env")

	_initSocketServer()
	_initMqttConnection(mqttInfo{os.Getenv("mqtt_server"), os.Getenv("mqtt_port")})
}
