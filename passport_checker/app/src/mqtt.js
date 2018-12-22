const w = require("winston");
mqtt = require('mqtt')
var client = null;
var passport_temp = [],
  ongoing_request = false;
module.exports = {
  _initMqttConnexion: function () {
    client = mqtt.connect(
      "mqtt://" + process.env.mqtt_host + ":" + process.env.mqtt_port
    );

    client.on("connect", function () {
      w.info("Connected to MQTT");
      var topic = "+/" + process.env.answer_topic
      client.subscribe(topic, err => {
        if (err) throw err;
        w.info("Subscribed to topic : " + topic);
      });
    });
    client.on("message", (topic, message) => {
      w.info("[" + topic + "] : " + message + "\n*************************\n");
      if (ongoing_request && _checkTopic(topic)) {
        w.info('topic check ok')
        try {
          var passport = JSON.parse(message)
          passport_temp.push(passport)
        } catch (error) {
          passport_temp.push({ error: "Can't parse mqtt response : " + error })
        }
      }
      else {
        passport_temp = { "error": "improper topic" }

      }
    })
  },
  getPassport: function (name, callback) {
    w.info("Getting MQTT passport for model : " + name);
    ongoing_request = true;
    var topic = name + "/" + process.env.question_topic
    client.publish(topic, "ASK")
    setTimeout(() => {
      w.info((process.env.mqtt_timeout / 1000) + " seconds have passed, " + (passport_temp.length || 0) + " passport(s) received")
      callback(passport_temp)
      passport_temp.length = 0;
      ongoing_request = false;
    }, process.env.mqtt_timeout || 5000);
  }
};

function _checkTopic(topic) {
  w.info("checking topic : " + topic)
  var temp = topic.split("/");
  w.info("topic parts" + temp.length + temp)
  w.info(temp[temp.length - 1] + " = " + process.env.answer_topic + " ? " + (temp[temp.length] === process.env.answer_topic))
  w.info(temp.length + " = " + 2 + " ? " + (temp.length === 2))

  return temp[temp.length - 1] === process.env.answer_topic && temp.length === 2 ? true : false
}
