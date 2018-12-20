w = require("winston");

// w.add(new w.transports.Console({
//     format: w.format.simple()
// }));

module.exports = {
  comparePassports: function(pMqtt, pMongo, callback) {
    var final = {
      "MongoDB_Passport": pMongo,
      "MQTT_Passport": pMqtt || "none yet"
    }
    //TODO Implement passport concatenation and comparaison
    callback(final);
  }
};
