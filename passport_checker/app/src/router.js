const express = require("express");
const app = express(),
  _ = require("lodash"),
  w = require("winston"),
  mqtt_helper = require("./mqtt"),
  mongo_helper = require("./mongo"),
  compartor = require("./passport_comparator");

require("dotenv").config();

w.add(
  new w.transports.Console({
    format: w.format.simple()
  })
);
mqtt_helper._initMqttConnexion();
app.get("/hello", function (req, res) {
  res.send("Hello from passport_checker").end();
});

app.get("/passports", function (req, res) {
  w.info("********** New Request **********");
  if (req.query.model_Name == undefined || ""){
    w.info("model_Name missing in query parameters")
    res.send('{"error":"model_Name required in query parameters"}').end();
  }
  else {
    console.log(
      "Requested passport checking for model : " + req.query.model_Name
    );
    var name = req.query.model_Name;
    var response_from_mqtt, response_from_mongo;

    w.info("Checking passort");
    var passort_loaded = _.after(2, comparePassports);

    function comparePassports() {
      w.info("Got the response from MQTT and MongoDB : " + response_from_mqtt + response_from_mongo);
      compartor.comparePassports(req.query.model_Name, response_from_mqtt, response_from_mongo, function (str) {
        w.info("********** Comparaison finished **********");
        res.send(str).end();
        w.info("********** End of Request **********");
      });
    }

    mqtt_helper.getPassport(name, str => {
      response_from_mqtt = str;
      w.info("Got MQTT response : " + JSON.stringify(str));
      passort_loaded();
    });

    mongo_helper.getPassport(name, str => {
      response_from_mongo = str;
      w.info("Got MongoDB response : " + JSON.stringify(str));
      passort_loaded();
    });
  }
});

app.listen(process.env.port, () => {
  w.info("Passport Checker listening on port : " + process.env.port);
});
