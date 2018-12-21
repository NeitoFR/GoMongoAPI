const w = require("winston"),
  io = require("socket.io-client");
//   WebSocket = require("ws");
// var ws;

// w.add(new w.transports.Console({
//     format: w.format.simple()
// }));

module.exports = {
  _initWebSocketClient: function() {
    var socket = io("ws://mqtt_api/ws", { transports: ["websocket"] });
    // w.info("Trying to connect to websocket");
    // var host =
    //   "ws://" +
    //   process.env.mqtt_api_host +
    //   ":" +
    //   process.env.mqtt_api_port +
    //   "/ws?EIO=3&transport=websocket";
    // ws = new WebSocket(host);
    // ws.on("open", () => {
    //   console.log(
    //     "Connected to WebSocket Server",
    //     process.env.mqtt_api_host + ":" + process.env.mqtt_api_port
    //   );
    //   ws.send('["ask_passport", { Data: "Data test" }]');
    // });
    // ws.on("message", data => {
    //   console.log("message from go", data);
    // });
  },
  getPassport: function(name, callback) {
    w.info("Getting MQTT passport for model : " + name);

    // var host =
    //   "http://" +
    //   process.env.mqtt_api_host +
    //   ":" +
    //   process.env.mqtt_api_port +
    //   "/";
    // var socket = io.connect(
    //   process.env.mqtt_api_host,
    //   { port: 3502, path: "/" }
    // );
    // w.info("Trying to connect to", host);

    // socket.on("connect", () => {
    //   w.info("Connected to Go MQTT_API");
    // });
    // socket.on("hello", msg => {
    //   w.info("Message from go server: " + msg);
    // });
  }
};
