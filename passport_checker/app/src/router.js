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

app.get("/hello", function(req, res) {
  res.send("Hello from passport_checker").end();
});

app.get("/checkpassports", function(req, res) {
  var mqtt_passport, mongo_passport;
  w.info("Checking passort");
  var passort_loaded = _.after(2, comparePassports);

  function comparePassports() {
    w.info("Got the 2 passports" + mqtt_passport + mongo_passport);
    compartor.comparePassports(mqtt_passport, mongo_passport, function(str) {
      w.info("Comparaison finished");
      res.send(str).end();
    });
  }

  mqtt_helper.getPassport(str => {
    mqtt_passport = str;
    w.info("Got MQTT passport" + JSON.stringify(str));
    passort_loaded();
  });

  mongo_helper.getPassport(str => {
    mongo_passport = str;
    w.info("Got MongoDB passport" + JSON.stringify(str));
    passort_loaded();
  });
});

app.listen(process.env.port, () => {
  w.info("Example app listening on port : " + process.env.port);
});
