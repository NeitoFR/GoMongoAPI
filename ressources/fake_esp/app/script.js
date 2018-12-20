const mqtt = require("mqtt");

require("dotenv").config();
const json_passport = require("/data/iot_passport_example.json");

var client = mqtt.connect(
  "mqtt://" + process.env.mqtt_host + ":" + process.env.mqtt_port
);

client.on("connect", function() {
  console.log("Connected to MQTT");
  client.subscribe(process.env.question_topic, err => {
    if (err) throw err;
    console.log("Subscribed to topic", process.env.question_topic);
  });
});

client.on("message", (topic, message) => {
  console.log("[" + topic + "] : " + message);

  switch (topic) {
    case process.env.question_topic:
      if (message == "ASK") _sendPassport();
      break;

    default:
      console.log("Not a valid demand");
      client.publish(
        process.env.error_topic,
        '{"error": "Not a valid command"}'
      );
      break;
  }
});

function _sendPassport() {
  client.publish(process.env.answer_topic, JSON.stringify(json_passport));
  console.log("Passport sent");
}
