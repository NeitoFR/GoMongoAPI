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
type WsData struct {
	Data string
}

func _initSocketServer() {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("New client connected")
		c.Emit("hello", WsData{"Struct data"})
	})
	server.On("ask_passport", func(c *gosocketio.Channel, msg WsData) string {
		//send event to all in room
		log.Println("Message : " + msg.Data)
		return "OK"
	})
	serveMux := http.NewServeMux()
	serveMux.Handle("/ws", server)
	log.Println("Listening on port : " + os.Getenv("socket_server_port"))
	log.Panic(http.ListenAndServe(":"+os.Getenv("socket_server_port"), serveMux))
}

func main() {
	godotenv.Load("./.env")

	_initMqttConnection(mqttInfo{os.Getenv("mqtt_server"), os.Getenv("mqtt_port")})
	_initSocketServer()
}
